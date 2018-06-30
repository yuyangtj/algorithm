package main

import (
	"algorithm/queue"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// TODO:
// func NewTree
// func NewEmptyTree

//Node represents a node in the BST
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	key    int
	depth  int
}

//Tree represents the whole BST
type Tree struct {
	root   *Node
	height int
}

//NewNode returns a node
func NewNode(k int) *Node {
	return &Node{parent: (*Node)(nil), left: (*Node)(nil), right: (*Node)(nil), key: k}
}

//NewNodes returns a list of nodes
func NewNodes(k ...int) []*Node {
	if len(k) == 1 {
		return []*Node{NewNode(k[0])}
	}
	nodes := []*Node{}
	for _, v := range k {
		nodes = append(nodes, NewNode(v))
	}
	return nodes
}

//String prints the whole tree
func (t *Tree) String() string {
	s := &[]string{}
	inorderTreeWalk(t.root, s)
	return strings.Join(*s, " ")
}

//InorderTreeWalk walks the tree in order

func PrintTree(t *Tree) {
	q := queue.NewQueue(128)
	q.EnQueue(t.root)
	depth := 0
	for q.Length > 0 {
		node, err := q.DeQueue()
		if err != nil {
			fmt.Printf("failed due to %v", err)
		}
		// fmt.Printf("depth: %v\n", node.(*Node).depth)
		if node.(*Node).depth > depth {
			fmt.Println()
			fmt.Println()
			depth = node.(*Node).depth
		}
		fmt.Printf("%v ", node.(*Node).key)
		if node.(*Node).left != (*Node)(nil) {
			q.EnQueue(node.(*Node).left)
		}
		if node.(*Node).right != (*Node)(nil) {
			q.EnQueue(node.(*Node).right)
		}
	}

}
func inorderTreeWalk(x *Node, s *[]string) {

	if x != nil {
		inorderTreeWalk(x.left, s)
		*s = append(*s, strconv.Itoa(x.key))
		inorderTreeWalk(x.right, s)
	}

}

func InorderTreeWalk(x *Node) {

	if x != (*Node)(nil) {

		InorderTreeWalk(x.left)
		fmt.Printf("..%+v", x)
		InorderTreeWalk(x.right)

	}

}

//Search searches node in the tree
func Search(x *Node, k int) *Node {
	if x == nil || k == x.key {
		return x
	}
	if k < x.key {
		return Search(x.left, k)
	} else {
		return Search(x.right, k)
	}

}

//IterSearch searches node in the tree iteratively
func IterSearch(x *Node, k int) *Node {
	for x != nil && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}

	}
	return x
}

//MinNode find the node with min value
func MinNode(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

//MaxNode find the node with min value
func MaxNode(x *Node) *Node {
	for x.right != nil {
		x = x.right
	}
	return x
}

//Successor returns the successor of node x
func Successor(x *Node) *Node {
	if x.right != nil {
		return MinNode(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

//InsertNode inserts an element in the tree
func InsertNode(t *Tree, n ...*Node) {
	for _, v := range n {
		insert(t, v)
	}
}

//InsertNodes inserts a list of elements in the tree
func InsertNodes(t *Tree, ns []*Node) {
	for _, v := range ns {
		insert(t, v)
	}
}

func InsertNodesRandom(t *Tree, ns []*Node) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(len(ns))
	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[l[i]].key)
		insert(t, ns[l[i]])
	}
}

func insert(t *Tree, z *Node) {
	y := (*Node)(nil)
	x := t.root
	h := 1
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
		h++

	}
	z.depth = h - 1
	if h > t.height {
		t.height = h
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func main() {

	// r := NewNode(0)
	// r1 := NewNode(1)
	// r2 := NewNode(2)
	// r3 := NewNode(3)
	// r4 := NewNode(4)
	// t := &Tree{root: r}
	// r.left = r1
	// r.right = r2
	// r1.left = r3
	// r1.right = r4
	r := (*Node)(nil)
	t := &Tree{root: r}
	ns1 := NewNodes(1, 2, 4, 3, 6, 5)
	InsertNodesRandom(t, ns1)
	// fmt.Printf("%v\n", t)
	fmt.Printf("root of the tree: %v\n", t.root.key)
	fmt.Printf("height ot the tree: %v\n", t.height)
	// key := 1
	// fmt.Printf("depth node with key %v is %v\n", key, Search(t.root, key).depth)
	PrintTree(t)
	InorderTreeWalk(t.root)
}
