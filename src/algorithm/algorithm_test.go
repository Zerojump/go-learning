package algorithm_learing

import (
	"testing"
	"math/rand"
)

func TestSumOf2Num(t *testing.T) {
	if idx1, idx2, ok := sumOf2Num(12, []int{1, 2, 3, 5, 6, 5, 7, 8, 9}); ok {
		t.Logf("idx1 = %v", idx1)
		t.Logf("idx2 = %v", idx2)
	}
}

func TestSumOfListNum(t *testing.T) {
	ints := sumOfListNum([]int{6, 3, 5, 9}, []int{5, 8, 4, 2, 1})
	t.Logf("ints = %v", ints)
}

func TestLongestUncomplicatableSubStr(t *testing.T) {
	t.Log(longestUncomplicatableSubStr("abceabcddddd"))
}

func TestNodeList_RevRemove(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 4})
	t.Logf("nodeList = %v", nodeList)
	head, removed := nodeList.RevRemove(5)
	t.Logf("nodeList = %v", nodeList)
	t.Logf("head = %v", *head)
	t.Logf("removed = %v", removed)
}

func TestNodeList_ReverseList(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 4})
	t.Logf("nodeList = %v", nodeList)
	nodeList.ReverseList()
	t.Logf("nodeList = %v", nodeList)
}

func TestRemoveComplicate(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 5, 1, 0, 3, 4})
	result := RemoveComplicate(nodeList)
	t.Logf("*nodeList = %v", *nodeList)
	t.Logf("result = %v", *result)
}

func TestNodeList_RemoveComplicate(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 5, 1, 0, 3, 4})
	t.Logf("*nodeList = %v", *nodeList)
	nodeList.RemoveComplicate()
	t.Logf("*nodeList = %v", *nodeList)
}

func TestNodeList_IsRing(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 4})
	t.Logf("nodeList.isRing() = %v", nodeList.isRing())

	target, targetN := rand.Intn(nodeList.Size()), nodeList.head
	tailN := targetN
	for i := 0; tailN.next != nil; i++ {
		if i < target {
			targetN = targetN.next
		}
		tailN = tailN.next
	}
	tailN.next = targetN

	t.Logf("nodeList.isRing() = %v", nodeList.isRing())
}

func TestRecursionReverse(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 4})
	t.Logf("nodeList = %v", *nodeList)
	nodeList.RecursionReverse()
	t.Logf("*nodeList = %v", *nodeList)
}

func TestRecursionReverseNodeListPerK(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 5, 6})
	t.Logf("*nodeList = %v", *nodeList)
	recursionReverseNodeListPerK(5, nodeList)
	t.Logf("*nodeList = %v", *nodeList)
}

func TestMerge2SortedList(t *testing.T) {
	l1 := NewNodeList([]int{1, 3, 4, 7, 8,8,10,11})
	l2 := NewNodeList([]int{0, 2, 2, 5, 9})
	list := merge2SortedList(l2, l1)
	t.Logf("list = %v" , list)

}
