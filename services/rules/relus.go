package rules

import (
	"github.com/gomodule/redigo/redis"
	"kgqa/global"
)

func DeleteKG() {

}

// SetRule 设置知识图谱对应的规则和查询语句
// res  rule是否已经存在
func SetRule(kgName string, rule string, sql string) (ok bool, err error) {
	tmp, err := global.DbRedis.Do("hset", kgName, rule, sql)
	if err != nil {
		return
	}
	if tmp.(int64) == 1 {
		ok = true
	}
	return
}

// DeleteRule 删除KgName(知识图谱名称)中的rule
// res  false      true		false
// err  error      nil		nil
// 情况 过程报错	   删除成功	删除失败，不存在rule
func DeleteRule(kgName string, rule string) (ok bool, err error) {
	tmp, err := global.DbRedis.Do("hdel", kgName, rule)
	if err != nil {
		ok = false
		return
	}
	if tmp.(int64) == 1 {
		ok = true
	}
	return
}

// FindRule 查找kgName(知识图谱) rule中对应的查询语句
func FindRule(kgName string, rule string) (res string, err error) {
	tmp, err := global.DbRedis.Do("hget", kgName, rule)
	if err != nil {
		return
	}
	if tmp != nil {
		res = string(tmp.([]uint8))
	}
	return
}

func GetAllRules(kgName string) (all map[string]string, err error) {
	all, err = redis.StringMap(global.DbRedis.Do("hgetall", kgName))
	return
}
