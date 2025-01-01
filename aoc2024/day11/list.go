package day11

import (
	"container/list"

	"github.com/shadiestgoat/aoc/utils/sparse"
)

func ParseInputList(inp string) *list.List {
	l := list.New()

	for _, v := range sparse.SplitAndParseInt(inp, " ") {
		l.PushFront(v)
	}

	return l
}

func RunGameList(l *list.List, moves int) int {
	for i := 0; i < moves; i++ {
		cur := l.Front()
		for cur != nil {
			v := cur.Value.(int)

			o := numberLogic(v)
			cur.Value = o[0]
			if len(o) > 1 {
				cur = l.InsertAfter(o[1], cur)
			}

			cur = cur.Next()
		}
	}

	return l.Len()
}
