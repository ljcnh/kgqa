package utils

import (
	"github.com/go-ego/gse"
)

type Participle struct {
	segment gse.Segmenter
}

func (seg *Participle) Initialization() {
	seg.segment.LoadDict()
}

// Cut HMM+ALL模式分词   包含重复元素
func (seg *Participle) Cut(text string) []string {
	text = RemoveSpaceAndTab(text)
	dag := seg.CutHmmSearch(text)
	cut := seg.CutAll(text)
	dag = append(dag, cut...)
	return seg.segment.Trim(dag)
}

// CutDeduplication HMM+ALL模式分词   不包含重复元素
func (seg *Participle) CutDeduplication(text string) []string {
	return RemoveRep(seg.Cut(text))
}

func (seg *Participle) CutHmmSearch(text string) []string {
	return seg.segment.CutSearch(text, true)
}

func (seg *Participle) CutAll(text string) []string {
	return seg.segment.CutAll(text)
}
