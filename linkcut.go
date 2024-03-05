package linkcut

import "github.com/JonathanYuan23/link-cut-tree/splay"

// MakeTree returns a new node as a singleton tree
func MakeTree() *splay.Node {
	return splay.NewNode()
}

// Link adds c as a child of p. Assumes that c is the root of
// its own tree
func Link(c, p *splay.Node) {
	c.Access()
	p.Access()
	c.Attach(0, p)
}

// Cut detaches n from its parent
func Cut(n *splay.Node) {
	n.Access()
	n.Detach(0, n.PathParent)
	n.Child[0] = nil
	// n becomes the root of its own represented tree
	n.PathParent = nil
}

// FindRoot finds the root of the represented tree containing n.
// This operation may be time-consuming because the path to the root
// could be very long
func FindRoot(n *splay.Node) *splay.Node {
	n.Access()
	m := n
	for ; m.Child[0] != nil; m = m.Child[0] {
	}
	m.Access()
	return m
}

// LCA returns the lowest common ancestor of n and m in the
// represented tree.
func LCA(n, m *splay.Node) *splay.Node {
	if FindRoot(n) != FindRoot(m) {
		return nil
	}
	n.Access()
	return m.Access()
}
