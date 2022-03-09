package config

type GlobalConfig struct {
	Name        string
	Port        int
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
	LogsAddr    string
	ImagePath   string
}

type MysqlConfig struct {
	Host     string
	Port     int
	Name     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}
