package binarytree

import (
	"fmt"
	"math/rand"
	"testing"
)

type ObjString struct {
	Id string
}

type ObjInt struct {
	Id int
}

func (o ObjInt) String() string {
	return fmt.Sprintf("%d", o.Id)
}

func (o ObjString) Compare(obj Obj) int {
	o2 := obj.(ObjString)
	if o.Id == o2.Id {
		return 0
	} else if o.Id < o2.Id {
		return -1
	} else {
		return 1
	}
}

func (o ObjInt) Compare(obj Obj) int {
	o2 := obj.(ObjInt)
	if o.Id == o2.Id {
		return 0
	} else if o.Id < o2.Id {
		return -1
	} else {
		return 1
	}
}

func TestCompare(t *testing.T) {
	int_arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	objInt_arr := make([]ObjInt, 0, len(int_arr))
	for _, id := range int_arr {
		objInt_arr = append(objInt_arr, ObjInt{id})
	}
	var res int
	for i := 0; i < len(objInt_arr)-1; i++ {
		res = objInt_arr[i].Compare(objInt_arr[i+1])
		if res != -1 {
			t.Error("Error value should be grater")
		}
	}

	for i := len(objInt_arr) - 1; i > 0; i-- {
		res = objInt_arr[i].Compare(objInt_arr[i-1])
		if res != 1 {
			t.Error("Error value should be smaller")
		}
	}

	str_arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	objstring_arr := make([]ObjString, 0, len(str_arr))
	for _, id := range str_arr {
		objstring_arr = append(objstring_arr, ObjString{id})
	}

	for i := 0; i < len(objstring_arr)-1; i++ {
		res = objstring_arr[i].Compare(objstring_arr[i+1])
		if res != -1 {
			t.Error("Error value should be grater")
		}
	}

	for i := len(objstring_arr) - 1; i > 0; i-- {
		res = objstring_arr[i].Compare(objstring_arr[i-1])
		if res != 1 {
			t.Error("Error value should be smaller")
		}
	}
}

func newObjInt() ObjInt {
	return ObjInt{rand.Intn(100)}
}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

func assertTrue(assertion bool, t *testing.T, msg string) {
	if !assertion {
		t.Error(msg)
	}
}

func TestTreeInsertion(t *testing.T) {
	tree := NewTree()
	err := tree.Insert(newObjInt())
	assertNoError(err, t)
	err = tree.Insert(newObjInt())
	assertNoError(err, t)
	err = tree.Insert(newObjInt())
	assertNoError(err, t)
	err = tree.Insert(newObjInt())
	assertNoError(err, t)
}

func TestInOrder(t *testing.T) {
	tree := NewTree()
	nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
	for i := 0; i < len(nums); i++ {
		tree.Insert(ObjInt{nums[i]})
	}
	ch_result := make(chan Obj, 10)
	go InOrder(tree, ch_result)
	var previous Obj
	for obj := range ch_result {
		if previous == nil {
			previous = obj
			continue
		}
		assertTrue(previous.Compare(obj) == -1, t,
			"Previous obj should be smaller than current object")
		previous = obj
	}
}

func TestPostOrder(t *testing.T) {
	tree := NewTree()
	nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
	for i := 0; i < len(nums); i++ {
		tree.Insert(ObjInt{nums[i]})
	}
	ch_result := make(chan Obj, 10)
	answer := []int{1, 3, 5, 6, 4, 2, 8, 7, 10, 9}
	result := make([]Obj, 0, 10)
	go PostOrder(tree, ch_result)
	for obj := range ch_result {
		result = append(result, obj)
	}
	for i, obj := range result {
		assertTrue(obj.Compare(ObjInt{answer[i]}) == 0, t,
			"Previous obj should be smaller than current object")
	}
}

func TestPreOrder(t *testing.T) {
	tree := NewTree()
	nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
	for i := 0; i < len(nums); i++ {
		tree.Insert(ObjInt{nums[i]})
	}
	ch_result := make(chan Obj, 10)
	answer := []int{9, 7, 2, 1, 4, 3, 6, 5, 8, 10}
	result := make([]Obj, 0, 10)
	go PreOrder(tree, ch_result)
	for obj := range ch_result {
		result = append(result, obj)
	}
	for i, obj := range result {
		assertTrue(obj.Compare(ObjInt{answer[i]}) == 0, t,
			"Previous obj should be smaller than current object")
	}
}