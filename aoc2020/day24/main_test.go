package day24_test

import (
	"testing"

	"github.com/shadiestgoat/aoc/aoc2020/day24"
	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

const INPUT = `
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
`

func TestSolve1(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day24.Solve1, 10)	
}

func TestSolve2(t *testing.T) {
	tutils.AssertFunc(t, INPUT, day24.Solve2, 2208)	
}

func TestDoDirs(t *testing.T) {
	t.Run(`nwwswee`, func(t *testing.T) {
		dirs := day24.ParseInput(`nwwswee`)[0]
		v := day24.DoDirs(dirs)
		tutils.Assert(t, utils.XY{}, v)
	})

	t.Run(`esew`, func(t *testing.T) {
		dirs := day24.ParseInput(`esew`)[0]
		v := day24.DoDirs(dirs)
		tutils.Assert(t, utils.XY{1, 1}, v)
	})
}
