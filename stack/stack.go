package stack

import "errors"

// Hide node struct
// Two members : a data interface{}
// and a pointer to node
type node struct {
  Data interface{}
  Next *node
}

// Member fonction init
// Init node with d parameter for its data
// next to nil
func (n *node) init(d interface{}) {
  n.Data = d
  n.Next = nil
}

// Stack structure
// One member : a pointer of node that is
// the head of the Stack
type Stack struct {
  Head *node
}

// Member function IsEmpty
// Check if stack is empty
func (st *Stack) IsEmpty() bool {
  return nil == st.Head
}

// Member function Push
// Add d (data) parameter to the stack
func (st *Stack) Push(d interface{}) {
  elem := new(node)
  elem.init(d)
  elem.Next = st.Head
  st.Head = elem
}

// Member function Pop
// Delete the head of the stack
func (st *Stack) Pop() (interface{}, error) {
  if st.IsEmpty() {
    return nil, errors.New("Stack is empty")
  }
  ret := st.Head.Data
  st.Head = st.Head.Next
  return ret, nil
}
