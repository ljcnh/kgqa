package main

//  Category 的判断  需要到时候做出规定

func main() {
	/*	answer := ruleMatch.Answer{
			SpaceID:    0,
			CategoryId: 1,
			AnsPro: &ruleMatch.AnswerProperty{
				VarDefinition: &ruleMatch.VarDefinition{
					VarName:  "X",
					VarValue: "C++程序设计实践",
					GraphNode: &ruleMatch.GraphNode{
						BelongEntity: "course",
						PropertyName: "name",
					},
				},
				AnsPropertyName: "created_at",
			},
			AnsByRel: nil,
		}
		gql, _ := global.Templates.RuleToNGQL(&answer)
		fmt.Println(gql)
		answer1 := ruleMatch.Answer{
			SpaceID:    0,
			CategoryId: 2,
			AnsPro:     nil,
			AnsByRel: &ruleMatch.AnswerByRelation{
				VarDefinition: &ruleMatch.VarDefinition{
					VarName:  "X",
					VarValue: "C++程序设计实践",
					GraphNode: &ruleMatch.GraphNode{
						BelongEntity: "course",
						PropertyName: "name",
					},
				},
				RelationShip:    "course_exercise",
				IsPositive:      0,
				AnsPropertyName: "publish_time",
			},
		}
		gql, _ = global.Templates.RuleToNGQL(&answer1)
		fmt.Println(gql)*/
	//rule := "<?x>的<?阿斯顿撒旦>[联系方式|电话](是[什么|啥](啊)?<?阿斯顿撒旦>)?"
	//newRule := ruleMatch.Preprocessing(rule)
	//fmt.Println(newRule)
	//rule = "<?x>[联系方式<?sca>|电话<?sdsd速度>](是[什么|啥](啊)?)?"
	//newRule = ruleMatch.Preprocessing(rule)
	//fmt.Println(newRule)

	//rule := `<!x>的[联系方式|电话](是[什么|啥](啊)?)?`
	//text := "罗老师的联系方式是啥"
	//fmt.Println(ruleMatch.RegularMatch(rule, text))
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
