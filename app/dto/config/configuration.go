package config

type Configuration struct {
	ServerConfigure   `toml:"server_configure"`
	BaiduMapConfigure `toml:"baidu_map_configure"`
	MysqlConfigure    `toml:"mysql_configure"`
	LogConfigure      `toml:"log_configure"`
	CronTaskConfigure `toml:"corn_task_configure"`
}

type LogConfigure struct {
	LogFilePath string `toml:"log_file_path"`
	LogFileName string `toml:"log_file_name"`
}

type ServerConfigure struct {
	ServerAddr string `toml:"server_addr"`
}

type BaiduMapConfigure struct {
	AK string `toml:"ak"`
}

type MysqlConfigure struct {
	Driver   string `toml:"driver"`
	UserName string `toml:"user_name"`
	Pw       string `toml:"pw"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DbName   string `toml:"dbname"`
	TimeOut  string `toml:"timeout"`
}

type CronTaskConfigure struct {
	UpdateShop   string `toml:"update_shop"`
	UpdateItem   string `toml:"update_item"`
	ExUser       string `toml:"ex_user"`
	UpdateCommon string `toml:"update_common"`
	Reservation  string `toml:"reservation"`
	AddRecord    string `toml:"add_record"`
}
