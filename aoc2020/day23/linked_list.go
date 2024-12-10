package day23

import (
	"strconv"
)

type LinkedList struct {
	nID map[int]*Node
}

type Node struct {
	ll *LinkedList

	Value int
	Next *Node
	Prev *Node
}

func (n *Node) InX(v int) *Node {
	cur := n
	for i := 0; i < v; i++ {
		cur = n.Next
	}

	return cur
}

func (c *Node) Append(v int) *Node {
	n := &Node{
		Value: v,
		ll: c.ll,
		Next: c.Next,
		Prev: c,
	}
	if c.ll.nID[v] != nil {
		n = c.ll.nID[v]

		n.Next = c.Next
		n.Prev = c
	}

	c.Next = n
	n.Next.Prev = n

	c.ll.nID[v] = n

	return n
}

func (n *Node) String() string {
	return strconv.Itoa(n.Prev.Value) + " <-- " + strconv.Itoa(n.Value) + " --> " + strconv.Itoa(n.Next.Value)
}