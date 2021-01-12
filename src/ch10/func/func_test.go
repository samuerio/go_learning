package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int)  {
	return rand.Intn(10), rand.Intn(20)
}

func logTimeSpend(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		rs := inner(op)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return rs;
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second*1)
	return op;
}

func TestFn(t *testing.T)  {
	a, _ := returnMultiValues()
	t.Log(a)
	slowFuncWrapper := logTimeSpend(slowFunc)
	t.Log(slowFuncWrapper(3))
}

func Sum(ops ...int) int {
	ret := 0;
	for _, op := range ops{
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("Clear Resource")
}

func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}