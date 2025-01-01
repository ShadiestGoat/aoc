package xprint

import (
	"encoding/json"
	"fmt"

	"github.com/shadiestgoat/aoc/utils"
)

type JSONAnswer struct {
	V any
}

func (a JSONAnswer) String() string {
	v, err := json.MarshalIndent(a.V, "", "\t")
	utils.PanicIfErr(err, "encoding answer to str")

	return string(v)
}

// Prints v on new lines as json
func PrintJSON(args ...any) {
	for _, a := range args {
		fmt.Println(JSONAnswer{V: a})
	}
}
