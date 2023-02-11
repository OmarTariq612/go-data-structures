package binarysearchtree

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	data        T
	left, right *node[T]
}

func (n *node[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}

func newNode[T any](data T, left, right *node[T]) *node[T] {
	return &node[T]{
		data:  data,
		left:  left,
		right: right,
	}
}

type Tree[T any] struct {
	root     *node[T]
	nodesNum uint
	less     func(T, T) bool
}

func NewTree[T any](less func(T, T) bool) *Tree[T] {
	return &Tree[T]{
		root:     nil,
		nodesNum: 0,
		less:     less,
	}
}

func (b *Tree[T]) IsEmpty() bool {
	return b.nodesNum == 0
}

func (b *Tree[T]) Size() uint {
	return b.nodesNum
}

func (b *Tree[T]) Add(elem T) bool {
	if b.Contains(elem) {
		return false
	}
	b.root = b.add(b.root, elem)
	b.nodesNum++
	return true
}

func (b *Tree[T]) Remove(elem T) bool {
	if !b.Contains(elem) {
		return false
	}
	b.root = b.remove(b.root, elem)
	b.nodesNum--
	return true
}

func (b *Tree[T]) Contains(elem T) bool {
	return b.contains(b.root, elem)
}

func (b *Tree[T]) add(root *node[T], elem T) *node[T] {
	if root == nil {
		return newNode(elem, nil, nil)
	}
	if b.less(root.data, elem) {
		root.right = b.add(root.right, elem)
	} else {
		root.left = b.add(root.left, elem)
	}
	return root
}

func (b *Tree[T]) remove(root *node[T], elem T) *node[T] {
	if root == nil {
		return nil
	}

	isCurrLessElem := b.less(root.data, elem)
	isElemLessCurr := b.less(elem, root.data)

	if !isCurrLessElem && !isElemLessCurr {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			tmp := b.findMin(root.right)
			root.data = tmp.data
			root.right = b.remove(root.right, tmp.data)
		}
	} else if isCurrLessElem {
		root.right = b.remove(root.right, elem)
	} else {
		root.left = b.remove(root.left, elem)
	}

	return root
}

func (b *Tree[T]) findMin(root *node[T]) *node[T] {
	for root.left != nil {
		root = root.left
	}
	return root
}

func (b *Tree[T]) contains(root *node[T], elem T) bool {
	if root == nil {
		return false
	}

	isCurrLessElem := b.less(root.data, elem)
	isElemLessCurr := b.less(elem, root.data)

	if !isCurrLessElem && !isElemLessCurr {
		return true
	} else if isCurrLessElem {
		return b.contains(root.right, elem)
	} else {
		return b.contains(root.left, elem)
	}
}

func (b *Tree[T]) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("BST\n")
	if !b.IsEmpty() {
		output(b.root, "", true, &strBuilder)
	}
	return strBuilder.String()
}

// from GoDS :D
// func output[T any](node *node[T], prefix string, isTail bool, str *string) {
func output[T any](node *node[T], prefix string, isTail bool, strBuilder *strings.Builder) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, strBuilder)
	}
	strBuilder.WriteString(prefix)
	if isTail {
		strBuilder.WriteString("└── ")
	} else {
		strBuilder.WriteString("┌── ")
	}
	strBuilder.WriteString(node.String() + "\n")
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, strBuilder)
	}
}
