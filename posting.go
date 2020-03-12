package wiser

type PostingList struct {
	DocumentID     int
	Positions      Positions
	PositionsCount int
	next           *PostingList
}
type Positions struct {
	list []string
}
