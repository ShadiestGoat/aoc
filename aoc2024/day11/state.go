package day11

type State struct {
	// I SO cannot be fucked with the cache state being an array
	cache map[int]map[int]int
	maxMoves int
}

// Add to cache, 
func (s *State) addCache(n int, movesLeft int, v int) {
	if s.cache[n] == nil {
		s.cache[n] = map[int]int{}
	}

	s.cache[n][movesLeft] = v
}

// movesLeft - the number of moves to do after this. 0 indicates this is the last move.
func (s *State) runOne(n int, movesLeft int) int {
	if s.cache[n] != nil && s.cache[n][movesLeft] != 0 {
		return s.cache[n][movesLeft]
	}

	arr := numberLogic(n)
	s.addCache(n, 0, len(arr))

	if movesLeft == 0 {
		return len(arr)
	}

	t := 0
	for _, v := range arr {
		t += s.runOne(v, movesLeft - 1)
	}

	s.addCache(n, movesLeft, t)

	return t
}

func RunGameState(v []int, moves int) int {
	s := &State{
		cache:    map[int]map[int]int{},
		maxMoves: moves,
	}

	t := 0
	for _, n := range v {
		t += s.runOne(n, moves - 1)
	}

	return t
}
