package binarytree

import (
	"errors"
	"sync"
)

type Obj interface {
	Compare(node Obj) int
}

type Tree struct {
	lock        sync.RWMutex
	Item        Obj
	Rigth, Left *Tree
	height      int16
}

func NewTree() *Tree {
	return &Tree{Item: nil, Rigth: nil, Left: nil, height: 0}
}

func inOrder(t *Tree, chTree chan Obj) {
	if t != nil {
		inOrder(t.Left, chTree)
		chTree <- t.Item
		inOrder(t.Rigth, chTree)
	}
}

func InOrder(t *Tree, chTree chan Obj) {
	inOrder(t, chTree)
	close(chTree)
}

func postOrder(t *Tree, chTree chan Obj) {
	if t != nil {
		postOrder(t.Left, chTree)
		postOrder(t.Rigth, chTree)
		chTree <- t.Item
	}
}

func PostOrder(t *Tree, chTree chan Obj) {
	postOrder(t, chTree)
	close(chTree)
}

func preOrder(t *Tree, chTree chan Obj) {
	if t != nil {
		chTree <- t.Item
		preOrder(t.Left, chTree)
		preOrder(t.Rigth, chTree)
	}
}

func PreOrder(t *Tree, chTree chan Obj) {
	preOrder(t, chTree)
	close(chTree)
}

func (t *Tree) Insert(item Obj) error {
	if t.Item == nil {
		t.lock.Lock()
		defer t.lock.Unlock()
		t.Item = item
		return nil
	}
	if t.Item.Compare(item) == 1 { //Left
		if t.Left == nil {
			t.lock.Lock()
			defer t.lock.Unlock()
			t.Left = &Tree{Item: item,
				Rigth: nil,
				Left:  nil,
			}
			return nil
		} else {
			return t.Left.Insert(item)
		}
	} else if t.Item.Compare(item) == -1 { //Rigth
		if t.Rigth == nil {
			t.lock.Lock()
			defer t.lock.Unlock()
			t.Rigth = &Tree{Item: item,
				Rigth: nil,
				Left:  nil,
			}
			return nil
		} else {
			return t.Rigth.Insert(item)
		}
	}
	return errors.New("Item already exist")
}

func (t *Tree) Delete(item Obj) error {
	return nil
}

func Find(t *Tree, item Obj, ch_result chan Tree) error {
	if t != nil {
		if t.Item.Compare(item) == 0 {
			ch_result <- *t
			close(ch_result)
			return nil
		} else if t.Item.Compare(item) == -1 { //Rigth
			return Find(t.Rigth, item, ch_result)
		} else { //Left
			return Find(t.Left, item, ch_result)
		}
	}
	close(ch_result)
	return errors.New("Item not Found")
}
