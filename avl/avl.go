package main

import (
	"math"
)

// AVL represent an AVL Tree
type AVL struct {
	root *Node
}

// NewAVL create an AVL
func NewAVL() *AVL {
	return &AVL{
		root: NewNode(0),
	}
}

// Add an element into AVL Tree
func (a *AVL) Add(val int) {
	a.root = a.root.Add(val)
}

// Node avl node
type Node struct {
	val         int
	left, right *Node
}

// NewNode create a Node
func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

// Add 添加一个节点
func (n *Node) Add(val int) *Node {
	if n.val == val {
		return n
	}
	if n.val > val {
		if n.left == nil {
			n.left = NewNode(val)
		} else {
			n.left = n.left.Add(val)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(val)
		} else {
			n.right = n.right.Add(val)
		}
	}
	return n.autoRotate()
}

// autoRotate automate rotate Nodes
func (n *Node) autoRotate() *Node {
	ld := n.leftMaxHeight()
	rd := n.rightMaxHeight()
	// 差值在 1 以内，不需要自旋
	if math.Abs(float64(ld-rd)) <= 1 {
		return n
	}
	// 左侧深度大于右侧深度，需要右旋
	if ld > rd {
		// 右旋之前先检查左侧节点的右子节点是否比左子节点更深
		// 如果 true 的话，左侧节点需要先左旋
		if n.left.rightMaxHeight() > n.left.leftMaxHeight() {
			n.left = n.left.leftRotate()
		}
		return n.rightRotate()
	} else {
		if n.right.leftMaxHeight() > n.right.rightMaxHeight() {
			n.right = n.right.rightRotate()
		}
		return n.leftRotate()
	}
}

// leftMaxHeight max height of left child node
func (n *Node) leftMaxHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.maxHeight() + 1
}

// rightMaxHeight max height of right child node
func (n *Node) rightMaxHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.maxHeight() + 1
}

// maxHeight calculate max height of node
func (n *Node) maxHeight() int {
	if n.left == nil && n.right == nil {
		return 0
	}
	if n.left == nil {
		return n.right.maxHeight() + 1
	}
	if n.right == nil {
		return n.left.maxHeight() + 1
	}
	return max(n.left.maxHeight()+1, n.right.maxHeight()+1)
}

// rightRotate 右旋
/*
有这样的一棵树

 			4
 	      /
 	    2
 	  /  \
 	1	  3

可以看到 4 作为 root 节点，左树的最大深度为 2，右侧的最大深度为 0，
差值为 2，说明不是平衡二叉树(差值大于 |1| )

因为是左侧比右侧更深，所以需要右旋:
	root 成为左子树的右子树，
	左子树的右子树成为 root 的左子树（如果有）

经过了右旋以后的树会变成这个样子

		2
      /  \
    1     4
         /
       3

成为了一棵平衡二叉树


有这样的一棵树

		5
      /  \
    3     8
        /  \
      6     9

当我们向树中添加 7 时会变成这个样子
		5
      /  \
    3     8
        /  \
      6     9
	   \
		7

直接左旋
			8
	 	  /	 \
		5	  9
      /  \
    3	  6
		   \
			7

还是不平衡，按照规则需要右旋
		5
	  /  \
    3	  8
		/  \
  	  6		9
	   \
		7

好了，由变成了原来的样子

----

先以 8 作为 root 进行右旋
		5
	  /  \
    3	  6
		   \
			8
		  /  \
		7	  9

再以 5 为 root 进行左旋

			6
		  /  \
	    5     8
	  /		/  \
	3	  7     9

就变成了一棵平衡二叉树

还是这棵树
		5
      /  \
    3     8
        /  \
      6     9

如果添加 10 会变成这个样子

		5
      /  \
    3     8
        /  \
      6     9
			 \
			 10

左旋
			8
	 	  /	 \
		5	  9
      /  \     \
    3	  6    10


当 root 节点需要左旋操作时，如果右侧子节点的左侧最大深度大于右侧最大深度，
那么需要先对 root 的右子节点进行右旋，然后再将 root 节点左旋
*/
func (n *Node) rightRotate() *Node {
	root := n.left
	n.left = root.right
	root.right = n
	return root
}

/*
左旋
		4
		 \
		  6
        /  \
	  7  	8

同理，这棵树的右侧深度为 2，左侧深度为 0，是一棵不平衡的二叉树
需要左旋

		6
      /  \
    4	  8
     \
	  7
*/
func (n *Node) leftRotate() *Node {
	root := n.right

	n.right = root.left
	root.left = n
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
