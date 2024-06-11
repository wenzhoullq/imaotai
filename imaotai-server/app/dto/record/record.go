package record

import "time"

type Record struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	Deleted    int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`                         // 是否删除
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	UserID     int       `gorm:"column:user_id;default:0" json:"user_id"`
	UserName   string    `gorm:"column:user_name" json:"user_name"`
	ItemID     string    `gorm:"column:item_id;default:0" json:"item_id"`
	ItemName   string    `gorm:"column:item_name" json:"item_name"`
}

func (rd *Record) TableName() string {
	return "record"
}
