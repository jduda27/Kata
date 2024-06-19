package dlkata1

import (
	"fmt"
)

type List struct {
	head *Node
}

type Node struct {
	val  interface{}
	next *Node
	prev *Node
}

func (l *List) Insert(d interface{}) {
	if l.head == nil {
		l.head = &Node{val: d, next: nil, prev: nil}
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = &Node{val: d, next: nil, prev: p}
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.val)
		p = p.next
	}
}
