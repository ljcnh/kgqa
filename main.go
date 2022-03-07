package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	_ "kgqa/global"
	"reflect"
	"strconv"
)

const pi float32 = 89.89877
const zero = 0.0

func main() {
	a := 0.0
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())
	fmt.Println(reflect.TypeOf(zero), reflect.ValueOf(zero).Kind())

	//fmt.Println(rules.GetAllRules("key"))
	//fmt.Println(rules.GetAllRules("xcxc"))

	//rules.GetAllRules("key")
	//rules.GetAllRules("xcxc")
	//fmt.Println(global.Seg.CutDeduplication("gse 是 结巴 分词 jieba 的 golang 实现 并 尝试 添加 nlp 功能 和 更多 属性 g s e j i b a o l n p"))
}

//func main() {
//pool := &redis.Pool{
//	MaxActive:   6,
//	MaxIdle:     5,
//	IdleTimeout: time.Second * 100,
//	Wait:        true,
//	Dial: func() (redis.Conn, error) {
//		conn, err := redis.Dial("tcp", "82.156.183.226:6379")
//		if err != nil {
//			return nil, err
//		}
//		if _, err := conn.Do("AUTH", "ljwudao5"); err != nil {
//			conn.Close()
//			return nil, err
//		}
//		return conn, err
//	},
//}
//defer pool.Close()
//for i := 0; i < 10; i++ {
//	go getConnFromPoolAndHappy(pool, i)
//}
//time.Sleep(3 * time.Second)
//}
func getConnFromPoolAndHappy(pool *redis.Pool, i int) {
	//通过连接池获得连接
	conn := pool.Get()
	//延时关闭连接
	defer conn.Close()
	//使用连接操作数据
	reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
	if err != nil {
		fmt.Println(err)
	}

	s, _ := redis.String(reply, err)
	fmt.Println(s)
}
