package avl

import (
	"fmt"
)

type AvlNode struct {
	left   *AvlNode
	right  *AvlNode
	value  int
	height int
}

func (a *AvlNode) Insert(value int) *AvlNode {
	node := NewAvlNode(value)
	return a.insert(node)
	// return a.insert(node)
}

func (a *AvlNode) getBalance() int {
	if a.left != nil && a.right != nil {
		return a.left.height - a.right.height
	} else if a.left != nil {
		return a.left.height
	} else if a.right != nil {
		return 0 - a.right.height
	}
	return 0
}

func (a *AvlNode) insert(node *AvlNode) *AvlNode {
	if node.value < a.value {
		if a.left != nil {
			b := a.left.insert(node)
			a.left = b
		} else {
			a.left = node
		}
	} else {
		if a.right != nil {
			b := a.right.insert(node)
			a.right = b
		} else {
			a.right = node
		}
	}
	a.setHeight()
	return a.reBalance()
}

func (a *AvlNode) setHeight() {
	if a.left != nil && a.right != nil {
		a.height = max(a.right.height, a.left.height) + 1
	} else if a.left != nil {
		a.height = a.left.height + 1
	} else if a.right != nil {
		a.height = a.right.height + 1
	} else {
		a.height = 1
	}
}

func (a *AvlNode) reBalance() *AvlNode {
	balance := a.getBalance()
	if balance > 1 {
		balance1 := a.left.getBalance()
		if balance1 > 0 {
			return a.rightRotate()
		} else {
			b := a.left.leftRotate()
			a.left = b
			return a.rightRotate()
		}
	} else if balance < -1 {
		balance1 := a.right.getBalance()
		if balance1 > 0 {
			b := a.right.rightRotate()
			a.right = b
			return a.leftRotate()
		} else {
			return a.leftRotate()
		}
	}
	return a
}

func (a *AvlNode) rightRotate() *AvlNode {
	fmt.Println("RightRotate: ", a.value)
	leftChild := a.left
	rightGrandChild := a.left.right
	a.left = rightGrandChild
	leftChild.right = a
	a.setHeight()
	leftChild.setHeight()
	return leftChild
}

func (a *AvlNode) leftRotate() *AvlNode {
	fmt.Println("LeftRotate: ", a.value)
	rightChild := a.right
	leftGrandChild := a.right.left
	a.right = leftGrandChild
	rightChild.left = a
	a.setHeight()
	rightChild.setHeight()
	return rightChild
}

func (a *AvlNode) Print() {
	fmt.Println("Value : ", a.value, "Height: ", a.height)
	if a.left != nil {
		fmt.Println("Left")
		a.left.Print()
	}
	if a.right != nil {
		fmt.Println("Right")
		a.right.Print()
	}
}

func NewAvlNode(value int) *AvlNode {
	a := &AvlNode{value: value, height: 1}
	return a
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
