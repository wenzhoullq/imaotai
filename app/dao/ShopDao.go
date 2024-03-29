package dao

import (
	"github.com/jinzhu/gorm"
	"zuoxingtao/constant"
	"zuoxingtao/dto/shop"
	"zuoxingtao/init/db"
)

type ShopDao struct {
	*gorm.DB
}

func NewShopDao(ops ...func(*ShopDao)) *ShopDao {
	sd := &ShopDao{
		DB: db.DB,
	}
	for _, op := range ops {
		op(sd)
	}
	return sd
}

func (sd *ShopDao) AddShop(shop *shop.Shop) error {
	if err := sd.Table(shop.TableName()).Create(shop).Error; err != nil {
		return err
	}
	return nil
}

func (sd *ShopDao) UpdateShop(shop *shop.Shop) error {
	if err := sd.Table(shop.TableName()).Where("shop_id  = ?", shop.ShopId).Updates(shop).Error; err != nil {
		return err
	}
	return nil
}

func (sd *ShopDao) GetShopsByStatus(status int) ([]*shop.Shop, error) {
	s := &shop.Shop{}
	shops := make([]*shop.Shop, 0)
	if err := sd.Table(s.TableName()).Where("status = ?", status).Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (sd *ShopDao) GetShopByID(id string) (*shop.Shop, error) {
	shop := &shop.Shop{}
	if err := sd.Table(shop.TableName()).Where("shop_id = ?", id).Find(shop).Error; err != nil {
		return nil, err
	}
	return shop, nil
}

func (sd *ShopDao) CloseAllShop() error {
	shop := shop.Shop{}
	if err := sd.Table(shop.TableName()).Where("status = ?", constant.SHOP_OPEN).Update("status", constant.SHOP_CLOSE).Error; err != nil {
		return err
	}
	return nil
}
