package binarysearchtree

// provides inorder traversal for the BST
type Iterator[T any] struct {
	stack []*node[T]
}

func (b *Tree[T]) Iterator() Iterator[T] {
	iterator := Iterator[T]{stack: []*node[T]{}}
	for curr := b.root; curr != nil; curr = curr.left {
		iterator.stack = append(iterator.stack, curr)
	}
	return iterator
}

func (iterator *Iterator[T]) Next() bool {
	return len(iterator.stack) != 0
}

func (iterator *Iterator[T]) Value() T {
	reqNode := iterator.stack[len(iterator.stack)-1]
	iterator.stack[len(iterator.stack)-1] = nil
	iterator.stack = iterator.stack[:len(iterator.stack)-1]
	for curr := reqNode.right; curr != nil; curr = curr.left {
		iterator.stack = append(iterator.stack, curr)
	}
	return reqNode.data
}
