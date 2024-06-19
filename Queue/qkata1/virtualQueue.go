package qkata1

type Visitor struct {
  visID string
  next *Visitor
}

type Queue struct {
  head *Visitor
}



func (q Queue) EnterQueue (v string) {
  visitor := &Visitor{visID:v ,next:nil }

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
