package dao

import (
	"github.com/jinzhu/gorm"
	"imaotai_helper/constant"
	"imaotai_helper/dto/admin"
	"imaotai_helper/init/db"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao(ops ...func(*AdminDao)) *AdminDao {
	ad := &AdminDao{
		DB: db.DB,
	}
	for _, op := range ops {
		op(ad)
	}
	return ad
}

func (ad *AdminDao) AddAdmin(admin *admin.Admin) error {
	if err := ad.Table(admin.TableName()).Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (ad *AdminDao) UpdateAdmin(admin *admin.Admin) error {
	if err := ad.Table(admin.TableName()).Where("uid  = ?", admin.UID).Updates(admin).Error; err != nil {
		return err
	}
	return nil
}

func (ad *AdminDao) GetAdminByMobile(mobile string) (*admin.Admin, error) {
	admin := &admin.Admin{}
	if err := ad.Table(admin.TableName()).Where("mobile = ? and status = ? and deleted = ?", mobile, constant.ADMIN_NORMAL, constant.NORMAL).Find(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
func (ad *AdminDao) GetAdminByUid(uid string) (*admin.Admin, error) {
	admin := &admin.Admin{}
	if err := ad.Table(admin.TableName()).Where("uid = ? and status = ? and deleted = ?", uid, constant.ADMIN_NORMAL, constant.NORMAL).Find(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

//func (itemDao *ItemDao) GetItemByID(id string) (*item.Item, error) {
//	item := &item.Item{}
//	if err := itemDao.Table(item.TableName()).Where("item_code = ?", id).Find(item).Error; err != nil {
//		return nil, err
//	}
//	return item, nil
//}

//func (itemDao *ItemDao) GetItemByStatus(status int) ([]*item.Item, error) {
//	it := &item.Item{}
//	items := make([]*item.Item, 0)
//	if err := itemDao.Table(it.TableName()).Where("status = ?", status).Find(&items).Error; err != nil {
//		return nil, err
//	}
//	return items, nil
//}
