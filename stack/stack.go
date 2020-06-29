package stack

import "sync"

type Item interface {
}

type Stack struct {
	items []Item
	mutex sync.Mutex
}

// Dump 查看栈内容
func (stack *Stack) Dump() []Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	var copiedStack = make([]Item, len(stack.items))
	copy(copiedStack, stack.items)

	return copiedStack
}

// Peek 查看栈顶元素
func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}

// Reset 重置
func (stack *Stack) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}

// Push 入栈
func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

// IsEmpty 是否为空
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

// Pop 出栈
func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}
