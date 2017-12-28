// see https://godoc.org/gonum.org/v1/gonum/floats
package main

import (
	"math"
)

type valstream chan float64

// Create a sequence with the given input value.
func Val(x float64) valstream {
	vs := make(valstream)
	go func() {
		defer close(vs)
		for i:=0; i<10; i++ {
			vs <-x
		}
	}()
	return vs
}

// Use two channels as tuples of input to a function and send values to a
// valstream channel.
func ApplyTo(fn func(x, y float64) float64, xs, ys valstream) valstream {
	vs := make(valstream)
	go func() {
		defer close(vs)
		for x := range xs {
			y := <-ys
			vs <- fn(x,y)
		}
	}()
	return vs
}

// Add two valstreams of floats element-wise.
func Add(x, y valstream) valstream {
	return ApplyTo(func(x, y float64) float64 { return x+y }, x, y)
}

// Subtract to valstreams of floas element-wise.
func Sub(x, y valstream) valstream {
	return ApplyTo(func(x, y float64) float64 { return x-y }, x, y)
}

// Multiply two valstreams of floats element-wise.
func Mul(x, y valstream) valstream {
	return ApplyTo(func(x, y float64) float64 { return x*y }, x, y)
}

// Divide two valstreams of floats element-wise.
func Div(x, y valstream) valstream {
	return ApplyTo(func(x, y float64) float64 { return x/y }, x, y)
}

// Exponentiate element-wise x to the power y.
func Exp(x, y valstream) valstream {
	return ApplyTo(func(x, y float64) float64 { return math.Pow(x, y) }, x, y)
}
