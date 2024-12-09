package day22

import (
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

type Game struct {
	Players [2][]int
	History [][2][]int
	Layer int
}

func (g *Game) winnerDeck(w, l int) {
	wd, ld := g.Players[w], g.Players[l]

	g.Players[w], g.Players[l] = append(wd[1:], []int{wd[0], ld[0]}...), ld[1:]
}

func (g *Game) win(p int) {
	w, l := 0, 1
	if p == 2 {
		w, l = 1, 0
	}

	g.winnerDeck(w, l)
}

func (g *Game) checkIfWon() int {
	if len(g.Players[0]) == 0 {
		return 2
	} else if len(g.Players[1]) == 0 {
		return 1
	}

	return 0
}

func (g *Game) roundIsSame(r [2][]int) bool {
	for i, d := range r {
		pd := g.Players[i]

		if len(d) != len(pd) {
			return false
		}
		for j, v := range d {
			if v != pd[j] {
				return false
			}
		}
	}

	return true
}

func (g *Game) checkHistory() bool {
	for _, r := range g.History {
		if g.roundIsSame(r) {
			return true
		}
	}

	return false
}

func (g *Game) copyDecks() [2][]int {
	decks := [2][]int{
		make([]int, len(g.Players[0])),
		make([]int, len(g.Players[1])),
	}
	copy(decks[0], g.Players[0])
	copy(decks[1], g.Players[1])

	return decks
}

func (g *Game) SubGame() *Game {
	decks := g.copyDecks()
	newDecks := [2][]int{}

	for p, d := range decks {
		amt := d[0]

		newDecks[p] = d[1:amt + 1]
	}

	return &Game{
		Players: newDecks,
		Layer: g.Layer + 1,
	}
}

func (g *Game) regularCombat() {
	if g.Players[0][0] > g.Players[1][0] {
		g.win(1)
	} else {
		g.win(2)
	}
}

func (g *Game) playerCanRecurse(p int) bool {
	return g.Players[p][0] < len(g.Players[p])
}

func (g *Game) RecursiveCombat() int {
	if won := g.checkIfWon(); won != 0 {
		return won
	}
	if g.checkHistory() {
		return 1
	}
	g.History = append(g.History, g.copyDecks())

	if g.playerCanRecurse(0) && g.playerCanRecurse(1) {
		g.win(Forever(g.SubGame().RecursiveCombat))
	} else {
		g.regularCombat()
	}

	return 0
}

// Runs the Combat game once. Returns winning player's deck if theres a winner. Otherwise, returns empty array
func (g *Game) Combat() int {
	if won := g.checkIfWon(); won != 0 {
		return won
	}

	g.regularCombat()

	return 0
}

// Runs game until a winner emerges
func Forever(gameFunc func () int) int {
	res := gameFunc()
	for res == 0 {
		res = gameFunc()
	}

	return res
}

func WinningScore(g *Game, winner int) int {
	winningDeck := g.Players[winner - 1]

	sum := 0
	for i, w := range winningDeck {
		sum += (len(winningDeck) - i) * w
	}

	return sum
}

func parseInput(inp string) *Game {
	spl := strings.Split(inp, "\n\n")
	players := [2][]int{}

	for i, p := range spl {
		players[i] = utils.Map(strings.Split(p, "\n")[1:], utils.ParseInt)
	}

	return &Game{
		Players: players,
		History: [][2][]int{},
		Layer:   1,
	}
}

func Solve1(inp string) any {
	g := parseInput(inp)
	winner := Forever(g.Combat)

	return WinningScore(g, winner)
}

func Solve2(inp string) any {
	g := parseInput(inp)
	winner := Forever(g.RecursiveCombat)

	return WinningScore(g, winner)
}
