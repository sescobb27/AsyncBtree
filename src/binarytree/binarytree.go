package binarytree

type Obj interface {
	Compare(node Obj) int
}

type Tree struct {
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

func insert(t *Tree, item Obj, done chan struct{}) {
	if t.Item == nil {
		t.Item = item
		var signal struct{}
		done <- signal
		return
	}
	if t.Item.Compare(item) == 1 { //Left
		if t.Left == nil {
			t.Left = &Tree{Item: item,
				Rigth: nil,
				Left:  nil,
			}
			var signal struct{}
			done <- signal
		} else {
			insert(t.Left, item, done)
			return
		}
	} else if t.Item.Compare(item) == -1 { //Rigth
		if t.Rigth == nil {
			t.Rigth = &Tree{Item: item,
				Rigth: nil,
				Left:  nil,
			}
			var signal struct{}
			done <- signal
		} else {
			insert(t.Rigth, item, done)
			return
		}
	}
	close(done)
	return //errors.New("Item already exist")
}

func Insert(t *Tree, item Obj) chan struct{} {
	done := make(chan struct{}, 1)
	go insert(t, item, done)
	return done
}

func (t *Tree) Delete(item Obj) error {
	// TO-DO
	return nil
}

func find(t *Tree, item Obj, ch_result chan Tree) {
	if t != nil {
		if t.Item.Compare(item) == 0 {
			ch_result <- *t
			return
		} else if t.Item.Compare(item) == -1 { //Rigth
			Find(t.Rigth, item, ch_result)
			return
		} else { //Left
			Find(t.Left, item, ch_result)
			return
		}
	}
	ch_result <- *NewTree()
}

func Find(t *Tree, item Obj, ch_result chan Tree) {
	find(t, item, ch_result)
}
