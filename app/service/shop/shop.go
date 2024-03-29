package shop

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"zuoxingtao/constant"
	"zuoxingtao/dao"
	"zuoxingtao/dto/shop"
	"zuoxingtao/init/log"
	"zuoxingtao/lib/client"
)

type ShopModel struct {
	*logrus.Logger
	*client.ShopClient
	*dao.ShopDao
}

func SetLog() func(model *ShopModel) {
	return func(model *ShopModel) {
		model.Logger = log.Logger
	}
}

func NewShopModel(ops ...func(model *ShopModel)) *ShopModel {
	sm := &ShopModel{
		ShopClient: client.NewShopClient(),
		ShopDao:    dao.NewShopDao(),
	}
	for _, op := range ops {
		op(sm)
	}
	return sm
}

func (sm *ShopModel) ParseShopJson(jsonStr []byte) (map[string]*shop.Shop, error) {
	shopMap := make(map[string]*shop.Shop)
	err := json.Unmarshal(jsonStr, &shopMap)
	if err != nil {
		return nil, err
	}
	return shopMap, nil
}

func (sm *ShopModel) UpdateShop() {
	url, err := sm.GetShopListUrl()
	if err != nil {
		return
	}
	jsonStr, err := sm.GetShopList(url)
	if err != nil {
		sm.Logln(logrus.ErrorLevel, err.Error())
		return
	}
	shopMap, err := sm.ParseShopJson(jsonStr)
	if err != nil {
		sm.Logln(logrus.ErrorLevel, err.Error())
		return
	}
	// 先将所有店铺关闭
	err = sm.CloseAllShop()
	if err != nil {
		sm.Logln(logrus.ErrorLevel, err.Error())
		return
	}
	// 批量更新Shop信息
	failTask := 0
	for _, v := range shopMap {
		_, err = sm.GetShopByID(v.ShopId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				v.Status = constant.SHOP_OPEN
				err = sm.AddShop(v)
				if err != nil {
					failTask++
					sm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
				}
			} else {
				failTask++
				sm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
			}
			continue
		}
		v.Status = constant.SHOP_OPEN
		//lib.WriteToTxt(fmt.)
		err = sm.ShopDao.UpdateShop(v)
		if err != nil {
			failTask++
			sm.Logln(logrus.ErrorLevel, "shopID is "+string(v.ShopId)+err.Error())
			continue
		}
	}
	sm.Logln(logrus.InfoLevel, "fail shop Update/Create time  is ", failTask)
	return
}
