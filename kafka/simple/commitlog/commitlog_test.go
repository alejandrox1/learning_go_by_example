package commitlog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)


func check(t *testing.T, got []byte, want []byte) {
	if !bytes.Equal(got, want) {
		t.Errorf("got = %s, wanted = %s", string(got), string(want))
	}
}

func TestNewCommitLog(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("commitlogtest%d", rand.Int63()))
	fmt.Println(path)
	opts := Options{
		Path:         path,
		SegmentBytes: 3,
	}
	// Create a new CommitLog.
	l, err := New(opts)
	if err != nil {
		t.Fatal(err)
	}

	// Remove old data.
	l.deleteAll()
	// Recreate environment.
	l.init()
	// Create a new segment
	l.open()

	if _, err = l.Write([]byte("one")); err != nil {
		t.Error("Error writing to log: ", err)
	}

	if _, err = l.Write([]byte("two")); err != nil {
		t.Error("Error writing to log: ", err)
	}

	r, err := l.NewReader(0)
	if err != nil {
		t.Error(err)
	}
	p, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	check(t, p, []byte("onetwo"))
}


