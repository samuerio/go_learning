package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T)  {
	m1 := map[int]func(op int)int{}
	m1[1] = func(op int) int {
		return op
	}
	m1[2] = func(op int) int {
		return op * op
	}
	m1[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T){
	myset := map[int]bool{}
	myset[1] = true
	n :=3
	if myset[1]{
		
	}
}