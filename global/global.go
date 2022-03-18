package global

import (
	"kgqa/config"
	"kgqa/models/redis"
	_ "kgqa/models/redis"
	"kgqa/utils"
	"kgqa/utils/intentDetection/ruleMatch"
)

// Config 配置初始化
var Config = new(config.Config)

// RedisConfig 配置初始化
var RedisConfig = new(config.RDConfig)

// DbRedis Redis 初始化
var DbRedis redis.RedisPool

func init() {
	config.InitConfig(Config)
	RedisConfig = &Config.RedisConfig
	DbRedis = redis.PollInit(RedisConfig)
}

// ————————————————————————————————————————————————————————————————————————————————————————————————————

// Seg 分词初始化
var Seg utils.Participle

func init() {
	//Seg.Initialization()
}

// ————————————————————————————————————————————————————————————————————————————————————————————————————

// Templates 自定义模板 能将rule转为查询语句
// 这些全部需要自定义
var Templates *ruleMatch.AnswerTemplate

func init() {
	Templates = ruleMatch.NewAnswerTemplate()
	Templates.AddTemplate(1, `MATCH (v:{{.BelongEntity}} { {{.PropertyName}}:"{{.VarValue}}"}) RETURN v.{{.AnsPropertyName}};`)
	Templates.AddTemplate(2, `MATCH (v:{{.BelongEntity}} { {{.PropertyName}}:"{{.VarValue}}"}) -[e:{{.RelationShip}}]-(v2) RETURN v2.{{.AnsPropertyName}};`)
}
