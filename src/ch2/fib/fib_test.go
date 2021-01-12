package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {

	a := 1;
	b := 1;

	for i:=0; i<5; i++{
		fmt.Println(" ",b)
		temp := a
		a=b
		b=temp+a
	}
	
}

func TestExchange(t *testing.T)  {
	a := 1
	b := 2
	a,b = b,a
	t.Log(a, b)
}
