package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of tracing events
// thorugh code.
type Tracer interface {
	Trace(...interface{})
}



// tracer is aTracer that writes to an io.Writer.
type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New creates a new Tracer that will write the output to a specified
// io.Writer.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}


type nilTracer struct{}

func (n *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
