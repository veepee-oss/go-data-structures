package generaltree

import (
  "testing"
)

func TestTree(t *testing.T) {
  tree := New(1)
  tree.Child = append(tree.Child, New(2))
  tree.Child = append(tree.Child, New(3))
  tree.Child[0].Child = append(tree.Child[0].Child, New(4))
  tree.Child[0].Child = append(tree.Child[0].Child, New(5))
  tree.Child[0].Child = append(tree.Child[0].Child, New(6))
  tree.Child[1].Child = append(tree.Child[1].Child, New(7))
  tree.Child[1].Child = append(tree.Child[1].Child, New(8))
  tree.Child[1].Child = append(tree.Child[1].Child, New(9))
}
