package day11

func Solve1(inp string) any {
	l := ParseInputList(inp)
	return RunGameList(l, 25)
}

func Solve2(inp string) any {
	l := ParseInputArray(inp)
	return RunGameState(l, 75)
}
