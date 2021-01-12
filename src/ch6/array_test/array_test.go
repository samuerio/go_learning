package array_test

import "testing"

func TestArrayTest(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr1[1]=5
	arr2 := [...]int{1,2,3,4}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTraval(t *testing.T){
	arr1 := [4]int{1,2,3,4}
	for i:= 0; i<len(arr1); i++ {
		t.Log(arr1[i])
	}

	for idx,item := range(arr1){
		t.Log(idx, item)
	}
}

func TestArraySelection(t *testing.T)  {
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[2:]
	t.Log(arr3_sec)
}
