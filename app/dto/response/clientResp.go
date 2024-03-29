package response

type ImaotaiResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type MTVersionResp struct {
	ImaotaiResp
}

type BaiduParseIPResp struct {
	Status  int `json:"status"`
	Content struct {
		Address string `json:"address"`
		Point   struct {
			Lng string `json:"x"`
			Lat string `json:"y"`
		} `json:"point"`
	} `json:"content"`
}

type BaiduParseAddressResp struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
	} `json:"result"`
}

type BaiduParseLngAndLatsResp struct {
	Status int `json:"status"`
	Result struct {
		AddressComponent struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
		} `json:"addressComponent"`
	} `json:"result"`
}

type MTLogInResp struct {
	ImaotaiResp
	Data struct {
		UserID   int    `json:"userId" `
		UserName string `json:"userName" `
		Token    string `json:"token"`
		Cookie   string `json:"cookie"`
	} `json:"data"`
}

type ShopListResp struct {
	ImaotaiResp
	Data struct {
		MtshopsPc struct {
			Md5     string `json:"md5"`
			Size    int    `json:"size"`
			Url     string `json:"url"`
			Version int    `json:"version"`
		} `json:"mtshops_pc"`
	} `json:"data"`
}

type SessionResp struct {
	ImaotaiResp
	Data struct {
		ItemList  []*Item `json:"itemList"`
		SessionID int     `json:"sessionId"`
	} `json:"Data"`
}
type Item struct {
	Content   string `json:"content"`
	ItemCode  string `json:"itemCode"`
	JumpURL   string `json:"jumpUrl"`
	Picture   string `json:"picture"`
	PictureV2 string `json:"pictureV2"`
	Title     string `json:"title"`
}

type ShopByProvinceResp struct {
	ImaotaiResp
	Data struct {
		Shops []struct {
			ShopID string `json:"shopId"`
		} `json:"shops"`
	} `json:"Data"`
}

type ReserveResp struct {
	ImaotaiResp
}

type RecordResp struct {
	ImaotaiResp
	Data struct {
		ReservationItemVOS []struct {
			Status   int    `json:"status"`
			ItemID   string `json:"itemId"`
			ItemName string `json:"itemName"`
		} `json:"reservationItemVOS"`
	} `json:"data"`
}
