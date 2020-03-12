package wiser

type CompressionMethod int

const (
	CompressionMethodNone CompressionMethod = iota
	CompressionMethodGolomb
)
