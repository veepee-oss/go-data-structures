package queue

import (
  "fmt"
  "testing"
)

// Tools function for display the queue
func display(q *Queue, val int) {
  fmt.Printf("\n")
  tmp := q.Tail
  for idx := 0; val > idx && tmp != nil; idx++ {
    fmt.Printf("elem %d = %v\n", idx, tmp)
    tmp = tmp.next
  }
}

// Simple test : 10 elements push
func TestPushQueueSimple(t *testing.T) {
  queue := new(Queue)
  for idx := 0; 10 > idx; idx++ {
    queue.Push(idx)
  }
}

// Medium test : 1000 elements push
func TestPushQueueMedium(t *testing.T) {
  queue := new(Queue)
  for idx := 0; 1000 > idx; idx++ {
    queue.Push(idx)
  }
}

// Big test : 100000 elements push
func TestPushQueueBig(t *testing.T) {
  queue := new(Queue)
  for idx := 0; 100000 > idx; idx++ {
    queue.Push(idx)
  }
}

// Test pop an empty queue
func TestPopQueueEmpty(t *testing.T) {
  queue := new(Queue)
  ret, err := queue.Pop()
  if nil != ret {
    t.Errorf("Fonction pop is not define correctly indeed %v", err)
  }
}

// Test pop one element
func TestPopQueueOneElement(t *testing.T) {
  queue := new(Queue)
  queue.Push(1)
  ret, err := queue.Pop()
  if 1 == ret {
    fmt.Printf("")
  } else {
    t.Error(err)
  }
}

// Test push 10 elements and pop the 10 elements
func TestPopQueueSimple(t *testing.T) {
  queue := new(Queue)
  for idx := 0; 10 > idx; idx++ {
    queue.Push(idx)
  }
  for idx := 0; 10 > idx; idx++ {
    v, e := queue.Pop()
    t.Logf("value = %v, err = %s\n", v, e)
  }
  if false == queue.IsEmpty() {
    t.Error("Pop is not correctly")
  }
}
