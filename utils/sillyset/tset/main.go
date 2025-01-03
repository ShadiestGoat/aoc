
package tset

import "math/bits"

type Set [3]uint64
const SET_LEN uint64 = 3 * 64

func (s Set) And(s2 Set) Set {
	res := Set{}

	for i, p := range s {
		p2 := s2[i]
		res[i] = p & p2
	}

	return res
}

func CreateSet(poses ...uint64) Set {
	s := Set{}

	for _, p := range poses {
		s[p/64] |= 1 << (p % 64)
	}

	return s
}

// Update val at pos, v should be 1/0
func (s Set) UpdateAt(pos uint64, v uint64) Set {
	m := pos % 64
	s[pos/64] = (s[pos/64] & ^(1 << m)) | (v << m)

	return s
}

func (s Set) At(id uint64) bool {
	return ((s[id/64] >> (id % 64)) & 1) == 1
}

func (s Set) Len() int {
	t := 0
	for _, p := range s {
		t += bits.OnesCount64(p)
	}

	return t
}


// Get the position of the closest 1 bit to off in the right direction
func (s Set) ClosestBig(off int) int {
	m := off % 64
	first := s[off/64] >> m << m
	if first != 0 {
		return bits.TrailingZeros64(first) + off - m
	}

	for i := off/64 + 1; i < len(s); i++ {
		if s[i] != 0 {
			return i * 64 + bits.TrailingZeros64(s[i])
		}
	}

	return -1
}

// Get the position of the closest 1 bit to off in the direction of small bit pos
func (s Set) ClosestSmall(off int) int {
	m := off % 64

	first := s[off/64] << (63 - m)
	if first != 0 {
		return off - bits.LeadingZeros64(first)
	}

	for i := int(off/64) - 1; i >= 0; i-- {
		if s[i] != 0 {
			return (i + 1) * 64 - 1 - bits.LeadingZeros64(s[i])
		}
	}

	return -1
}
