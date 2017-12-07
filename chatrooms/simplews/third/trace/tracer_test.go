package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Tracer should not be nill")
	}

	tracer.Trace("Hello trace pkg")
	if buf.String() != "Hello trace pkg\n" {
		t.Errorf("Trace should not write '%s'", buf.String())
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
