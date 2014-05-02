package binarytree

import (
        "fmt"
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

func assertSignal(ch_signal chan struct{}, t *testing.T) {
        _, done := <-ch_signal
        if !done {
                t.Error("Error: it should send a signal of empty struct")
        }
}

func assertNoSignal(ch_signal chan struct{}, t *testing.T) {
        _, done := <-ch_signal
        if done {
                t.Error("Error: it should no send a signal")
        }
}

func TestInsertion(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        var ch_signal chan struct{}
        for i := 0; i < len(nums); i++ {
                ch_signal = Insert(tree, ObjInt{nums[i]})
                assertSignal(ch_signal, t)
        }
}

func TestNoInsertion(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        var ch_signal chan struct{}
        for i := 0; i < len(nums); i++ {
                ch_signal = Insert(tree, ObjInt{nums[i]})
                assertSignal(ch_signal, t)
        }
        ch_signal = Insert(tree, ObjInt{9})
        assertNoSignal(ch_signal, t)
}

func TestInOrder(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        for i := 0; i < len(nums); i++ {
                <-Insert(tree, ObjInt{nums[i]})
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
                <-Insert(tree, ObjInt{nums[i]})
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
                        "Error: Objects should be equals")
        }
}

func TestPreOrder(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        for i := 0; i < len(nums); i++ {
                <-Insert(tree, ObjInt{nums[i]})
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
                        "Error: Objects should be equals")
        }
}

func TestFind(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        for i := 0; i < len(nums); i++ {
                <-Insert(tree, ObjInt{nums[i]})
        }
        var ch_result chan Tree
        for i := 0; i < len(nums); i++ {
                ch_result = Find(tree, ObjInt{nums[i]})
                obj := <-ch_result
                assertTrue(obj.Item.Compare(ObjInt{nums[i]}) == 0, t,
                        "Error: I should find the requested object")
        }
}

func TestNotFound(t *testing.T) {
        tree := NewTree()
        nums := []int{9, 7, 2, 4, 6, 10, 1, 5, 8, 3}
        for i := 0; i < len(nums); i++ {
                <-Insert(tree, ObjInt{nums[i]})
        }
        var ch_result chan Tree
        not_in_tree := []int{90, 70, 20, 40, 60, 100, 110, 50, 80, 30}
        for i := 0; i < len(not_in_tree); i++ {
                ch_result = Find(tree, ObjInt{not_in_tree[i]})
                _, done := <-ch_result
                if done {
                        t.Error("Error: it should no send a signal")
                }
        }
}
