package paasio

import (
	"io"
	"sync"
)

func NewReadCounter(reader io.Reader) ReadCounter {
	n := int64(0)
	nops := 0
	return PaasReader{&(sync.Mutex{}), reader, &n, &nops}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	n := int64(0)
	nops := 0
	return PaasWriter{&(sync.Mutex{}), writer, &n, &nops}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return PaasReaderWriter{NewReadCounter(rw), NewWriteCounter(rw)}
}

type PaasReader struct {
	mutex  *sync.Mutex
	reader io.Reader
	n      *int64
	nops   *int
}

func (a PaasReader) ReadCount() (int64, int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return *a.n, *a.nops
}

func (a PaasReader) Read(p []byte) (int, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	n, err := a.reader.Read(p)
	*a.n += int64(n)
	*a.nops++
	return n, err
}

type PaasWriter struct {
	mutex  *sync.Mutex
	writer io.Writer
	n      *int64
	nops   *int
}

func (a PaasWriter) WriteCount() (int64, int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return *a.n, *a.nops
}

func (a PaasWriter) Write(p []byte) (int, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	n, err := a.writer.Write(p)
	*a.n += int64(n)
	*a.nops++
	return n, err
}

type PaasReaderWriter struct {
	reader ReadCounter
	writer WriteCounter
}

func (a PaasReaderWriter) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	return n, err
}

func (a PaasReaderWriter) Write(p []byte) (int, error) {
	n, err := a.writer.Write(p)
	return n, err
}

func (a PaasReaderWriter) ReadCount() (int64, int) {
	return a.reader.ReadCount()
}

func (a PaasReaderWriter) WriteCount() (int64, int) {
	return a.writer.WriteCount()
}
