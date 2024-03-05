package splay

// Node represents a node in a splay tree
type Node struct {
	Parent *Node
	Child  [2]*Node

	// PathParent is used for link-cut tree
	PathParent *Node
}

func NewNode() *Node {
	return &Node{}
}

// Side returns if n is a left (0) or right (1) child
func (n *Node) Side() (i int) {
	if n.Parent.Child[1] == n {
		i = 1
	} else {
		i = 0
	}
	return
}

// Attach adds a new child on side i and updates the child's
// parent link to n
func (n *Node) Attach(i int, c *Node) {
	if c != nil {
		c.Parent = n
	}
	n.Child[i] = c
}

// Detach separates subtree on side i into its own splay tree with
// pathParent set to pp
func (n *Node) Detach(i int, pp *Node) {
	if n.Child[i] != nil {
		n.Child[i].PathParent = pp
		n.Child[i].Parent = nil
	}
}

// Rotate rotates n around its parent
func (n *Node) Rotate() {
	parent := n.Parent
	grandParent := parent.Parent

	i := n.Side()

	// n replaces its parent as the grandparent's child
	if grandParent != nil {
		grandParent.Attach(parent.Side(), n)
	} else {
		n.Parent = nil
	}

	parent.Attach(i, n.Child[i^1])
	n.Attach(i^1, parent)

	n.PathParent = parent.PathParent
	parent.PathParent = nil
}

// Splay rotates n until it becomes the root of the splay tree
func (n *Node) Splay() {
	for ; n.Parent != nil; n.Rotate() {
		if n.Parent.Parent != nil {
			if n.Side() == n.Parent.Side() { // LL and RR cases
				n.Parent.Rotate()
			} else { // LR and RL cases
				n.Rotate()
			}
		}
	}
}

// Access restructures the link cut tree with a new splay path
// from the root. Returns the last pathParent before n becomes the
// root of this new splay path
func (n *Node) Access() *Node {
	n.Splay()
	n.Detach(1, n)
	n.Child[1] = nil

	parent := n
	for n.PathParent != nil {
		parent = n.PathParent
		parent.Splay()

		parent.Detach(1, n)
		parent.Attach(1, n)
		n.PathParent = nil

		n.Splay()
	}
	return parent
}
