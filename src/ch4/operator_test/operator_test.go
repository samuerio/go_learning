package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T)  {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	c := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
}

func TestConstantTry1(t *testing.T)  {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable,
		a&Writable == Writable, a&Executable == Executable)
}