package day24

import (
	"slices"
	"strings"
)

type Tree struct {
	Left, Right *Tree
	Res string
	op string
}

// abc -\
//      OR -> efg
// cba -/

func longestLine(lines []string) (i, size int) {
	for j, l := range lines {
		if len(l) > size {
			i, size = j, len(l)
		}
	}

	return
}

func (t Tree) addSelf(last []string, down bool) []string {
	mid, baseOff := longestLine(last)

	slashChar := "/"
	if down {
		slashChar = "\\"
	}

	cur := []string{" --" + slashChar}

	dir := -1
	lineAmt := mid
	if down {
		dir = 1
		lineAmt = len(last) - mid - 1
	}

	lastOff := 4
	for i := 0; i < lineAmt; i++ {
		lastOff = 4 + i
		cur = append(cur, strings.Repeat(" ", lastOff) + slashChar)
	}

	for i, l := range cur {
		oldI := mid + i * dir
		off := baseOff - len(last[oldI])
		last[oldI] += strings.Repeat(" ", off) + l
	}

	return last
}

// Returns the lines of a string
func (t Tree) recString() []string {
	if t.Left == nil {
		return []string{t.Res}
	}

	l, r := t.Left.recString(), t.Right.recString()
	l = t.addSelf(l, true)
	r = t.addSelf(r, false)

	_, ll := longestLine(l)
	_, lr := longestLine(r)
	if ll != lr {
		diff := ll - lr
		if diff < 0 {
			diff = -diff
		}

		off := strings.Repeat(" ", diff)
		sl := l
		if ll > lr {
			sl = r
		}

		for i, line := range sl {
			sl[i] = off + line
		}
	}

	opP := t.op
	if t.op == "OR" {
		opP += " "
	}

	opLine := strings.Repeat(" ", len(r[0]) - 1) + opP + " -> " + t.Res

	return append(append(slices.Clone(l), opLine), r...)
}

func mkDepTree(s *State, elm string, car string) *Tree {
	g, ok := s.OutToGate[elm]
	if !ok {
		return &Tree{
			Res: elm,
		}
	}
	if elm == car {
		return &Tree{
			Res:   "{c}",
		}
	}

	return &Tree{
		Left:  mkDepTree(s, g.Nums[0], car),
		Right: mkDepTree(s, g.Nums[1], car),
		Res:   elm,
		op:    string(g.Op),
	}
}

func fullDrawDepMap(s *State, root, car string) string {
	tree := mkDepTree(s, root, car)
	treeLen := 0

	cur := tree
	for cur != nil {
		treeLen++
		cur = cur.Left
	}

	return strings.Join(tree.recString(), "\n")
}