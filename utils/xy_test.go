package utils_test

import (
	"fmt"
	"testing"

	"github.com/shadiestgoat/aoc/utils"
	"github.com/shadiestgoat/aoc/utils/tutils"
)

func TestRotateUnitVector(t *testing.T) {
	for ogI, ogD := range utils.ALL_DIRS {
		for i := -8; i <= 8; i++ {
			t.Run(fmt.Sprintf("%v rot %d", ogD, i), func(t *testing.T) {
				expI := (ogI + i) % 8
				if expI < 0 {
					expI += 8
				}

				tutils.Assert(t, utils.ALL_DIRS[expI], ogD.RotateUnitVector(i))
			})
		}
	}
}