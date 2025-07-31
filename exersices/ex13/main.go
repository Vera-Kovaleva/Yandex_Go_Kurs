package main

import (
	mathpakage "exersices/ex13/mathPakage"
	"fmt"
)

func main() {

	s := mathpakage.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", mathpakage.SumSlice(s))

	mathpakage.MapSlice(s, func(i mathpakage.Element) mathpakage.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", mathpakage.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		mathpakage.FoldSlice(s,
			func(x mathpakage.Element, y mathpakage.Element) mathpakage.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		mathpakage.FoldSlice(s,
			func(x mathpakage.Element, y mathpakage.Element) mathpakage.Element {
				return x + y
			},
			0))

}
