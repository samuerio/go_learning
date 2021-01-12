package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt = MyInt(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T)  {
	a := 1
	aPoint := &a
	t.Log(a, aPoint)
	t.Logf("%T %T", a, aPoint)
}

func TestString(t *testing.T)  {
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
	if s == ""{
		t.Log(true)
	}
}