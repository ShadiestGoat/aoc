package utils

import "fmt"

func PanicIfErr(err error, ctx string, args ...any) {
	if err == nil {
		return
	}

	panic("Error when " + fmt.Sprintf(ctx, args...) + ": " + err.Error())
}
