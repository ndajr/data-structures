package datastructures

import "fmt"

type Queue struct {
	list LinkedList
}

func (q *Queue) Enqueue(n Node) {
	q.list.Append(n)
}

func (q Queue) isEmpty() bool {
	return q.list.Head == nil
}

func (q *Queue) Dequeue() *Node {
	if q.isEmpty() {
		return nil
	}
	head := *q.list.Head
	q.list.Delete(head)
	return &head
}

func (q Queue) String() string {
	return fmt.Sprint(q.list)
}
