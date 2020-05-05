package conf

type AppConfig struct {
	MysqlConfig  `ini:"mysql"`
	RedisConfig  `ini:"redis"`
	LoggerConfig `ini:"logger"`
}

type MysqlConfig struct {
	UserName  string `ini:"userName"`
	Password  string `ini:"password"`
	IpAddress string `ini:"ipAddress"`
	Port      int    `ini:"port"`
	DbName    string `ini:"dbName"`
	Charset   string `ini:"charset"`
}

type RedisConfig struct {
	IpAddress string `ini:"ipAddress"`
	Port      int    `ini:"port"`
}

type LoggerConfig struct {
	FileName string `ini:"fileName"`
	FilePath string `ini:"filePath"`
}
