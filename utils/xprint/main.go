package xprint

import "fmt"

func Clear() {
	fmt.Print(string([]rune{27, '[', '3', 'J', 27, '[', ';', 'H'}))
}

func ClearAndPrint(v string) {
	Clear()
	fmt.Println(v)
}
