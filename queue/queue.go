package queue

import (
  "errors"
)

var (
  // ErrQueueEmpty is an error global to check if the queue is empty or not
  ErrQueueEmpty = errors.New("Queue is empty")
)

// Elem is the struct that represent the node of the queue
// elem.data is the data of the node
// elem.next is the pointer to the next element of the queue
type elem struct {
  data interface{}
  next *elem
}

// Member function of type elem
// Initalize elem with the value d.
func newElem(d interface{}) *elem {
  e := new(elem)
  e.data = d
  e.next = nil
  return e
}

// Queue contains :
// The last element of the queue
// The next element of the tail is the first element of the queue.
type Queue struct {
  Tail *elem
}

// IsEmpty Member function
// Check if the queue is empty or not.
func (qu *Queue) IsEmpty() bool {
  return nil == qu.Tail
}

// Push Member function
// Add the data in parameter in the queue.
func (qu *Queue) Push(data interface{}) {
  node := newElem(data)
  if !qu.IsEmpty() {
    node.next = qu.Tail.next
    qu.Tail.next = node
  } else {
    node.next = node
  }
  qu.Tail = node
}

// Pop Member function
// Delete the oldest element in the queue and return
// its value an error is return if the queue is empty.
func (qu *Queue) Pop() (interface{}, error) {
  if qu.IsEmpty() {
    return nil, ErrQueueEmpty
  }
  ret := qu.Tail.next.data
  if qu.Tail.next != qu.Tail {
    qu.Tail.next = qu.Tail.next.next
  } else {
    qu.Tail = nil
  }
  return ret, nil
}
