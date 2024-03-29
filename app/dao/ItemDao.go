package dao

import (
	"github.com/jinzhu/gorm"
	"zuoxingtao/constant"
	"zuoxingtao/dto/item"
	"zuoxingtao/init/db"
)

type ItemDao struct {
	*gorm.DB
}

func NewItemDao(ops ...func(*ItemDao)) *ItemDao {
	id := &ItemDao{
		DB: db.DB,
	}
	for _, op := range ops {
		op(id)
	}
	return id
}

func (itemDao *ItemDao) AddItem(item *item.Item) error {
	if err := itemDao.Table(item.TableName()).Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (itemDao *ItemDao) UpdateItem(item *item.Item) error {
	if err := itemDao.Table(item.TableName()).Where("item_code  = ?", item.ItemCode).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (itemDao *ItemDao) GetItemByID(id string) (*item.Item, error) {
	item := &item.Item{}
	if err := itemDao.Table(item.TableName()).Where("item_code = ?", id).Find(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (itemDao *ItemDao) GetItemByStatus(status int) ([]*item.Item, error) {
	it := &item.Item{}
	items := make([]*item.Item, 0)
	if err := itemDao.Table(it.TableName()).Where("status = ?", status).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (itemDao *ItemDao) CloseAllItem() error {
	item := item.Item{}
	if err := itemDao.Table(item.TableName()).Where("status = ?", constant.ITEM_OPEN).Update("status", constant.ITEM_CLOSE).Error; err != nil {
		return err
	}
	return nil
}
