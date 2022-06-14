package list

// List 双链表集合
// Double linked list set
type List[T any] struct {
	Front, Back *Node[T]
}

// Node 链表中的伪节点
// Pseudo node in linked list
type Node[T any] struct {
	Value      T
	Prev, Next *Node[T]
}

// New 返回空链表
// Return empty collection
func New[T any]() *List[T] {
	return &List[T]{}
}

// Push 传统添加集合数据
// Traditional add collection data
func (l *List[T]) Push(v T) {
	l.PushBackNode(&Node[T]{
		Value: v,
	})
}

// PushBack 将“v”添加到列表的末尾。
// adds 'v' to the end of the list.
func (l *List[T]) PushBack(v T) {
	l.PushBackNode(&Node[T]{
		Value: v,
	})
}

// PushFront 将“v”添加到列表的开头。
// adds 'v' to the beginning of the list.
func (l *List[T]) PushFront(v T) {
	l.PushFrontNode(&Node[T]{
		Value: v,
	})
}

// PushBackNode 将节点“n”添加到列表的后面。
// adds the node 'n' to the back of the list.
func (l *List[T]) PushBackNode(n *Node[T]) {
	n.Next = nil
	n.Prev = l.Back
	if l.Back != nil {
		l.Back.Next = n
	} else {
		l.Front = n
	}
	l.Back = n
}

// PushFrontNode 将节点“n”添加到列表的前面。
// adds the node 'n' to the front of the list.
func (l *List[T]) PushFrontNode(n *Node[T]) {
	n.Next = l.Front
	n.Prev = nil
	if l.Front != nil {
		l.Front.Prev = n
	} else {
		l.Back = n
	}
	l.Front = n
}

// Remove 从列表中删除n的节点。
// removes the node 'n' from the list.
func (l *List[T]) Remove(n *Node[T]) {
	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		l.Back = n.Prev
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		l.Front = n.Next
	}
}

// ForEach  对列表中从此节点开始的每个元素调用“fn”。
// calls 'fn' on every element from this node onward in the list.
func (n *Node[T]) ForEach(fn func(val T)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Next
	}
}

// EachReverse 从后面向前面进行遍历
// calls 'fn' on every element from this node backward in the list
func (n *Node[T]) EachReverse(fn func(val T)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Prev
	}
}
