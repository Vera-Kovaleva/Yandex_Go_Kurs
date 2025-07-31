package main

import (
	"fmt"
)

func FindPairIndexsForSum(givenArray []int, sum int) []int {
	unicNumbersWithTheirIndex := make(map[int]int)

	for i, curNum := range givenArray {
		value, ok := unicNumbersWithTheirIndex[sum-curNum]
		if ok {
			return []int{i, value}
		}
		unicNumbersWithTheirIndex[curNum] = i
	}
	return nil
}

func RemoveDublicates(givenArr []string) []string {
	var result []string
	check := make(map[string]int)
	for _, el := range givenArr {
		_, ok := check[el]
		if !ok {
			result = append(result, el)
		}
		check[el]++
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 7, 2, 3}
	fmt.Println(FindPairIndexsForSum(nums, 6))
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemoveDublicates(input))

}
