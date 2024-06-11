package item

import "time"

type Item struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	Status     int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态
	Deleted    int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`                         // 是否删除
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	ItemCode   string    `gorm:"column:item_code" json:"item_code"`                                        // 酒编码
	Title      string    `gorm:"column:title" json:"title"`
}

func (item *Item) TableName() string {
	return "item"
}
