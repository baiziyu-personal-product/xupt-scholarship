package config

type ServerConfig struct {
	Name        string
	Port        int
	MsqlConfig  MysqlConfig
	RedisConfig RedisConfig
	LogsAddr    string
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
	Port     string
	Password string
}
