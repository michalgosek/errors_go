package stream

type Reader interface {
	Read(bb ...byte) error
}

// Idea: https://cs.opensource.google/go/go/+/refs/tags/go1.21.6:src/bufio/bufio.go;l=676
type ErrStreamReader struct {
	err error
	sr  Reader
}

func (e *ErrStreamReader) Read(bb ...byte) {
	if e.err != nil {
		return
	}
	e.err = e.sr.Read(bb...)
}

func (e *ErrStreamReader) Error() error {
	return e.err
}

func NewErrStreamReader() *ErrStreamReader {
	sr := ErrStreamReader{sr: &Stream{}}
	return &sr
}
