package main

import (
	"fmt"
)

func main() {
	pricelist := map[string]int{"хлеб": 50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500}

	var delicates []string

	zakaz := []string{"хлеб", "буженина", "сыр", "огурцы"}
	cost := 0

	for k, v := range pricelist {
		if v >= 500 {
			delicates = append(delicates, k)
		}
	}

	for _, v := range zakaz {
		cost += pricelist[v]
	}
	fmt.Println("Деликатесы:", delicates, ", стоимость заказа: ", cost)
}
