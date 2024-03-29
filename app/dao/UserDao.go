package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"zuoxingtao/dto/user"
	"zuoxingtao/init/db"
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

func (ud *UserDao) GetUsersByStatus(status int) ([]*user.User, error) {
	u := &user.User{}
	users := make([]*user.User, 0)
	if err := ud.Table(u.TableName()).Where("status = ?", status).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ud *UserDao) GetUserByMobile(mobile string) (*user.User, error) {
	user := &user.User{}
	fmt.Println(ud.DB)
	if err := ud.Table(user.TableName()).Where("mobile = ?", mobile).Find(user).Error; err != nil {
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
	if err := ud.Table(user.TableName()).Where("mobile = ?", user.Mobile).Update("status", user.Status).Error; err != nil {
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
