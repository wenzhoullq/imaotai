package item

import "time"

type Item struct {
	ID         int       `json:"id" gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)"`
	Status     int       `json:"status" gorm:"column:status;type:tinyint(4)"`
	Deleted    int       `json:"deleted" gorm:"column:deleted;type:tinyint(4)"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null;default CURRENT_TIMESTAMP;type:timestamp"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;not null;default CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP;type:timestamp"`
	ItemCode   string    `json:"item_code" gorm:"column:item_code;type:varchar(20)"`
	Title      string    `json:"title" gorm:"column:title;type:varchar(20)"`
}

func (item *Item) TableName() string {
	return "item"
}
