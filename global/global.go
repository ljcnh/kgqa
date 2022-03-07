package global

import (
	"kgqa/config"
	"kgqa/models/redis"
	_ "kgqa/models/redis"
	"kgqa/utils"
)

// Config 配置初始化

var Config = new(config.Config)

// RedisConfig 配置初始化  Redis 初始化

var RedisConfig = new(config.RDConfig)
var DbRedis redis.RedisPool

func init() {
	config.Init(Config)
	RedisConfig = &Config.RedisConfig
	DbRedis = redis.PollInit(RedisConfig)
}

// Seg 分词初始化
var Seg utils.Participle

func init() {
	Seg.Initialization()
}
