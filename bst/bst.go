package bst

// Node 节点
type Node struct {
	val                 int
	parent, left, right *Node
}

// NewNode .
func NewNode(val int) *Node {
	return &Node{val: val}
}

// NewBST .
func NewBST() *BST {
	return &BST{}
}

// BST .
type BST struct {
	root *Node
}

// Insert .
func (b *BST) Insert(n *Node) {
	if b.root == nil {
		b.root = n
		return
	}

	var (
		node   = b.root
		parent = node.parent
	)

	for node != nil {
		// 已存在的节点，不处理
		if node.val == n.val {
			return
		}

		parent = node
		if n.val > node.val {
			node = node.right
		} else {
			node = node.left
		}
	}

	if n.val > parent.val {
		parent.right = n
	} else {
		parent.left = n
	}
	n.parent = parent
}

// Delete .
func (b *BST) Delete(n *Node) {
	if n.left == nil {
		b.Transplant(n, n.right)
		return
	}
	if n.right == nil {
		b.Transplant(n, n.left)
		return
	}

	successor := b.Successor(n)
	if successor != n.right {
		// 后继节点不是 n 的直属右节点的话，这个后继节点一定没有左节点
		b.Transplant(successor, successor.right)
		successor.right = n.right
		successor.right.parent = successor
	}
	b.Transplant(n, successor)

	successor.left = n.left
	successor.left.parent = successor
}

// Transplant 使用 v 取代 n 的位置
func (b *BST) Transplant(n *Node, v *Node) {
	if n == b.root {
		b.root = v
		if v != nil {
			v.parent = nil
		}
		return
	}
	if n.parent.left == n {
		n.parent.left = v
	} else {
		n.parent.right = v
	}
	if v != nil {
		v.parent = n.parent
	}
}

// Search .
func (b *BST) Search(val int) *Node {
	node := b.root
	for node != nil {
		if node.val == val {
			break
		}
		if val > node.val {
			node = node.right
		} else {
			node = node.left
		}
	}
	return node
}

// Max 树中最大节点
func (b *BST) Max(n *Node) *Node {
	if n == nil {
		return nil
	}
	node := n
	for node.right != nil {
		node = node.right
	}
	return node
}

// Min 树中最小节点
func (b *BST) Min(n *Node) *Node {
	if n == nil {
		return nil
	}
	node := n
	for node.left != nil {
		node = node.left
	}
	return node
}

// Successor 查找后继节点
func (b *BST) Successor(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return b.Min(n.right)
	}

	var (
		node      = n
		successor = n.parent
	)

	for successor != nil && successor.right == node {
		node = successor
		successor = successor.parent
	}
	return successor
}

// Predecessor 查找前驱节点
func (b *BST) Predecessor(n *Node) *Node {
	if n.left != nil {
		return b.Max(n.left)
	}

	var (
		node        = n
		predecessor = n.parent
	)

	for predecessor != nil && predecessor.left == node {
		node = predecessor
		predecessor = predecessor.parent
	}
	return predecessor
}
