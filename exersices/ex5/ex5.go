/*
func main() {
	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Printf("%d",*p)
  }*/

package main

import (
	"fmt"
)

func f(c *int) {
	*c++
}

func main() {
	/*
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}*/
	
	weekTempArr := [7]int{1, 2, 3, 5, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
}
