package stack

import "testing"

var stack Stack

func TestStack_IsEmpty(t *testing.T) {
	if !stack.IsEmpty() {
		t.Errorf("stack should be empty")
	}
	stack.Push("first")
	if stack.IsEmpty() {
		t.Errorf("stack should not be empty")
	}
	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("stack should be empty after pop all elements")
	}
	stack.Push("second")
	stack.Reset()
	if !stack.IsEmpty() {
		t.Errorf("stack should be empty after reset stack")
	}
}

func TestStack_Pop(t *testing.T) {
	item := stack.Pop()
	if item != nil {
		t.Errorf("empty stack pop should nil")
	}
	stack.Push("first")
	item = stack.Pop()
	if item == nil {
		t.Errorf("stack should pop element")
	}

	stack.Push("second")
	stack.Reset()
	item = stack.Pop()
	if item != nil {
		t.Errorf("stack pop should return nil after reset")
	}
}
