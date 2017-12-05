package queue

import (
	"errors"
)

// Elem is the struct that represent the node of the queue
// elem.data is the data of the node
// elem.next is the pointer to the next element of the queue
type elem struct {
	data interface{}
	next *elem
}

// Member function of type elem
// Initalize elem with the value d
func (e *elem) init(d interface{}) {
	e.data = d
	e.next = nil
}

// Type queue contains :
// The last element of the queue
// The next element of the tail is the first element of the queue
type Queue struct {
	Tail *elem
}

// Member function IsEmpty
// Check if the queue is empty or not
func (qu *Queue) IsEmpty() bool {
	return nil == qu.Tail
}

// Member function Push
// Add the data in parameter in the queue
func (qu *Queue) Push(data interface{}) {
	node := new(elem)
	node.init(data)
	if !qu.IsEmpty() {
		node.next = qu.Tail.next
		qu.Tail.next = node
	} else {
		node.next = node
	}
	qu.Tail = node
}

// Member function Pop
// Delete the oldest element in the queue and return
// its value an error is return if the queue is empty
func (qu *Queue) Pop() (interface{}, error) {
	if qu.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	ret := qu.Tail.next.data
	if qu.Tail.next != qu.Tail {
		qu.Tail.next = qu.Tail.next.next
	} else {
		qu.Tail = nil
	}
	return ret, nil
}
