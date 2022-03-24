package config

// GetDevGlobalConfig 获取全局配置信息
func GetDevGlobalConfig() GlobalConfig {
	return GlobalConfig{
		Name: "xupt-scholarship",
		Port: 8096,
		MysqlConfig: MysqlConfig{
			Host:     "127.0.0.1",
			Port:     3308,
			Name:     "root",
			Password: "586014BZYbzy",
			DBName:   "xupt-scholarship",
		},
		RedisConfig: RedisConfig{
			Host: "127.0.0.1",
			Port: 6379,
		},
		LogsAddr:   "./logs",
		ImagePath:  "./images/",
		FilePath:   "./uploads/",
		AvatarPath: "./avatars/",
	}
}
