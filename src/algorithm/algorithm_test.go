package algorithm_learing

import (
	"testing"
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
	t.Logf("*nodeList = %v" , *nodeList)
	t.Logf("result = %v" , *result)
}

func TestNodeList_RemoveComplicate(t *testing.T) {
	nodeList := NewNodeList([]int{0, 1, 2, 3, 5, 1, 0, 3, 4})
	t.Logf("*nodeList = %v" , *nodeList)
	nodeList.RemoveComplicate()
	t.Logf("*nodeList = %v" , *nodeList)
}
