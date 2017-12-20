package commitlog

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type CommitLog struct {
	Options
	name           string
	mu             sync.RWMutex
	segments       []*segment
	vActiveSegment atomic.Value
}

type Options struct {
	Path         string
	// Used to indicate the maxBytes for segment structs.
	SegmentBytes int64
}


// Create an instance of COmmitLog by creating the path dir and using method
// open().
func New(opts Options) (*CommitLog, error) {
	if opts.Path == "" {
		return nil, errors.New("Options struct path is empty")
	}

	if opts.SegmentBytes == 0 {
		// TODO default
	}

	path, err := filepath.Abs(opts.Path)
	if err != nil {
		return nil, err
	}

	l := &CommitLog{
		Options: opts,
		name:    filepath.Base(path),
	}
	return l, nil
}

// Create the appropiate directories where to store logs.
func (l *CommitLog) init() error {
	return os.MkdirAll(l.Path, 0755)
}

// Check path exists and create a new segment.
func (l *CommitLog) open() error {
	if _, err := ioutil.ReadDir(l.Path); err != nil {
		return err
	}

	activeSegment, err := NewSegment(l.Path, 0, l.SegmentBytes)
	if err != nil {
		return err
	}
	l.vActiveSegment.Store(activeSegment)
	l.segments = append(l.segments, activeSegment)
	return nil
}


func (l *CommitLog) activeSegment() *segment {
	return l.vActiveSegment.Load().(*segment)
}


func (l *CommitLog) checkSplit() bool {
	return l.activeSegment().IsFull()
}


func (l *CommitLog) newestOffset() int64 {
	return l.activeSegment().NewestOffset()
}


func (l *CommitLog) split() error {
	seg, err := NewSegment(l.Path, l.newestOffset(), l.SegmentBytes)
	if err != nil {
		return err
	}
	l.segments = append(l.segments, seg)
	l.vActiveSegment.Store(seg)
	return nil
}


func (l *CommitLog) Read(m []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.activeSegment().Read(m)
}


func (l *CommitLog) Write(m []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.checkSplit() {
		if err := l.split(); err != nil {
			return 0, err
		}
	}
	return l.activeSegment().Write(m)
}

// Remove data.
func (l *CommitLog) deleteAll() error {
	return os.RemoveAll(l.Path)
}
