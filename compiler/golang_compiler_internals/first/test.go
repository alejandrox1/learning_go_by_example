// go tool compile test.go
package main

type I interface {
	DoSomeWork()
}

type T struct {
	a int
}

func (t *T) DoSomeWork() {
}

func main() {
	t := &T{}
	i := I(t)
	print(i)
}
