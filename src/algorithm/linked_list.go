package algorithm_learing

import (
	"fmt"
)

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//示例：
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
type node struct {
	next *node
	code int
}

func (self *node) clone() *node {
	return &node{next: self.next, code: self.code}
}

func (self node) String() string {
	if self.next == nil {
		return fmt.Sprintf("[%d]->", self.code)
	} else {
		return fmt.Sprintf("[%d]->[%d]", self.code, self.next.code)
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
	result := self.aggregate(func(pos int, n *node) interface{} {
		return fmt.Sprintf("[%d:%d]", pos, n.code)
	}, func(i1, i2 interface{}) interface{} {
		if i1 == nil {
			i1 = ""
		}
		return i1.(string) + i2.(string) + "->"
	})
	return result.(string)
}

func (self nodeList) aggregate(foreach func(pos int, n *node) interface{}, merge func(i1, i2 interface{}) interface{}) (result interface{}) {
	pos := 0
	for cur := self.head; cur != nil; {
		result = merge(result, foreach(pos, cur))
		pos++
		cur = cur.next
	}
	return
}

//反转链表
func (self *nodeList) ReverseList() {
	//先反转链表，再删除
	var prev *node
	for crt := self.head; crt != nil; {
		next := crt.next
		crt.next = prev
		prev = crt
		crt = next
	}
	self.head = prev
}

func (self *nodeList) Size() int {
	return self.aggregate(func(pos int, n *node) interface{} {
		return 1
	}, func(i1, i2 interface{}) interface{} {
		if i1 == nil {
			i1 = 0
		}
		return i1.(int) + i2.(int)
	}).(int)
}

//删除倒数第idx个
func (self *nodeList) RevRemove(idx uint) (head, removed *node) {
	size := self.Size()
	if size >= int(idx) {
		targetPrev := &node{next: self.head}
		for i := 0; i < size-int(idx); i++ {
			targetPrev = targetPrev.next
		}
		removed = targetPrev.next
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

//去重
func (self *nodeList) RemoveComplicate() {
	for cur := self.head; cur != nil; cur = cur.next {
		prevN := cur
		for next := cur.next; next != nil; next = next.next {
			if cur.code == next.code {
				prevN.next = next.next
			} else {
				prevN = next
			}
		}
	}
}

//去重
func RemoveComplicate(source *nodeList) *nodeList {
	n := nodeList{head: &node{code: source.head.code}}
	for cur := source.head; cur != nil; cur = cur.next {
		fmt.Println("cur:", cur.code)
		nd := n.head
		var exists bool
		for ; nd.next != nil; nd = nd.next {
			if nd.code == cur.code {
				exists = true
				break
			}
		}
		if !exists && nd != nil {
			nd.next = &node{code: cur.code}
		}
	}
	return &n
}

//给定链表 Lo一＞L1一＞L2… Ln-1一＞Ln，把链表重新排序为 Lo->Ln一＞L1一＞Ln-1->L2一＞Ln-2。要求：
//1、在原来链表的基础上进行排序，即不能申请新的结点
//2、只能修改结点的 next 域，不能修改数据域

//如何检测一个较大的单链表是否有环
func (self *nodeList) isRing() bool {
	walk2Step := func(n *node) *node {
		if next := n.next; next != nil && next.next != nil {
			return next.next
		}
		return nil
	}

	fast := walk2Step(self.head)
	for slow := self.head; slow != nil && fast != nil; slow = slow.next {
		if fast == slow {
			return true
		}
		fast = walk2Step(fast)
	}
	return false
}

func (self *nodeList) RecursionReverse() {
	head := self.head
	//self.head=recursionReverseNode(head, head.next)
	self.head = recursionReverseNode2(head)
}

func recursionReverseNode(prev, cur *node) *node {
	if cur == nil {
		return prev
	} else {
		head := recursionReverseNode(cur, cur.next)
		cur.next = prev
		prev.next = nil
		fmt.Println("cur:", *cur, "head:", *head)
		return head
	}
}

func recursionReverseNodeListPerK(k int, list *nodeList) {
	from, to := 0, k
	joinN := list.head
	var pieceN, nextPieceHead *node
	pieceN, nextPieceHead = recursionReverseNodePerK(from, to, nil, joinN)
	list.head = pieceN
	for ; nextPieceHead != nil; {

		from += k
		to += k

		joinN.next = nextPieceHead
		fmt.Println("jsonN:", joinN, "nextPieceHead:", nextPieceHead)
		fmt.Println("nodeList:", list)
		joinN, nextPieceHead = recursionReverseNodePerK(from+k, to+k, joinN, nextPieceHead)
	}
}

func recursionReverseNodePerK(k, limit int, prev, cur *node) (*node, *node) {
	//todo
	if cur == nil {
		return prev, nil
	} else if k == limit {
		return prev, cur
	} else {
		head, nextK := recursionReverseNodePerK(k+1, limit, cur, cur.next)
		cur.next = prev
		if k+1 < limit && prev != nil {
			prev.next = nil
		}
		fmt.Println("cur:", cur, "  head:", head, "  nextK:", nextK, "  k:", k, "  limit:", limit)
		return head, nextK
	}
}

func recursionReverseNode2(cur *node) *node {
	if cur == nil || cur.next == nil {
		return cur
	} else {
		head := recursionReverseNode2(cur.next)
		cur.next.next = cur
		fmt.Println("cur:", *cur, "head:", *head)
		cur.next = nil
		return head
	}
}

//如何把链表以 K个结点为－组进行翻转
//func ReversePer(k int, list nodeList) nodeList {
//	size := list.Size()
//	if size < k {
//		return list
//	}
//
//	i := 0
//	in1st := true
//	firstK_L := nodeList{head: &node{code: list.head.code}}
//	firstK_N := firstK_L.head
//	var secondK_N *node
//	var secondK_L nodeList
//
//	for cur := list.head; cur.next != nil; cur = cur.next {
//		if in1st {
//			if i < k {
//				firstK_N.next = &node{code: cur.code}
//				firstK_N = firstK_N.next
//				i++
//			} else {
//				firstTail := firstK_L.head
//				firstK_L.ReverseList()
//				firstK_N = firstTail
//
//				secondK_N = &node{code: cur.code}
//				firstK_N.next = secondK_N
//				secondK_L = nodeList{head: secondK_N}
//				i, in1st = 0, false
//			}
//		} else {
//			if i < k {
//				secondK_N.next = &node{code: cur.code}
//				secondK_N = secondK_N.next
//				i++
//			} else {
//				secondTail := secondK_L.head
//				secondK_L.ReverseList()
//				secondK_N = secondTail
//
//				firstK_N = &node{code: cur.code}
//				secondK_N.next = firstK_N
//				firstK_L = nodeList{head: firstK_N}
//				i, in1st = 0, true
//			}
//		}
//	}
//
//}

//合并两个有序链表
//其实可以比较一次移动一次，不用多次比较再移动一次，这样增加了复杂度，又省不了多少时间
func merge2SortedList(l1, l2 *nodeList) *nodeList {
	if l1== l2 {
		return l1
	}
	var headN *node
	cur1 := l1.head
	cur2 := l2.head
	if cur1.code <= cur2.code {
		headN = cur1
	} else {
		headN = cur2
	}
	for ; cur1 != nil && cur2 != nil; {
		fmt.Println("cur1:", cur1)
		if cur1.code <= cur2.code {
			from1 := cur1
			for ; from1.next != nil; from1 = from1.next {
				if from1.next.code > cur2.code {
					break
				}
			}
			cur1 = from1.next
			fmt.Println("cur1:", cur1)

			from1.next = cur2
		}
		if cur1!= nil {
			from2 := cur2
			for ; from2.next != nil; from2 = from2.next {
				if from2.next.code > cur1.code {
					break
				}
			}
			cur2 = from2.next
			fmt.Println("cur2:", cur2)
			from2.next = cur1
		}
	}

	return &nodeList{head: headN}
}
