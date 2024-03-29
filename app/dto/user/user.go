package user

import (
	"zuoxingtao/constant"
	"zuoxingtao/lib"
)

type User struct {
	ID           int     `json:"id" gorm:"id"`
	Status       int     `json:"status" gorm:"status"`
	Deleted      int     `json:"deleted" gorm:"deleted"`
	Mobile       string  `json:"mobile" gorm:"mobile"`
	Md5          string  `json:"md5" gorm:"md5"`
	DeviceID     string  `json:"device_id" gorm:"device_id"`
	Token        string  `json:"token" gorm:"token"`
	Lat          float64 `json:"lat" gorm:"lat"`
	Lng          float64 `json:"lng" gorm:"lng"`
	CityName     string  `json:"city_name" gorm:"city_name"`
	UserID       int     `json:"user_id" gorm:"user_id"`
	Source       int     `json:"source" gorm:"source"`
	UserName     string  `json:"user_name" gorm:"user_name"`
	ProvinceName string  `json:"province_name" gorm:"province_name"`
	DistrictName string  `json:"district_name" gorm:"district_name"`
	Cookie       string  `json:"cookie" gorm:"cookie"`
}

func (user *User) TableName() string {
	return "user"
}

func SetMobile(mobile string) func(user *User) {
	return func(user *User) {
		user.Mobile = mobile
		user.Md5 = lib.Signature(mobile)
	}
}

func (user *User) SetUser(ops ...func(user *User)) {
	for _, op := range ops {
		op(user)
	}
}

func NewUser(ops ...func(user *User)) *User {
	user := &User{
		Status:   constant.USER_INIT,
		DeviceID: lib.GetDeviceID(),
	}
	for _, op := range ops {
		op(user)
	}
	return user
}
