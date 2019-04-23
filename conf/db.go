package conf

const DriverName = "mysql"

// 数据库连接信息
type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

var MasterDbConfig DbConfig = DbConfig{
	Host:     "localhost",
	Port:     "3306",
	Username: "root",
	Password: "joyiever",
	DbName:   "stars",
}

var SlvaeDbConfig DbConfig = DbConfig{
	Host:     "localhost",
	Port:     "3306",
	Username: "root",
	Password: "joyiever",
	DbName:   "stars",
}
