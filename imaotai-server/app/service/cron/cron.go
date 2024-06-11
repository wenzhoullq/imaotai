package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"imaotai_helper/constant"
	"imaotai_helper/dao"
	"imaotai_helper/dto/item"
	item2 "imaotai_helper/dto/item"
	"imaotai_helper/dto/record"
	"imaotai_helper/dto/response"
	"imaotai_helper/dto/shop"
	u "imaotai_helper/dto/user"
	"imaotai_helper/init/common"
	"imaotai_helper/init/log"
	"imaotai_helper/lib"
	"imaotai_helper/lib/client"
	"math"
	"time"
)

type CronModel struct {
	*logrus.Logger
	*client.ShopClient
	*dao.ShopDao
	*dao.ItemDao
	*dao.UserDao
	*dao.RecordDao
	*client.UserClient
	*client.ItemClient
}

func SetLog() func(model *CronModel) {
	return func(model *CronModel) {
		model.Logger = log.Logger
	}
}
func NewCronModel(ops ...func(model *CronModel)) *CronModel {
	sm := &CronModel{
		ShopClient: client.NewShopClient(),
		ShopDao:    dao.NewShopDao(),
		ItemDao:    dao.NewItemDao(),
		UserDao:    dao.NewUserDao(),
		RecordDao:  dao.NewRecordDao(),
		ItemClient: client.NewItemClient(),
		UserClient: client.NewUserClient(),
	}
	for _, op := range ops {
		op(sm)
	}
	return sm
}
func (cm *CronModel) ParseShopJson(jsonStr []byte) (map[string]*shop.Shop, error) {
	shopMap := make(map[string]*shop.Shop)
	err := json.Unmarshal(jsonStr, &shopMap)
	if err != nil {
		return nil, err
	}
	return shopMap, nil
}

func (cm *CronModel) UpdateShop() error {
	url, err := cm.GetShopListUrl()
	if err != nil {
		return err
	}
	jsonStr, err := cm.GetShopList(url)
	if err != nil {
		return err
	}
	shopMap, err := cm.ParseShopJson(jsonStr)
	if err != nil {
		return err
	}
	// 先将所有店铺关闭
	err = cm.CloseAllShop()
	if err != nil {
		return err
	}
	// 批量更新Shop信息
	failTask := 0
	for _, v := range shopMap {
		_, err = cm.GetShopByID(v.ShopId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				v.Status = constant.SHOP_OPEN
				err = cm.AddShop(v)
				if err != nil {
					failTask++
					cm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
				}
			} else {
				failTask++
				cm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
			}
			continue
		}
		v.Status = constant.SHOP_OPEN
		err = cm.ShopDao.UpdateShop(v)
		if err != nil {
			failTask++
			cm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
			continue
		}
	}
	cm.Logln(logrus.InfoLevel, "fail shop Update/Create time  is ", failTask)
	return nil
}
func (cm *CronModel) UpdateItems() error {
	// 先将所有酒关闭
	items, err := cm.GetItemList()
	if err != nil {
		return err
	}
	err = cm.CloseAllItem()
	if err != nil {
		return err
	}
	failTask := 0
	for _, v := range items {
		item, err := cm.GetItemByID(v.ItemCode)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				item = &item2.Item{
					ItemCode: v.ItemCode,
					Title:    v.Title,
					Status:   constant.ITEM_OPEN,
				}
				err = cm.AddItem(item)
				if err != nil {
					failTask++
					cm.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
				}
			} else {
				failTask++
				cm.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
			}
			continue
		}
		item.Status = constant.ITEM_OPEN
		err = cm.UpdateItem(item)
		if err != nil {
			failTask++
			cm.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
			continue
		}
	}
	cm.Logln(logrus.InfoLevel, "fail item Update/Create time  is ", fmt.Sprintf("%d", failTask))
	return nil
}
func (cm *CronModel) ExpUser() error {
	users, err := cm.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	expNum := 0
	for _, u := range users {
		token, err := lib.ParseImaoTaiToken(u.Token)
		//解析失败直接改变状态
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err.Error())
			u.Status = constant.USER_ABNORMAL
			err = cm.UpdateUserStatus(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err.Error())
			}
			expNum++
			continue
		}
		if lib.Overdue(token.Exp) {
			u.Status = constant.USER_ABNORMAL
			err = cm.UpdateUserStatus(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			expNum++
			cm.Logln(logrus.InfoLevel, "user :", u.UserID, u.UserName, "has overdue")
		}
	}
	cm.Logln(logrus.InfoLevel, "overDue user num:", expNum)
	return nil
}

