package llkata1

type SNode struct {
  val interface{}
  next *SNode
}

type DNode struct {
  val int
  prev *DNode
  next *DNode
}

type SList struct {
  head *SNode
}

func (l *SList) Insert(d interface{}) {
  list := &SNode{val:d,next: nil}

  if l.head == nil {
    l.head = list
  } else {
    p := l.head
    for p.next != nil {
      p = p.next
    }
    p.next = list
  }
}

