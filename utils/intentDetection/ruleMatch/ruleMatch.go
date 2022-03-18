package ruleMatch

import (
	"regexp"
)

// 规则匹配
// 使用 <?变量> 匹配   匹配前需要替换 <?变量>
var ruleVariable = `<\?.*?>`
var replace = ".*"

func RegularMatch(rule string, text string) (ok bool, err error) {
	rule = Preprocessing(rule)
	ok, err = regexp.MatchString(rule, text)
	return
}

func Preprocessing(rule string) (newRule string) {
	variableRe := regexp.MustCompile(ruleVariable)
	newRule = variableRe.ReplaceAllString(rule, replace)
	return
}
