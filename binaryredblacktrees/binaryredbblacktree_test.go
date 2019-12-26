package binaryredblacktrees

import (
	"container/list"
	"math/rand"
	"strconv"
	"testing"
)

func TestSimple(t *testing.T) {
	var tree2 BinaryRedBlackTree
	lst := list.New()
	for i := 0; i<10000; i++ {
		key := rand.Int()
		lst.PushBack(key)
	}

	for element := lst.Front(); element != nil; element = element.Next() {
		tree2.Put(element.Value.(int), rand.Int())
		tree2.BlackHeight()
	}

	for element := lst.Front(); element != nil; element = element.Next() {
		tree2.Delete(element.Value.(int))
		height, check := tree2.BlackHeight()
		if !check {
			t.Error("Black height not equal among subtrees. Black-height: " + strconv.Itoa(height))
		}
	}
}
