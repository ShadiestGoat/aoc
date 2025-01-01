package day23

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/shadiestgoat/aoc/utils/xarr"
)

func parseInput(inp string) (*LinkedList, *Node) {
	ll := &LinkedList{}
	var (
		a = &Node{
			ll:    ll,
			Value: int(inp[0] - '0'),
		}
		b = &Node{
			ll:    ll,
			Value: int(inp[1] - '0'),
		}
		c = &Node{
			ll:    ll,
			Value: int(inp[2] - '0'),
		}
	)

	ll.nID = map[int]*Node{
		a.Value: a,
		b.Value: b,
		c.Value: c,
	}

	a.Next = b
	b.Next = c
	c.Next = a

	c.Prev = b
	b.Prev = a
	a.Prev = c

	last := c

	for _, r := range inp[3:] {
		last = last.Append(int(r - '0'))
	}

	a.Prev = last

	return ll, a
}

func RunGame(ll *LinkedList, cur *Node, min, max int, moves int) {
	l := time.Now()

	for m := 0; m < moves; m++ {
		if m%100_000 == 0 && m != 0 {
			fmt.Println("On Move", m, time.Since(l))
			l = time.Now()
		}

		dstVals := []int{}
		lastTake := cur
		for i := 0; i < 3; i++ {
			lastTake = lastTake.Next
			dstVals = append(dstVals, lastTake.Value)
		}

		var dst *Node
		dstV := cur.Value

		for dst == nil {
			dstV--
			if dstV < min {
				dstV = max
			}

			if !slices.Contains(dstVals, dstV) {
				dst = ll.nID[dstV]
			}
		}

		cur.Next = lastTake.Next
		cur.Next.Prev = cur

		for _, v := range dstVals {
			dst = dst.Append(v)
		}

		cur = cur.Next
	}
}

func printState(cur *Node, pickedUp []int, dst int) {
	str := []string{}
	initV := cur.Value

	for {
		v := strconv.Itoa(cur.Value)
		if cur.Value == initV {
			v = "(" + v + ")"
		}
		str = append(str, v)

		cur = cur.Next
		if initV == cur.Value {
			break
		}
	}

	fmt.Println("cups:", strings.Join(str, " "))
	fmt.Println("pick up:", strings.Join(xarr.Map(pickedUp, strconv.Itoa), ", "))
	fmt.Println("destination:", dst)
	fmt.Println()
}

func Solve1(inp string) any {
	ll, a := parseInput(inp)
	RunGame(ll, a, 1, len(ll.nID), 100)

	str := ""

	cur := ll.nID[1].Next
	for {
		if cur.Value == 1 {
			return str
		}

		str += strconv.Itoa(cur.Value)
		cur = cur.Next
	}
}

func Solve2(inp string) any {
	ll, a := parseInput(inp)
	max := int(slices.Max([]rune(inp)) - '0')
	cur := a.Prev

	for i := max + 1; i <= 1_000_000; i++ {
		cur = cur.Append(i)
	}

	RunGame(ll, a, 1, 1_000_000, 10_000_000)

	b := ll.nID[1].Next
	c := b.Next

	return b.Value * c.Value
}
