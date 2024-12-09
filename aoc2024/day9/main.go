package day9

import (
	"fmt"
	"slices"

	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) []int {
	disk := []int{}

	fID := 0
	for i, r := range inp {
		size := utils.ParseInt(string(r))
		v := -1

		if i % 2 == 0 {
			v = fID
			fID++
		}

		disk = append(disk, slices.Repeat([]int{v}, size)...)
	}

	return disk
}

func debugPrint(d []int) {
	str := ""

	for _, v := range d {
		if v == -1 {
			str += "."
		} else {
			str += fmt.Sprint(v)
		}
	}

	fmt.Println(str)
}

func diskChecksum(d []int) int {
	tot := 0
	for i, v := range d {
		if v == -1 {
			continue
		}

		tot += i * v
	}

	return tot
}

func Solve1(inp string) any {
	disk := parseInput(inp)
	dI := 0

	for i := len(disk) - 1; i >= 0; i-- {
		for dI < len(disk) && disk[dI] != -1 {
			dI++
		}
		if dI >= len(disk) || dI >= i {
			break
		}

		disk[dI], disk[i] = disk[i], -1
	}

	return diskChecksum(disk)
}

func count(s []int, i, v int, dir int) int {
	c := 0

	for {
		if i < 0 || i >= len(s) || s[i] != v {
			return c
		}
		c++
		i += dir
	}
}

func Solve2(inp string) any {
	disk := parseInput(inp)
	var blockSize = 1

	for i := len(disk) - 1; i >= 0; i -= blockSize {
		if disk[i] == -1 {
			blockSize = 1
			continue
		}

		blockSize = count(disk, i, disk[i], -1)

		for j := 0; j < len(disk) && j < i; j++ {
			if disk[j] != -1 {
				continue
			}
			freeSize := count(disk, j, -1, 1)

			if freeSize < blockSize {
				continue
			}

			for k := 0; k < blockSize; k++ {
				disk[j + k] = disk[i - k]
				disk[i - k] = -1
			}

			break
		}
	}

	return diskChecksum(disk)
}