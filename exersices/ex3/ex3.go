package main

import (
	"fmt"
)

func main() {
	var bornYear int
	fmt.Scan(&bornYear)
	switch {
	case bornYear > 2013:
		fmt.Println("Привет, альфа!")
	case bornYear > 1997 && bornYear < 2013:
		fmt.Println("Привет, зумер!")
	case bornYear > 1981 && bornYear < 1997:
		fmt.Println("Привет, миллениал!")
	case bornYear > 1965 && bornYear < 1981:
		fmt.Println("Привет, представитель X!")
	case bornYear > 1946 && bornYear < 1965:
		fmt.Println("Привет, бумер!")
	default:
        fmt.Println("Здравствуйте!")
	}
}
