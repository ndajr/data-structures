package datastructures

import "fmt"

type Node struct {
	next *Node
	data interface{}
}

func (n Node) String() string {
	return fmt.Sprintf("%s", n.data)
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(n Node) {
	if l.Head == nil {
		l.Head = &n
	} else {
		l.Tail.next = &n
	}
	l.Tail = &n
}

func (l *LinkedList) Delete(n Node) {
	var prev *Node
	curr := l.Head

	// if n is l.Head
	if curr != nil && curr.data == n.data {
		l.Head = curr.next
		return
	}

	for curr != nil {
		if curr.data == n.data {
			prev.next = curr.next
			if curr.data == l.Tail.data {
				l.Tail = prev
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

func (l LinkedList) Find(n Node) int {
	curr := l.Head
	count := 0
	for curr != nil {
		if curr.data == n.data {
			return count
		}
		count++
		curr = curr.next
	}
	return -1
}

func (l *LinkedList) Reverse(args ...*Node) {
	if l.Head == nil {
		return
	}

	var prev *Node
	var curr *Node
	if len(args) > 1 {
		// Receiving curr and prev as arguments when iterating over the list
		curr = args[0]
		prev = args[1]
	} else {
		// Initializing prev and curr for the first call
		prev = nil
		curr = l.Head
	}

	// It is necessary to use recursive approach, to be able to iterate through curr.next even when changing curr.next = prev to reverse the list.
	if curr.next == nil {
		// Base case: when it reaches the end of the list
		l.Tail = l.Head
		curr.next = prev
		l.Head = curr
		return
	} else {
		// Recursive case: when it is iterating over the list
		next := curr.next
		curr.next = prev
		l.Reverse(next, curr)
	}

}

func (l *LinkedList) Prepend(n Node) {
	if l.Tail == nil {
		l.Tail = &n
	}
	n.next = l.Head
	l.Head = &n
}

func (l LinkedList) Len() int {
	count := 0
	curr := l.Head
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (l LinkedList) String() string {
	s := "["
	curr := l.Head
	for curr != nil {
		s += fmt.Sprint(curr.data)
		if curr.next != nil {
			s += " "
		}
		curr = curr.next
	}
	s += "]"
	return fmt.Sprintf("%s", s)
}
