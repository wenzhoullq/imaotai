package user

import (
	"imaotai_helper/constant"
	"imaotai_helper/lib"
	"time"
)

type User struct {
	ID           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	UID          string    `gorm:"column:uid;size:32;NOT NULL" json:"uid"`                                   // 自己的UID
	AdminUID     string    `gorm:"column:admin_uid;size:32;NOT NULL" json:"admin_uid"`                       // 管理员的UID
	Mobile       string    `gorm:"column:mobile;NOT NULL" json:"mobile"`                                     // 手机号
	Status       int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态
	Deleted      int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`                         // 是否删除
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	Md5          string    `gorm:"column:md5" json:"md5"`                                                    // md5
	DeviceID     string    `gorm:"column:device_id" json:"device_id"`                                        // 设备名称
	Token        string    `gorm:"column:token" json:"token"`
	Lat          float64   `gorm:"column:lat" json:"lat"`
	Lng          float64   `gorm:"column:lng" json:"lng"`
	CityName     string    `gorm:"column:city_name" json:"city_name"`
	UserID       int       `gorm:"column:user_id" json:"user_id"` //imaotai用户ID
	Source       int       `gorm:"column:source" json:"source"`
	UserName     string    `gorm:"column:user_name" json:"user_name"`
	ProvinceName string    `gorm:"column:province_name" json:"province_name"`
	DistrictName string    `gorm:"column:district_name" json:"district_name"`
	Cookie       string    `gorm:"column:cookie" json:"cookie"`
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

func SetUID(uid string) func(user *User) {
	return func(user *User) {
		user.UID = uid
	}
}

func (user *User) SetUser(ops ...func(user *User)) {
	for _, op := range ops {
		op(user)
	}
}
func SetAdminUID(adminUID string) func(user *User) {
	return func(user *User) {
		user.AdminUID = adminUID
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
