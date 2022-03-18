/**
 * @Author: lj
 * @Description:
 * @File:  ruleStruct
 * @Version: 1.0.0
 * @Date: 2022/03/09 17:59
 */

package ruleMatch

type Entity = string
type Property = string
type Category = int
type Relation = string

// GraphNode   知识图谱中的基本属性  实体和属性
//	BelongEntity Entity   所属实体
//  PropertyName Property 在实体中的属性
type GraphNode struct {
	BelongEntity Entity
	PropertyName Property
}

// ----------------------------------------------------------------------------

// VarDefinition 定义变量
//	VarName      string   变量的名字
//  GraphNode
type VarDefinition struct {
	VarName  string
	VarValue string
	*GraphNode
}

// ----------------------------------------------------------------------------

// AnswerProperty
// AnsPropertyName Property 知识图谱中对应属性名称
type AnswerProperty struct {
	*VarDefinition
	AnsPropertyName Property
}

// AnswerByRelation
// RelationShip Relation 关系名称
// IsPositive   uint8     流入还是流出
// AnsPropertyName Property 需要返回的属性
type AnswerByRelation struct {
	*VarDefinition
	RelationShip    Relation
	IsPositive      uint8
	AnsPropertyName Property
}

// ----------------------------------------------------------------------------

// VarDefinitions []*VarDefinition 删除
// 原本用来存放所有的变量及定义
// 但是对答案无用的变量有用吗？ 不应存在这种情况

// Answer
// SpaceID int 图空间的id
// CategoryId  Category 决定使用模板类别 两种不同的问答方式
// 目前只包含普通问答：
// Category
//	   1  : AnswerProperty
// 	   2  : AnswerByRelation
type Answer struct {
	SpaceID    int
	CategoryId Category
	AnsPro     *AnswerProperty
	AnsByRel   *AnswerByRelation
}

//  Category 的判断  需要到时候做出规定

func (ans *Answer) AnswerSelect() interface{} {
	switch ans.CategoryId {
	case 1:
		return ans.AnsPro
	case 2:
		return ans.AnsByRel
	}
	return nil
}
