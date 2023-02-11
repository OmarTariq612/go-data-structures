package binarysearchtree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	b := NewTree(func(i1, i2 int) bool {
		return i1 < i2
	})

	b.Add(5)
	b.Add(6)
	b.Add(1)
	b.Add(2)
	b.Add(500)
	b.Remove(1)
	b.Add(10)
	b.Add(600)

	t.Logf("\n%v\n", b)

	for itr := b.Iterator(); itr.Next(); {
		t.Log(itr.Value())
	}
}
