package item

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"zuoxingtao/constant"
	"zuoxingtao/dao"
	item2 "zuoxingtao/dto/item"
	"zuoxingtao/init/log"
	"zuoxingtao/lib/client"
)

type ItemModel struct {
	*logrus.Logger
	*dao.ItemDao
	*client.ItemClient
}

func SetLog() func(model *ItemModel) {
	return func(model *ItemModel) {
		model.Logger = log.Logger
	}
}

func NewItemModel(ops ...func(model *ItemModel)) *ItemModel {
	im := &ItemModel{
		ItemDao:    dao.NewItemDao(),
		ItemClient: client.NewItemClient(),
	}
	for _, op := range ops {
		op(im)
	}
	return im
}

func (im *ItemModel) UpdateItems() error {
	// 先将所有酒关闭
	items, err := im.GetItemList()
	if err != nil {
		return err
	}
	err = im.CloseAllItem()
	if err != nil {
		return err
	}
	failTask := 0
	for _, v := range items {
		item, err := im.GetItemByID(v.ItemCode)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				item = &item2.Item{
					ItemCode: v.ItemCode,
					Title:    v.Title,
					Status:   constant.ITEM_OPEN,
				}
				err = im.AddItem(item)
				if err != nil {
					failTask++
					im.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
				}
			} else {
				failTask++
				im.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
			}
			continue
		}
		item.Status = constant.ITEM_OPEN
		err = im.UpdateItem(item)
		if err != nil {
			failTask++
			im.Logln(logrus.ErrorLevel, "itemCode is "+string(v.ItemCode)+err.Error())
			continue
		}
	}
	im.Logln(logrus.InfoLevel, "fail item Update/Create time  is ", fmt.Sprintf("%d", failTask))
	return nil
}