func (cm *CronModel) FilterItem(items []*item.Item) []*item.Item {
	filterItems := make([]*item.Item, 0)
	for _, v := range items {
		if _, ok := common.FilterSet[v.ItemCode]; !ok {
			continue
		}
		filterItems = append(filterItems, v)
	}
	return filterItems
}

func (cm *CronModel) GetMinDistanceShop(user *u.User, shopIDs []string) (string, error) {
	var shopID string
	MinDis := math.MaxFloat64
	for _, v := range shopIDs {
		shop, err := cm.GetShopByID(v)
		if err != nil {
			return "", err
		}
		dis := lib.CalDis(user.Lat, user.Lng, shop.Lat, shop.Lng)
		if dis < MinDis {
			MinDis = dis
			shopID = v
		}
	}
	return shopID, nil
}

func (cm *CronModel) Reservation() error {
	users, err := cm.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	items, err := cm.GetItemByStatus(constant.ITEM_OPEN)
	if err != nil {
		return err
	}
	// 过滤掉不想要的酒
	items = cm.FilterItem(items)
	for _, user := range users {
		for _, item := range items {
			time.Sleep(time.Second * 3)
			shopsID, err := cm.GetShopsByProvince(user.ProvinceName, item.ItemCode)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			//选择距离最近的shop
			shopID, err := cm.GetMinDistanceShop(user, shopsID)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			err = cm.PostReserve(user, item.ItemCode, shopID)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err.Error(), "userID:", user.UserID, " reserve shopID:", shopID, " sessionID:", common.SessionID, " itemCode", item.ItemCode, " fail")
				continue
			}
		}
	}
	return nil
}

func (cm *CronModel) parseRecord(user *u.User, resp *response.RecordResp) []*record.Record {
	records := make([]*record.Record, 0)
	for _, v := range resp.Data.ReservationItemVOS {
		if v.Status != constant.AWARD || v.SessionType != constant.SESSION_TYPE_NORMAL_ORDER {
			continue
		}
		record := &record.Record{
			UserID:   user.UserID,
			UserName: user.UserName,
			ItemName: v.ItemName,
			ItemID:   v.ItemID,
		}
		records = append(records, record)
	}
	return records
}

func (cm *CronModel) AddRecord() error {
	users, err := cm.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	for _, u := range users {
		resp, err := cm.GetAppointmentRecord(u)
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err.Error())
			continue
		}
		records := cm.parseRecord(u, resp)
		for _, vv := range records {
			fmt.Sprintf("%#v", vv)
		}
		err = cm.AddRecords(records)
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err.Error())
			continue
		}
	}
	return nil
}

func (cm *CronModel) TravelReward() error {
	users, err := cm.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	for _, u := range users {
		pageData, err := cm.GetUserIsolationPageData(u)
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err)
			continue
		}
		//现存体力
		curEnergy := pageData.Data.Energy
		// 获得体力
		if pageData.Data.EnergyReward.Value > 0 {
			err = cm.GetEnergyAward(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err)
				continue
			}
			curEnergy = curEnergy + pageData.Data.EnergyReward.Value
		}
		//正在旅行中
		if pageData.Data.XmTravel.Status == constant.TRAVEL_STATUS_PROCESSING {
			cm.Logln(logrus.InfoLevel, "userID:", u.UserID, " is traveling")
			continue
		}
		// 如果旅行结束了,获取小茅运和首次分享获得体力值(该体力值不计入curEnergy,避免逻辑复杂化)
		if pageData.Data.XmTravel.Status == constant.TRAVEL_STATUS_FINISH {
			err = cm.ReceiveReward(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err)
				continue
			}
			err = cm.ShareReward(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err)
				continue
			}
		}
		travelRewardXmy, err := cm.GetXmTravelReward(u)
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err)
			continue
		}
		exchangeRateInfo, err := cm.GetExchangeRateInfo(u)
		if err != nil {
			cm.Logln(logrus.ErrorLevel, err)
			continue
		}
		// 本月小茅运还有余额;今日次数还有;体力值大于一次旅行的消耗量;
		if exchangeRateInfo > travelRewardXmy && curEnergy >= constant.TRAVEL_CONSUME && pageData.Data.XmTravel.RemainChance > 0 {
			err := cm.StartTravel(u)
			if err != nil {
				cm.Logln(logrus.ErrorLevel, err)
				continue
			}
		}
		cm.Logln(logrus.InfoLevel, fmt.Sprintf("TravelReward success,userID:%d", u.UserID))
	}
	cm.Logln(logrus.InfoLevel, "get reward Success")
	return nil
}
