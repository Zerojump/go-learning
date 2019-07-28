package algorithm_learing

import (
	"strings"
	"fmt"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
func sumOf2Num(target int, nums []int) (int, int, bool) {
	for i, size := 0, len(nums); i < size; i++ {
		iVal := nums[i]
		for j := 0; j < i; j++ {
			jVal := nums[j]
			if iVal+jVal == target {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//示例：
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func sumOfListNum(nums1, nums2 []int) []int {
	size1, size2 := len(nums1), len(nums2)
	var longerList, shorterList []int
	var longerSize, shorterSize int
	if size1 >= size2 {
		longerList, shorterList = nums1, nums2
		longerSize, shorterSize = size1, size2
	} else {
		longerList, shorterList = nums2, nums1
		longerSize, shorterSize = size2, size1
	}
	result := make([]int, shorterSize, longerSize+1)
	carry := 0
	for i := 0; i < shorterSize; i++ {
		num1, num2 := longerList[i], shorterList[i]
		sum := num1 + num2 + carry
		//余数
		remainder := sum % 10
		//进位数
		carry = sum / 10
		result[i] = remainder
	}
	if carry > 0 {
		for i := shorterSize; i < longerSize; i++ {
			if carry > 0 {
				//余数
				sum := longerList[i] + carry
				remainder := sum % 10
				//进位数
				carry = sum / 10
				result = append(result, remainder)
			} else {
				result = append(result, longerList[i])
			}
		}
		if carry > 0 {
			result = append(result, carry)
		}
	} else {
		for i := shorterSize; i < longerSize; i++ {
			result = append(result, longerList[i])
		}
	}
	return result
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func longestUncomplicatableSubStr(source string) string {
	unComplicatedSubStr := ""
	var subStr string
	for j, size := 0, len(source); j < size; j++ {
		jVal := source[j]
		if indexRune := strings.IndexRune(subStr, rune(jVal)); indexRune == -1 {
			subStr += string(jVal)
			if len(subStr) > len(unComplicatedSubStr) {
				unComplicatedSubStr = subStr
			}
		} else {
			if indexRune > (len(subStr)+1)/2 {
				if len(unComplicatedSubStr) < indexRune+1 {
					unComplicatedSubStr = subStr[:indexRune+1]
				}
			}
			if indexRune > len(subStr)-1 {
				subStr = string(jVal)
			} else {
				subStr = subStr[indexRune+1:] + string(jVal)
			}
		}
	}
	return unComplicatedSubStr
}

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//示例：
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
type node struct {
	next *node
	code int
}

func (self node) String() string {
	if self.next == nil {
		return fmt.Sprintf("[%d]->", self.code)
	} else {
		return fmt.Sprintf("[%d]->[%d]", self.code,self.next.code)
	}
}

type nodeList struct {
	head *node
}

func NewNodeList(arr []int) *nodeList {
	head := node{code: arr[0]}
	prev := &head
	for i := 1; i < len(arr); i++ {
		prev.next = &node{code: arr[i]}
		prev = prev.next
	}
	return &nodeList{head: &head}
}

func (self nodeList) String() string {
	result := self.groupby(func(pos int, n *node) interface{} {
		return fmt.Sprintf("[%d:%d]", pos, n.code)
	}, func(i1, i2 interface{}) interface{} {
		if i1 == nil {
			i1 = ""
		}
		return i1.(string) + i2.(string) + "->"
	})
	return result.(string)
}

func (self nodeList) groupby(f func(pos int, n *node) interface{}, merge func(i1, i2 interface{}) interface{}) (result interface{}) {
	pos := 0
	for cur := self.head; cur != nil; {
		result = merge(result, f(pos, cur))
		pos++
		cur = cur.next
	}
	return
}

func (self *nodeList)ReverseList() {
	//先反转链表，再删除
	var prev *node
	for crt :=self.head; crt!=nil;  {
		next:=crt.next
		crt.next = prev
		prev = crt
		crt = next
	}
	self.head=prev
}

func (self *nodeList) RevRemove(idx uint) (head,removed *node) {

	size := self.groupby(func(pos int, n *node) interface{} {
		return 1
	}, func(i1, i2 interface{}) interface{} {
		if i1 == nil {
			i1 = 0
		}
		return i1.(int) + i2.(int)
	}).(int)
	if size >= int(idx) {
		targetPrev := &node{next: self.head}
		for i := 0; i < size-int(idx); i++ {
			targetPrev = targetPrev.next
		}
		removed=targetPrev.next
		node := removed.next
		targetPrev.next.next = nil
		targetPrev.next = node

		if removed == self.head {
			self.head = node
		}
	}
	head = self.head
	return
}
