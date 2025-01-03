package tset_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/utils/sillyset/tset"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func testClosestX(t *testing.T, small bool) {
	t.Run("InSame", func(t *testing.T) {
		t.Run("Seg1", func(t *testing.T) {
			s := tset.CreateSet(0, 3, 5, 10)
	
			if small {
				tutils.Assert(t, 3, s.ClosestSmall(4))
			} else {
				tutils.Assert(t, 5, s.ClosestBig(4))
			}
		})

		t.Run("Seg2", func(t *testing.T) {
			s := tset.CreateSet(
				1, 3,
				64 + 1, 64 + 6,
				64 * 2 + 9, 64 * 2 + 20,
			)

			if small {
				tutils.Assert(t, 64 + 1, s.ClosestSmall(64 + 4))
			} else {
				tutils.Assert(t, 64 + 6, s.ClosestBig(64 + 4))
			}
		})
	})

	t.Run("InDiff", func(t *testing.T) {
		t.Run("Seg1", func(t *testing.T) {
			s := tset.CreateSet(
				64 * 2 + 6, 64 * 2 + 9,
			)
	
			if small {
				tutils.Assert(t, -1, s.ClosestSmall(7))
			} else {
				tutils.Assert(t, 64 * 2 + 6, s.ClosestBig(7))
			}
		})

		t.Run("Seg2", func(t *testing.T) {
			s := tset.CreateSet(
				1, 3,
				64 * 2 + 6, 64 * 2 + 9,
			)
	
			if small {
				tutils.Assert(t, 3, s.ClosestSmall(64 + 7))
			} else {
				tutils.Assert(t, 64 * 2 + 6, s.ClosestBig(64 + 7))
			}
		})

		t.Run("Seg3", func(t *testing.T) {
			s := tset.CreateSet(
				1, 3,
			)
	
			if small {
				tutils.Assert(t, 3, s.ClosestSmall(64 * 2 + 7))
			} else {
				tutils.Assert(t, -1, s.ClosestBig(64 * 2 + 7))
			}
		})
	})
}

func TestClosestSmall(t *testing.T) {
	testClosestX(t, true)
}

func TestClosestBig(t *testing.T) {
	testClosestX(t, false)
}

func TestUpdateAt(t *testing.T) {
	s := tset.Set{}

	tutils.Assert(t, s, s.UpdateAt(5, 1).UpdateAt(5, 0))
}
