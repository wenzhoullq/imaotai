package item

type Item struct {
	ID       int    `json:"id" gorm:"id"`
	Status   int    `json:"status" gorm:"status"`
	Deleted  int    `json:"deleted" gorm:"deleted"`
	ItemCode string `json:"item_code" gorm:"item_code"`
	Title    string `json:"title" gorm:"title"`
}

func (item *Item) TableName() string {
	return "item"
}
