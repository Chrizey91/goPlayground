package binaryredblacktrees

import (
	"fmt"
	"strconv"
)

type BinaryRedBlackNode struct {
	Parent     *BinaryRedBlackNode
	RightChild *BinaryRedBlackNode
	LeftChild  *BinaryRedBlackNode
	Key        int
	Value 	   interface{}
	isBlack    bool
}

type BinaryRedBlackTree struct {
	root *BinaryRedBlackNode
}

func New(k int, v interface{}, parent *BinaryRedBlackNode) *BinaryRedBlackNode {
	node := BinaryRedBlackNode{
		Key:        k,
		Value:      v,
		Parent:     parent,
	}
	return &node
}

func (n *BinaryRedBlackNode) isLeave() bool {
	return n != nil && n.RightChild == nil && n.LeftChild == nil
}

func (n *BinaryRedBlackNode) isRoot() bool {
	return n.Parent == nil
}

func (n *BinaryRedBlackNode) isLeftChild() bool {
	return n.Parent.LeftChild == n
}

func (n *BinaryRedBlackNode) isRightChild() bool {
	return n.Parent.RightChild == n
}

func (n *BinaryRedBlackNode) getUncle() *BinaryRedBlackNode {
	if n.Parent == nil || n.Parent.Parent == nil {
		return nil
	}
	if n.Parent.Parent.RightChild == n.Parent {
		return n.Parent.Parent.LeftChild
	} else {
		return n.Parent.Parent.RightChild
	}
}

func (t *BinaryRedBlackTree) Put(k int, v interface{}) {
	if t.root == nil {
		t.root = New(k, v, nil)
		t.root.isBlack = true
		return
	}

	t.putHelper(k, v, t.root)
}

func (t *BinaryRedBlackTree) PutWithPrint(k int, v interface{}) {
	if t.root == nil {
		t.root = New(k, v, nil)
		t.root.isBlack = true
		return
	}

	t.putHelper(k, v, t.root)

	t.PrintTree()
}

func (t *BinaryRedBlackTree) putHelper(k int, v interface{}, n *BinaryRedBlackNode) {
	// if value is smaller or equal the current nodes value, then insert into left child
	if k <= n.Key {
		if n.LeftChild == nil {
			n.LeftChild = New(k, v, n)
			t.insertFix1(n.LeftChild)
		} else {
			t.putHelper(k, v, n.LeftChild)
		}
		// otherwise into right child
	} else {
		if n.RightChild == nil {
			n.RightChild = New(k, v, n)
			t.insertFix1(n.RightChild)
		} else {
			t.putHelper(k, v, n.RightChild)
		}
	}
}

func (t *BinaryRedBlackTree) PrintTree() {
	fmt.Println(t.printTreeHelper(t.root))
}

func (t *BinaryRedBlackTree) printTreeHelper(n *BinaryRedBlackNode) string {
	if n == nil {
		return ""
	}
	black := ""
	if n.isBlack {
		black = "."
	}
	if n.isLeave() {
		return strconv.Itoa(n.Key) + black
	}
	return strconv.Itoa(n.Key) + black + "(" + t.printTreeHelper(n.LeftChild) + "," + t.printTreeHelper(n.RightChild) + ")"
}

func (t *BinaryRedBlackTree) rotateLeft(n *BinaryRedBlackNode) {
	if n == nil || n.RightChild == nil {
		return
	}
	parent := n.Parent
	right := n.RightChild
	if parent == nil {
		t.root = right
		right.Parent = nil
	} else {
		if parent.RightChild == n {
			parent.RightChild = right
		} else if parent.LeftChild == n {
			parent.LeftChild = right
		} else {
			panic("No child of parent matches this node!")
		}
		right.Parent = parent
	}
	rightLeft := right.LeftChild
	right.LeftChild = n
	n.Parent = right
	n.RightChild = rightLeft
	if rightLeft != nil {
		rightLeft.Parent = n
	}
}

