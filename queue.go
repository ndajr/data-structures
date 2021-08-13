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

func (q *Queue) Dequeue() bool {
	if q.isEmpty() {
		return false
	}
	q.list.Delete(*q.list.Head)
	return true
}

func (q Queue) String() string {
	return fmt.Sprint(q.list)
}
