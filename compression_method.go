package wiser

type CompressionMethod int

const (
	CompressionMethodNone CompressionMethod = iota
	CompressionMethodGolomb
)

func NewCompressionMethod(m string) CompressionMethod {
	return CompressionMethodNone
}
