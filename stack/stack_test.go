package stack

import (
  "testing"
)

// Test IsEmpty function
func TestStackEmpty(t *testing.T) {
  stack := new(Stack)
  if !stack.IsEmpty() {
    t.Error("Problem with the function IsEmpty")
  }
  stack.Push(0)
  if stack.IsEmpty() {
    t.Error("Problem with the function IsEmpty or Push")
  }
}

// Test Simple push
func TestStackPushSimple(t *testing.T) {
  stack := new(Stack)
  stack.Push(0)
  stack.Push(1)
}

// Test Pop an empty stack
func TestStackPopEmpty(t *testing.T) {
  stack := new(Stack)
  d, e := stack.Pop()
  if e != ErrEmptyStack {
    t.Errorf("Error is not null %v, indeed Stack is empty", d)
  }
}

// Test push 10 elements and pop 10 elements
func TestStackPushAndPop(t *testing.T) {
  stack := new(Stack)
  for idx := 0; 10 > idx; idx++ {
    stack.Push(idx)
  }
  for idx := 0; 10 > idx; idx++ {
    d, e := stack.Pop()
    if nil != e {
      t.Errorf("Problem with stack, last value is %v", d)
    }
  }
}
