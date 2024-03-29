package shop

type Shop struct {
	ID           int     `json:"id" gorm:"id"`
	Status       int     `json:"status" gorm:"status"`
	Address      string  `json:"address" gorm:"address"`
	CityName     string  `json:"cityName" gorm:"city_name"`
	DistrictName string  `json:"districtName" gorm:"district_name"`
	FullAddress  string  `json:"fullAddress" gorm:"full_address"`
	Lat          float64 `json:"lat" gorm:"lat"`
	Lng          float64 `json:"lng" gorm:"lng"`
	Name         string  `json:"name" gorm:"name"`
	ProvinceName string  `json:"provinceName" gorm:"province_name"`
	ShopId       string  `json:"shopId" gorm:"shop_id"`
}

func (shop *Shop) TableName() string {
	return "shop"
}
