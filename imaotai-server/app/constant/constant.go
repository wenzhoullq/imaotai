package constant

const (
	SALT = "2af72f100c356273d46284f6fd1dfc08" //盐值
)

// clent
const (
	CODESUCCESS     = 2000
	BAIDUMAPSUCCESS = 0
)

const (
	USER_INIT      = iota + 1 //初始化
	USER_NORMAL               //正常
	User_SUSPENDED            //暂停
	USER_ABNORMAL             //异常
)

const (
	USER_NOTEXIST   = "用户不存在"           //初始化
	USER_PORCESSING = "用户正常"            //正常
	USER_TOKENEX    = "用户已过期,重新更新Token" //异常
)

const (
	NORMAL  = 0 //未删除
	Deleted = 1 //删除
)

const (
	ADMIN_NORMAL  = 1
	ADMIN_DELETED = 2
)

const (
	Success = iota
	ParamErr
	ServerErr
	DBErr
	RPCErr
)

const (
	SHOP_OPEN  = 1
	SHOP_CLOSE = 2
)

const (
	SuperAdmin  = 1
	NorMalAdmin = 2
)

const (
	ITEM_OPEN  = 1
	ITEM_CLOSE = 2
)

const (
	AESKEY = "qbhajinldepmucsonaaaccgypwuvcjaa"
	AESIV  = "2018534749963515"
)

const (
	AWARD = 0
)

const (
	TRAVEL_STATUS_FREE       = 1
	TRAVEL_STATUS_PROCESSING = 2
	TRAVEL_STATUS_FINISH     = 3
)

const (
	TRAVEL_CONSUME = 100
)

const (
	SESSION_TYPE_NORMAL_ORDER = 1
)

const (
	JWT_SALT = "2af72f100c356273d46284f6fd1dfc"
)

const (
	HEADER_JWT = "Authorization"
)
