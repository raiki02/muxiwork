package conf

type UserConf struct {
	DatabaseConf `ini:"database"`
}

type DatabaseConf struct {
	Username     string `ini:"username"`
	Password     string `ini:"password"`
	DatabaseName string `ini:"databasename"`
}
