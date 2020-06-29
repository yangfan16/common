package list

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("no such node found ")

type Data interface {
}

type Node struct {
	data Data
	next *Node
}

// List
type List struct {
	head  *Node
	tail  *Node
	mutex sync.Mutex
}

func (node *Node) Data() Data {
	return node.data
}
func (node *Node) Next() *Node {
	return node.next
}

// InsertBefore
func (list *List) InsertBefore(value Data, mark *Node) *Node {

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if mark == nil {
		return nil
	}

	newNode := &Node{
		data: value,
		next: mark,
	}

	if list.head == mark {
		list.head = newNode
		return newNode
	}

	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			node.next = newNode
			return newNode
		}
	}

	return nil
}

// InsertAfter
func (list *List) InsertAfter(value Data, mark *Node) *Node {

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if mark == nil {
		return nil
	}

	for node := list.head; node != nil; node = node.next {
		if node == mark {
			newNode := &Node{
				data: value,
				next: node.next,
			}

			node.next = newNode

			if node == list.tail {
				list.tail = newNode
			}

			return newNode
		}
	}

	return nil
}

// InsertLast
func (list *List) InsertLast(value Data) *Node {

	list.mutex.Lock()
	defer list.mutex.Unlock()

	newNode := &Node{
		data: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	return newNode
}

// InsertFront
func (list *List) InsertFront(value Data) *Node {

	list.mutex.Lock()
	defer list.mutex.Unlock()

	newNode := &Node{
		data: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

	return newNode
}

// Delete
func (list *List) Delete(mark *Node) error {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if mark == nil {
		return nil
	}

	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			node.next = mark.next
			if list.tail == mark {
				list.tail = node
			}

			break
		}

		if node == mark {
			list.head = node.next
			if list.tail == mark {
				list.tail = nil
			}

			break
		}
	}

	return ErrNotFound
}

// Prev
func (list *List) Prev(mark *Node) *Node {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			return node
		}
	}

	return nil
}

// Range
func (list *List) Range(fn func(*Node)) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	for node := list.head; node != nil; node = node.next {
		fn(node)
	}
}

// Size
func (list *List) Size() int {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.head == nil {
		return 0
	}

	var size int

	for node := list.head; node != nil; node = node.next {
		size++
	}

	return size
}

// Head
func (list *List) Head() *Node {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.head
}

// Tail
func (list *List) Tail() *Node {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.tail
}

func (list *List) IsEmpty() bool {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	return list.head == nil
}
