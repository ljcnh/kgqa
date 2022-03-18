/**
 * @Author: lj
 * @Description:
 * @File:  dictionary
 * @Version: 1.0.0
 * @Date: 2022/03/15 15:20
 */

package dictionary

type Dictionary struct {
	ID          int
	Name        string
	Description string
}

type Vocabulary struct {
	ID           int
	Vocabulary   string
	Pos          string
	DictionaryID int // Dictionary 主键
}
