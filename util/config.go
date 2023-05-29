package util

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"sslmode"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	DB     DBConfig     `mapstructure:"crud_db"`
	Server ServerConfig `mapstructure:"crud_server"`
}

var vp *viper.Viper

func LoadConfig(path string) (config Config, err error) {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(path)
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		return Config{}, err
	}
	if err = vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
