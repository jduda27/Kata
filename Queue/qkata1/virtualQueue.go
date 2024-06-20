package qkata1

import (
	"fmt"
)

type Visitor struct {
	visID string
	next  *Visitor
}

type Queue struct {
	head *Visitor
}

type VisIDError struct{ visID string }
type QueueEmptyError struct{}

func (m *VisIDError) Error() string {
	return "Visitor ID " + m.visID + " not found"
}

func (m *QueueEmptyError) Error() string {
	return "Queue is empty."
}

func (q *Queue) EnterQueue(v string) {
	visitor := &Visitor{visID: v, next: nil}

	if q.head == nil {
		q.head = visitor
	} else {
		p := q.head
		for p.next != nil {
			p = p.next
		}
		p.next = visitor
	}
}

func (q *Queue) ExitQueue(target string) (err error) {
	err = nil
	tf := false
	p := q.head
	var visitorBefore, visitorAfter *Visitor

	for p != nil {
		if p.visID == target {
			tf = true
			visitorAfter = p.next
			break
		}
		if p.next != nil && p.next.visID == target {
			visitorBefore = p
		}

		p = p.next
	}

	if !tf {
		err = &VisIDError{visID: target}
	} else {
		if visitorBefore == nil {
			q.head = q.head.next
		} else {
			visitorBefore.next = visitorAfter
		}
	}

	return err
}

func (q *Queue) RemoveNextInQueue() {
	var err error
	if q.head != nil {
		p := q.head.visID
		err = q.ExitQueue(p)
	} else {
		err = &QueueEmptyError{}
	}
	if err != nil {
		fmt.Println(err)
	}
}

func Show(q Queue) {
	p := q.head
	for p != nil {
		fmt.Printf("%v -> ", p.visID)
		p = p.next
	}
	fmt.Println()
}
