package admin

import "time"

type Admin struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	UID        string    `gorm:"column:uid;size:32;NOT NULL" json:"uid"`                                   // 用户ID
	Mobile     string    `gorm:"column:mobile;size:16;NOT NULL" json:"mobile"`                             // 手机号,等同于用户名
	PassWord   string    `gorm:"column:password;size:64;NOT NULL" json:"pass_word"`                        // 密码
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态
	Role       int       `gorm:"column:role;default:0;NOT NULL" json:"role"`                               // 角色
	Email      string    `gorm:"column:email;size:32;NOT NULL" json:"email"`                               // 邮箱
	Deleted    int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`                         // 是否删除
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
}

func (Admin *Admin) TableName() string {
	return "admin"
}
