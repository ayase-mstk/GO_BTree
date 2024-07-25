package rbtree

type Color_t int

const (
  Red Color_t = iota
  Black
)

type RBNode struct {
  Parent, Left, Right *RBNode
  Color               Color_t
  Data                string
}

var NilNode = &RBNode{Color: Black}
