package record

type Record struct {
	ID       int    `json:"id" gorm:"id"`
	Deleted  int    `json:"deleted" gorm:"deleted"`
	UserID   int    `json:"user_id" gorm:"user_id"`
	UserName string `json:"user_name" gorm:"user_name"`
	ItemID   string `json:"item_id" gorm:"item_id"`
	ItemName string `json:"item_name" gorm:"item_name"`
}

func (rd *Record) TableName() string {
	return "record"
}
