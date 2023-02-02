package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type List struct {
	node *ListNode
}

func deleteHead(head *ListNode) *ListNode {
	needDelete := head
	head.Next = needDelete
	head = head.Next

	return head
}

func (l *List) Insert(d interface{}) {
	list := &ListNode{Val: d, Next: nil}
	if l.node == nil {
		l.node = list
	} else {
		p := l.node
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

func Show(l *List) {
	p := l.node
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

func main() {
	myList := List{}

	for i := 0; i < 5; i++ {
		myList.Insert(rand.Intn(100))
	}

	Show(&myList)
	fmt.Println()

	fmt.Println(deleteHead(myList.node))

}
