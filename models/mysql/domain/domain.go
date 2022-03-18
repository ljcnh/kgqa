/**
 * @Author: lj
 * @Description: 领域crud 包括词典映射 规则
 * @File:  domain
 * @Version: 1.0.0
 * @Date: 2022/03/15 11:00
 */

package domain

type Domain struct {
	ID          int
	Name        string
	Description string
}

type DomToDictionary struct {
	DictionaryID int // 外键id
	DomainId     int // 外键id
}

type Rule struct {
	ID          int
	Rule        string
	GQL         string
	Description string
	DomainId    int // 外键id
}