func (t *BinaryRedBlackTree) rotateRight(n *BinaryRedBlackNode) {
	if n == nil || n.LeftChild == nil {
		return
	}
	parent := n.Parent
	left := n.LeftChild
	if parent == nil {
		t.root = left
		left.Parent = nil
	} else {
		if parent.RightChild == n {
			parent.RightChild = left
		} else if parent.LeftChild == n {
			parent.LeftChild = left
		} else {
			panic("No child of parent matches this node!")
		}
		left.Parent = parent
	}
	leftRight := left.RightChild
	left.RightChild = n
	n.Parent = left
	n.LeftChild = leftRight
	if leftRight != nil {
		leftRight.Parent = n
	}
}

// ---
// Fixes below are implemented after the lecture "Algorithms in Bioinformatics" by Prof. Dr. Florian Erhard
// Uni Wuerzburg
// Same for deletion operations
//---
func (t *BinaryRedBlackTree) insertFix1(n *BinaryRedBlackNode) {
	if n.Parent == nil || n.Parent.isBlack {
		return
	}
	uncle := n.getUncle()
	if uncle != nil && !uncle.isBlack {
		uncle.isBlack = true
		n.Parent.isBlack = true
		n.Parent.Parent.isBlack = false
		t.insertFix1(n.Parent.Parent)
	} else {
		t.insertFix2(n)
	}
	t.root.isBlack = true
}

func (t *BinaryRedBlackTree) insertFix2(n *BinaryRedBlackNode) {
	if n.Key > n.Parent.Parent.Key {
		t.insertFix2Right(n)
	} else {
		t.insertFix2Left(n)
	}
}

func (t *BinaryRedBlackTree) insertFix2Left(n *BinaryRedBlackNode) {
	if n.Key > n.Parent.Key {
		t.rotateLeft(n.Parent)
		t.rotateRight(n.Parent)
		t.recolor(n)
	} else {
		t.rotateRight(n.Parent.Parent)
		t.recolor(n.Parent)
	}

}

func (t *BinaryRedBlackTree) insertFix2Right(n *BinaryRedBlackNode) {
	if n.Key <= n.Parent.Key {
		t.rotateRight(n.Parent)
		t.rotateLeft(n.Parent)
		t.recolor(n)
	} else {
		t.rotateLeft(n.Parent.Parent)
		t.recolor(n.Parent)
	}
}

func (t *BinaryRedBlackTree) recolor(n *BinaryRedBlackNode) {
	n.isBlack = true
	if n.LeftChild != nil {
		n.LeftChild.isBlack = false
	}
	if n.RightChild != nil {
		n.RightChild.isBlack = false
	}
}

func (t *BinaryRedBlackTree) Delete(k int) bool {
	n := t.search(k)
	if n == nil {
		return false
	}
	nextHigherNode := t.getNextHigherNode(n)
	n.Key = nextHigherNode.Key
	n.Value = nextHigherNode.Value
	t.deleteHelper(nextHigherNode)
	return true
}

func (t *BinaryRedBlackTree) DeleteWithPrint(k int) bool {
	res := t.Delete(k)
	t.PrintTree()
	return res
}

