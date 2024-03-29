package record

import "time"

type Record struct {
	ID         int       `json:"id" gorm:"column:id;not null;primary_key;AUTO_INCREMENT;type:int(11)"`
	Deleted    int       `json:"deleted" gorm:"column:deleted;type:tinyint(4)"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;not null;default CURRENT_TIMESTAMP;type:datetime"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;not null;default CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP;type:datetime"`
	UserID     int       `json:"user_id" gorm:"column:user_id;type:int(11)"`
	UserName   string    `json:"user_name" gorm:"column:user_name;type:varchar(20)"`
	ItemID     string    `json:"item_id" gorm:"column:item_id;type:varchar(50)"`
	ItemName   string    `json:"item_name" gorm:"column:item_name;type:varchar(100)"`
}

func (rd *Record) TableName() string {
	return "record"
}
