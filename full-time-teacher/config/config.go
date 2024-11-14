package config

import "github.com/spf13/viper"

// Config 结构体表示应用程序的配置
type Config struct {
	AppName    string `mapstructure:"app_name"`
	AppVersion string `mapstructure:"app_version"`
	ServerPort int    `mapstructure:"server_port"`
	MySQL      MySQLConfig
}

// MySQLConfig 结构体表示 MySQL 的配置
type MySQLConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
}

func LoggingConfig() (*Config, error) {

	configFilePath := "../config/config.toml"

	// 使用viper读取配置文件
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// 将配置文件绑定结构体
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