// TODO: I need to clean that up!
func (t *BinaryRedBlackTree) deleteHelper(n *BinaryRedBlackNode) {
	if n.isLeave() {
		if n.isRoot() {
			t.root = nil
			return
		} else if !n.isBlack {
			if n.isLeftChild() {
				n.Parent.LeftChild = nil
			} else {
				n.Parent.RightChild = nil
			}
			t.root.isBlack = true
			return
		}
	} else if !n.isBlack {
		panic("Should not be red")
	} else {
		if n.LeftChild == nil {
			if n.isRoot() {
				t.root = n.RightChild
				n.RightChild.Parent = nil
				t.root.isBlack = true
				return
			} else if !n.RightChild.isBlack {
				n.RightChild.isBlack = true
				n.RightChild.Parent = n.Parent
				if n.isLeftChild() {
					n.Parent.LeftChild = n.RightChild
				} else {
					n.Parent.RightChild = n.RightChild
				}
				return
			}
		} else if n.RightChild == nil {
			if n.isRoot() {
				t.root = n.LeftChild
				n.LeftChild.Parent = nil
				t.root.isBlack = true
				return
			} else if !n.LeftChild.isBlack {
				n.LeftChild.isBlack = true
				n.LeftChild.Parent = n.Parent
				if n.isLeftChild() {
					n.Parent.LeftChild = n.LeftChild
				} else {
					n.Parent.RightChild = n.LeftChild
				}
				return
			}
			n.RightChild = n.LeftChild
			n.LeftChild = nil
		} else {
			panic("Both children of n are not nil. This cannot be at this state!")
		}
	}
	for {
		if t.checkDeleteOp1(n) {
			t.deleteOp1(n)
		} else if t.checkDeleteOp2(n) {
			if t.deleteOp2(n) {
				return
			}
		} else if t.checkDeleteOp3Left(n) {
			t.deleteOp3(n)
			t.deleteOp4(n)
			return
		} else if t.checkDeleteOp3Right(n) {
			t.deleteOp3(n)
			t.deleteOp4(n)
			return
		} else if t.checkDeleteOp4Left(n) {
			t.deleteOp4(n)
			return
		} else if t.checkDeleteOp4Right(n) {
			t.deleteOp4(n)
			return
		} else {
			panic("No delete operation found")
		}
	}
}

func (t *BinaryRedBlackTree) checkDeleteOp1(n *BinaryRedBlackNode) bool {
	return !n.getBrother().isBlack && n.Parent.isBlack &&
			n.getBrother().LeftChild != nil && n.getBrother().LeftChild.isBlack &&
			n.getBrother().RightChild != nil && n.getBrother().RightChild.isBlack
}

func (t *BinaryRedBlackTree) checkDeleteOp2(n *BinaryRedBlackNode) bool {
	return 	n.getBrother().isBlack && (n.getBrother().isLeave() ||
			n.getBrother().LeftChild != nil && n.getBrother().LeftChild.isBlack &&
			n.getBrother().RightChild != nil && n.getBrother().RightChild.isBlack)
}

func (t *BinaryRedBlackTree) checkDeleteOp3Left(n *BinaryRedBlackNode) bool {
	return n.isLeftChild() && n.getBrother().isBlack && n.getBrother().LeftChild != nil && !n.getBrother().LeftChild.isBlack &&
		(n.getBrother().RightChild == nil || n.getBrother().RightChild.isBlack)
}

func (t *BinaryRedBlackTree) checkDeleteOp3Right(n *BinaryRedBlackNode) bool {
	return n.isRightChild() && n.getBrother().isBlack && n.getBrother().RightChild != nil && !n.getBrother().RightChild.isBlack &&
		(n.getBrother().LeftChild == nil || n.getBrother().LeftChild.isBlack)
}

func (t *BinaryRedBlackTree) checkDeleteOp4Left(n *BinaryRedBlackNode) bool {
	return n.isLeftChild() && n.getBrother().isBlack && n.getBrother().RightChild != nil && !n.getBrother().RightChild.isBlack
}

func (t *BinaryRedBlackTree) checkDeleteOp4Right(n *BinaryRedBlackNode) bool {
	return n.isRightChild() && n.getBrother().isBlack && n.getBrother().LeftChild != nil && !n.getBrother().LeftChild.isBlack
}

// Returns the next higher node or the same node if no right child exists
func (t *BinaryRedBlackTree) getNextHigherNode(n *BinaryRedBlackNode) *BinaryRedBlackNode {
	if n.RightChild == nil {
		return n
	}
	nextHigherNode := n.RightChild
	nextNode := n.RightChild
	for {
		if nextNode.Key < nextHigherNode.Key {
			nextHigherNode = nextNode
		}
		if nextNode.isLeave() {
			break
		}
		if nextNode.LeftChild != nil {
			nextNode = nextNode.LeftChild
		} else {
			nextNode = nextNode.RightChild
		}
	}
	return nextHigherNode
}

