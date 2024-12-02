package day14

import (
	"strconv"
	"strings"

	"shadygoat.eu/aoc/utils"
)

func toBin(v int) []rune {
	bin := strconv.FormatUint(uint64(v), 2)
	return []rune(strings.Repeat("0", 36 - len(bin)) + bin)
}

type memHandler = func(mask string, m map[int]uint, addr int, val int)

func parseInput(inp string, onMem memHandler) int {
	mask := ""
	mem := map[int]uint{}

	for _, v := range strings.Split(inp, "\n") {
		spl := strings.Split(v, " = ")

		switch spl[0][:3] {
		case "mas":
			mask = spl[1]
		case "mem":
			addr, _ := strconv.Atoi(spl[0][4:len(spl[0]) - 1])
			val, _ := strconv.Atoi(spl[1])

			onMem(mask, mem, addr, val)
		}
	}

	t := uint(0)

	for _, v := range mem {
		t += v
	}

	return int(t)
}

func parseBin(bin string) uint {
	v, err := strconv.ParseUint(string(bin), 2, 64)
	utils.PanicIfErr(err, "parsing bin '%v' as uint", bin)

	return uint(v)
}

func Solve1(inp string) any {
	return parseInput(inp, func(mask string, m map[int]uint, addr, val int) {
		bin := toBin(val)

		for i, r := range mask {
			if r == 'X' {
				continue
			}

			bin[i] = r
		}

		m[addr] = parseBin(string(bin))
	})
}

func allAddr(base string, mask string) []string {
	m := map[string]bool{
		base: true,
	}

	for i, r := range mask {
		if r != 'X' {
			continue
		}

		for _addr := range m {
			addr := []rune(_addr)

			addr[i] = '0'
			m[string(addr)] = true

			addr[i] = '1'
			m[string(addr)] = true
		}
	}

	resp := make([]string, 0, len(m))

	for addr := range m {
		resp = append(resp, addr)
	}

	return resp
}

func Solve2(inp string) any {
	return parseInput(inp, func(mask string, m map[int]uint, addr, val int) {
		bin := toBin(addr)

		for i, r := range mask {
			if r == '0' {
				continue
			}
			if r == '1' {
				bin[i] = '1'
			}
		}

		allAddr := allAddr(string(bin), mask)

		for _, addr := range allAddr {
			m[int(parseBin(addr))] = uint(val)
		}
	})
}
