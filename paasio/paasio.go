package paasio

import (
	"io"
	"sync"
)

type PaasReader struct {
	mutex  *sync.Mutex
	reader io.Reader
	n      *int64
	nops   *int
}

func (a PaasReader) ReadCount() (int64, int) {
	return *a.n, *a.nops
}

func (a PaasReader) Read(p []byte) (int, error) {
	a.mutex.Lock()
	n, err := a.reader.Read(p)
	*a.n += int64(n)
	*a.nops++
	a.mutex.Unlock()
	return n, err
}

func NewReadCounter(reader io.Reader) ReadCounter {
	n := int64(0)
	nops := 0
	return PaasReader{&(sync.Mutex{}), reader, &n, &nops}
}
