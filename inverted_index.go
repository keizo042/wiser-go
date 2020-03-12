package wiser

type InvertedIndex struct {
	TokenID        int
	DocsCount      int
	PositionsCount int
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{}
}
