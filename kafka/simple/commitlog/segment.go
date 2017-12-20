package commitlog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)


const (
	logNameFormat = "%020d.log"
	indexNameFormat = "%020d.index"
)

type segment struct {
	writer       io.Writer
	reader       io.Reader
	log          *os.File
	index        *os.File
	baseOffset   int64
	newestOffset int64
	bytes        int64
	maxBytes     int64
}

func (s *segment) Error() string {
	return "segment struct error"
}

func NewSegment(path string, baseOffset int64, maxBytes int64) (*segment, error) {
	logPath := filepath.Join(path, fmt.Sprintf(logNameFormat, baseOffset))
	log, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	file, err := log.Stat()
	if err != nil {
		return nil, err
	}

	indexPath := filepath.Join(path, fmt.Sprintf(indexNameFormat, baseOffset))
	index, err := os.OpenFile(indexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &segment{
		writer:       log,
		reader:       log,
		log:          log,
		index:        index,
		baseOffset:   baseOffset,
		newestOffset: baseOffset,
		bytes:        file.Size(),
		maxBytes:     maxBytes,
	}, nil
}


func (s *segment) NewestOffset() int64 {
	return s.newestOffset
}


func (s *segment) IsFull() bool {
	return s.bytes >= s.maxBytes
}


func (s *segment) Write(m []byte) (n int, err error) {
	n, err = s.writer.Write(m)
	if err!= nil {
		return n, err
	}

	_, err = s.index.Write([]byte(fmt.Sprintf("%d,%d\n", s.newestOffset, s.bytes)))
	if err != nil {
		return 0, err
	}

	s.newestOffset += 1
	s.bytes += int64(n)
	return
}


func (s *segment) Read(m []byte) (n int, err error) {
	return s.reader.Read(m)
}


func (s *segment) ReadAt(m []byte, offset int64) (int, error) {
	return s.log.ReadAt(m, offset)
}


func (s *segment) Close() error {
	if err := s.log.Close(); err != nil {
		return err
	}
	return s.index.Close()
}
