package commitlog

import (
	"errors"
	"io"
	"sync"
)


type Reader struct {
	segment  *segment
	segments []*segment
	idx      int
	mu       sync.Mutex
	offset   int64
}


func (r *Reader) Read(m []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var readSize int
	for {
		readSize, err = r.segment.ReadAt(m[n:], r.offset)
		n += readSize
		if err != io.EOF {
			break
		}

		r.idx++
		if len(r.segments) <= r.idx {
			err = io.EOF
			break
		}

		r.segment = r.segments[r.idx]
		r.offset = 0
	}
	return
}


func (l *CommitLog) NewReader(offset int64) (*Reader, error) {
	segment, idx := findSegment(l.segments, offset)
	if segment == nil {
		return nil, errors.New("NewReader did not find segment")
	}

	return &Reader{
		segment:  segment,
		segments: l.segments,
		idx:      idx,
		offset:   offset,
	}, nil
}
