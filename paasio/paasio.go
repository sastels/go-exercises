package paasio

import (
	"io"
)

type PaasReader struct {
	reader io.Reader
	n      *int64
	nops   *int
}

func (a PaasReader) ReadCount() (int64, int) {
	return *a.n, *a.nops
}

func (a PaasReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	*a.n += int64(n)
	*a.nops++
	return n, err
}

func NewReadCounter(reader io.Reader) ReadCounter {
	n := int64(0)
	nops := 0
	return PaasReader{reader, &n, &nops}
}
