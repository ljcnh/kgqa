/**
 * @Author: lj
 * @Description: 自定义规则转查询语句
 * @File:  ruleToNGQL
 * @Version: 1.0.0
 * @Date: 2022/03/08 10:06
 */

/*
问答的形式(由易至难)
	1. X 的 值 是什么？
		X:实体  值: X的属性或者与其他实体的关系

		属性都是正向
		关系可能分为正反向
			还需要考虑是否被动（流入流出）
				某个老师教什么课程？
				什么课程由什么老师教？
			所以还需要记录X和 值 之间的关系是正向还是反向

    2. X 的 Y 是什么？
		X:实体  Y: X的属性或者与其他实体的关系

  	3. X 和 Y 的 Z(关系)是什么？
		暂不考虑，先完成最基础的方式
*/

package ruleMatch

import (
	"bytes"
	"text/template"
)

func (ansTemp *AnswerTemplate) RuleToNGQL(ans *Answer) (gql string, err error) {
	var s bytes.Buffer

	temp, _ := ansTemp.GetTemplate(ans)
	tmpl, _ := template.New("").Parse(temp)
	err = tmpl.Execute(&s, ans.AnswerSelect())

	gql = s.String()
	return
}
