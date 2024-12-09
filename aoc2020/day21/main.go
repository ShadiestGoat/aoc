package day21

import (
	"slices"
	"strings"

	"github.com/shadiestgoat/aoc/utils"
)

func parseInput(inp string) (map[string]map[string]bool, map[string]int) {
	// Allergen -> ingredient list
	allergenReverseIndex := map[string]map[string]bool{}
	allIngredients := map[string]int{}

	for _, l := range strings.Split(inp, "\n") {
		spl := strings.Split(l[:len(l)-1], " (contains ")
		allergens := strings.Split(spl[1], ", ")
		ingredients := strings.Split(spl[0], " ")
		ingMap := utils.MapListKeys(ingredients)

		for _, v := range ingredients {
			allIngredients[v]++
		}
		for _, a := range allergens {
			if _, ok := allergenReverseIndex[a]; ok {
				tmpMap := map[string]bool{}

				for oldIng := range allergenReverseIndex[a] {
					if ingMap[oldIng] {
						tmpMap[oldIng] = true
					}
				}

				allergenReverseIndex[a] = tmpMap
			} else {
				allergenReverseIndex[a] = ingMap
			}
		}
	}

	return allergenReverseIndex, allIngredients
}

func Solve1(inp string) any {
	allergens, ingredients := parseInput(inp)
	tot := 0

	for v, count := range ingredients {
		found := false

		for _, poss := range allergens {
			if poss[v] {
				found = true
				break
			}
		}

		if !found {
			tot += count
		}
	}

	return tot
}

func Solve2(inp string) any {
	allergens, _ := parseInput(inp)
	resolvedAllergens := map[string]string{}
	resolvedIngredients := map[string]bool{}

	for len(allergens) != 0 {
		for a, ing := range allergens {
			if len(ing) == 1 {
				curIng := ""
				for v := range ing {
					curIng = v
				}
	
				resolvedIngredients[curIng] = true
				resolvedAllergens[a] = curIng
				delete(allergens, a)
			} else {
				for newIng := range ing {
					if resolvedIngredients[newIng] {
						delete(ing, newIng)
					}
				}
			}
		}
	}

	evilIng := []string{}
	sortedAllergens := utils.MapKeys(resolvedAllergens)
	slices.Sort(sortedAllergens)

	for _, a := range sortedAllergens {
		evilIng = append(evilIng, resolvedAllergens[a])
	}

	return strings.Join(evilIng, ",")
}
