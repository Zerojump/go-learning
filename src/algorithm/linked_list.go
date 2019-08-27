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

//删除倒数第idx个
func (self *nodeList) RevRemove(idx uint) (head, removed *node) {

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

func (self *nodeList) RemoveComplicate() {
	for cur := self.head; cur != nil; {
		prevN :=cur
		for next := cur.next; next != nil;  {
			if cur.code==next.code {
				prevN.next = next.next
			} else {
				prevN = next
			}
			next = next.next
		}
		cur = cur.next
	}
}

func RemoveComplicate(source *nodeList) *nodeList {
	n := nodeList{head: &node{code: source.head.code}}
	for cur := source.head; cur != nil; {
		fmt.Println("cur:", cur.code)
		nd := n.head
		var exists bool
		for ; nd.next != nil; {
			if nd.code == cur.code {
				exists = true
				break
			} else {
				nd = nd.next
			}
		}
		if !exists && nd != nil {
			nd.next = &node{code: cur.code}
		}
		cur = cur.next
	}
	return &n
}
