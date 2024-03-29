package shop

import "time"

type Shop struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	Status       int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                           // 状态
	Deleted      int       `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`                         // 是否删除
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	Address      string    `gorm:"column:address" json:"address"`
	CityName     string    `gorm:"column:city_name" json:"city_name"`
	DistrictName string    `gorm:"column:district_name" json:"district_name"`
	FullAddress  string    `gorm:"column:full_address" json:"full_address"`
	Lng          float64   `gorm:"column:lng;default:0" json:"lng"`
	Name         string    `gorm:"column:name" json:"name"`
	ProvinceName string    `gorm:"column:province_name" json:"province_name"`
	ShopId       string    `gorm:"column:shop_id" json:"shopId"`
	Lat          float64   `gorm:"column:lat;default:0" json:"lat"`
}

func (shop *Shop) TableName() string {
	return "shop"
}
