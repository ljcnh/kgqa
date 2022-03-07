package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	WebName     string   `mapstructure:"webName"`
	LogLevel    string   `mapstructure:"webName"`
	RedisConfig RDConfig `mapstructure:"redis"`
}

type RDConfig struct {
	Ip   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
	//user = "none"
	Password    string `mapstructure:"password"`
	MaxIdle     int    `mapstructure:"maxIdle"`
	MaxActive   int    `mapstructure:"maxActive"`
	IdleTimeout int    `mapstructure:"idleTimeout"`
	DefaultDB   int    `mapstructure:"defaultDB"`
}

func Init(config *Config) {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("解析配置文件失败 错误:%s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已被修改")
		if err := viper.Unmarshal(config); err != nil {
			log.Fatalf("解析配置文件失败 错误:%s \n", err)
		}
	})
}
