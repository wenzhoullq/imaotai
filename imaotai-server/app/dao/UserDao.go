package dao

import (
	"github.com/jinzhu/gorm"
	"imaotai_helper/constant"
	"imaotai_helper/dto/user"
	"imaotai_helper/init/db"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ops ...func(userDao *UserDao)) *UserDao {
	ud := &UserDao{
		DB: db.DB,
	}
	for _, op := range ops {
		op(ud)
	}
	return ud
}

func (ud *UserDao) GetUserByUidAndMobile(uid, mobile string) (*user.User, error) {
	user := &user.User{}
	if err := ud.Table(user.TableName()).Where("uid = ? and mobile = ? and status = ?", uid, mobile, constant.ADMIN_NORMAL).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ud *UserDao) GetUsersByStatus(status int) ([]*user.User, error) {
	u := &user.User{}
	users := make([]*user.User, 0)
	if err := ud.Table(u.TableName()).Where("status = ?", status).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ud *UserDao) GetFlowerUserListByAdminUid(uid string, page, pageSize int) ([]*user.User, error) {
	userList := make([]*user.User, 0)
	user := &user.User{}
	if err := ud.Table(user.TableName()).Where("admin_uid = ? and deleted = ?", uid, constant.NORMAL).Offset((page - 1) * pageSize).Limit(pageSize).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func (ud *UserDao) GetUserByMobile(mobile string) (*user.User, error) {
	user := &user.User{}
	if err := ud.Table(user.TableName()).Where("mobile = ?", mobile).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ud *UserDao) GetUserByUid(uid string) (*user.User, error) {
	user := &user.User{}
	if err := ud.Table(user.TableName()).Where("uid = ?", uid).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ud *UserDao) AddUser(user *user.User) error {
	if err := ud.Table(user.TableName()).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserDao) UpdateUser(user *user.User) error {
	if err := ud.Table(user.TableName()).Where("mobile = ?", user.Mobile).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserDao) UpdateUserStatus(user *user.User) error {
	if err := ud.Table(user.TableName()).Where("uid = ?", user.UID).Update("status", user.Status).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserDao) UpdateToken(user *user.User) error {
	if err := ud.Table(user.TableName()).Where("mobile = ?", user.Mobile).Update("token", user.Token).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserDao) DeleteUser(uid string) error {
	user := &user.User{}
	if err := ud.Table(user.TableName()).Where("uid = ?", uid).Update("deleted", constant.Deleted).Error; err != nil {
		return err
	}
	return nil
}
