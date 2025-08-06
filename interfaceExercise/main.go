package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch v2 := a.(type) {
	case int:
		return v2 * b
	case string:
		return strings.Repeat(v2, b)
	case fmt.Stringer:
		return strings.Repeat(v2.String(), b)
	default:
		fmt.Print("Неизвестный тип")
		return nil
	}
}


