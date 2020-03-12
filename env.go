package wiser

type Env struct {
	dbPath            string
	tokenLength       uint
	CompressionMethod CompressionMethod

	invertedIndex                      *InvertedIndex
	invertedIndexUpdateBufferThreshold int

	enablePhraseSearch bool
}

func NewEnv(
	invIndexBufferThreshold int,
	enablePhraseSearch bool,
	dbPath string,
) (*Env, error) {
	return &Env{
		dbPath:                             dbPath,
		enablePhraseSearch:                 enablePhraseSearch,
		invertedIndex:                      NewInvertedIndex(),
		invertedIndexUpdateBufferThreshold: invIndexBufferThreshold,
	}, nil
}

func (e *Env) Close() error {
	return nil
}

func (e *Env) Search(query string) error {
	return nil
}

func (e *Env) Index(
	compressionMethod CompressionMethod,
) error {
	return nil
}
