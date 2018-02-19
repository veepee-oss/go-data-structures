package stack

import "errors"

var (
  // ErrEmptyStack is an error global to check if a stack is empty or not
  ErrEmptyStack = errors.New("Stack is empty")
)

// node hide struct.
// Two members : a data interface{} and a pointer to node.
type node struct {
  Data interface{}
  Next *node
}

// init member function.
// Init node with d parameter for its data next to nil.
func newNode(d interface{}) *node {
  n := new(node)
  n.Data = d
  n.Next = nil
  return n
}

// Stack structure.
// One member : a pointer of node that is the head of the Stack.
type Stack struct {
  Head *node
}

// IsEmpty member function
// Check if stack is empty.
func (st *Stack) IsEmpty() bool {
  return nil == st.Head
}

// Push member function
// Add d (data) parameter to the stack.
func (st *Stack) Push(d interface{}) {
  elem := newNode(d)
  elem.Next = st.Head
  st.Head = elem
}

// Pop Member function.
// Delete the head of the stack.
func (st *Stack) Pop() (interface{}, error) {
  if st.IsEmpty() {
    return nil, ErrEmptyStack
  }
  ret := st.Head.Data
  st.Head = st.Head.Next
  return ret, nil
}
