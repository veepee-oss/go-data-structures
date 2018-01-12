package generaltree

// Tree Struct
// This struct represent the struct of general tree
type Tree struct {
  Key   interface{}
  Child []*Tree
}

// New funcion
// This function creates a pointer of Tree struct
// and set the value parameter in the Tree struct
func New(value interface{}) *Tree {
  t := new(Tree)
  t.Key = value
  t.Child = make([]*Tree, 0)
  return t
}