func (t *BinaryRedBlackTree) search(k int) *BinaryRedBlackNode {
	return t.searchHelper(k, t.root)
}

func (t *BinaryRedBlackTree) searchHelper(k int, n *BinaryRedBlackNode) *BinaryRedBlackNode {
	if n == nil {
		return nil
	}
	if n.Key == k {
		return n
	}
	if k <= n.Key {
		return t.searchHelper(k, n.LeftChild)
	} else {
		return t.searchHelper(k, n.RightChild)
	}
}

func (t *BinaryRedBlackTree) deleteOp1(n *BinaryRedBlackNode) {
	if n.isLeftChild() {
		t.rotateLeft(n.Parent)
	} else {
		t.rotateRight(n.Parent)
	}
	n.Parent.isBlack = false
	n.Parent.Parent.isBlack = true
	if n.getUncle() == nil {
		panic("should not be nil")
	}
}

// returns whether operation is finished or not
func (t *BinaryRedBlackTree) deleteOp2(n *BinaryRedBlackNode) bool {
	bColorWasBlack := n.Parent.isBlack

	n.Parent.isBlack = true
	n.getBrother().isBlack = false
	if n.isLeftChild() {
		n.Parent.LeftChild = n.RightChild
	} else {
		n.Parent.RightChild = n.RightChild
	}
	if n.RightChild != nil {
		n.RightChild.Parent = n.Parent
	}
	if n.Parent.isRoot() || !bColorWasBlack {
		return true
	}
	bParent := n.Parent.Parent
	if n.Parent.isLeftChild() {
		bParent.LeftChild = n
	} else {
		bParent.RightChild = n
	}
	n.Parent.Parent = n
	n.RightChild = n.Parent
	n.Parent = bParent
	return false
}

func (t *BinaryRedBlackTree) deleteOp3(n *BinaryRedBlackNode) {
	n.getBrother().isBlack = false
	if n.isLeftChild() {
		t.rotateRight(n.getBrother())
	} else {
		t.rotateLeft(n.getBrother())
	}
	n.getBrother().isBlack = true
}

func (t *BinaryRedBlackTree) deleteOp4(n *BinaryRedBlackNode) {
	nBrother := n.getBrother()
	parentBlack := n.Parent.isBlack
	if n.isLeftChild() {
		t.rotateLeft(n.Parent)
		n.Parent.LeftChild = n.RightChild
		if n.RightChild != nil {
			n.RightChild.Parent = n.Parent
		}
	} else {
		t.rotateRight(n.Parent)
		n.Parent.RightChild = n.RightChild
		if n.RightChild != nil {
			n.RightChild.Parent = n.Parent
		}
	}
	n.Parent.Parent.isBlack = parentBlack
	nBrother.LeftChild.isBlack = true
	nBrother.RightChild.isBlack = true
	if nBrother.isRoot() {
		nBrother.isBlack = true
	}
}

func (n *BinaryRedBlackNode) getBrother() *BinaryRedBlackNode {
	if n.Parent == nil {
		return nil
	}
	if n.Parent.LeftChild == n {
		return n.Parent.RightChild
	} else {
		return n.Parent.LeftChild
	}
}

func (t *BinaryRedBlackTree) BlackHeight() (int, bool) {
	return t.blackHeightHelper(t.root)
}

func (t *BinaryRedBlackTree) blackHeightHelper(n *BinaryRedBlackNode) (int, bool) {
	if n == nil {
		return 0, true
	}
	leftHeight, checkLeft := t.blackHeightHelper(n.LeftChild)
	rightHeight, checkRight := t.blackHeightHelper(n.RightChild)
	blackScore := 0
	if n.isBlack {
		blackScore = 1
	}
	checkMiddle := leftHeight == rightHeight
	return leftHeight + blackScore, checkMiddle && checkLeft && checkRight
}
