/**
 * @Author: lj
 * @Description: 自定义回答模板 将rule转为查询语句
 * @File:  answerTemplate
 * @Version: 1.0.0
 * @Date: 2022/03/14 15:48
 */

package ruleMatch

import "errors"

// AnswerTemplate
// Count int 共有多少个模板
// len(Templates) == Count
// Templates map[int]*Template 所有的回答模板
type AnswerTemplate struct {
	Count     int
	Templates map[int]string
}

func NewAnswerTemplate() (answerTemplate *AnswerTemplate) {
	answerTemplate = new(AnswerTemplate)
	answerTemplate.Templates = make(map[int]string, 0)
	return
}

func (ansTemp *AnswerTemplate) GetTemplate(answer *Answer) (string, error) {
	if ans, ok := ansTemp.Templates[answer.CategoryId]; !ok {
		return "", errors.New("没有对应模板")
	} else {
		return ans, nil
	}
}

func (ansTemp *AnswerTemplate) AddTemplate(id int, template string) {
	ansTemp.Templates[id] = template
	ansTemp.Count = len(ansTemp.Templates)
}

func (ansTemp *AnswerTemplate) RemoveTemplate(id int) {
	delete(ansTemp.Templates, id)
	ansTemp.Count = len(ansTemp.Templates)
}
