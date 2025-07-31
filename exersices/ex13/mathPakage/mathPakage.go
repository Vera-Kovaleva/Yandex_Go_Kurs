package mathpakage

type Slice []Element

type Element int

func Addition(firstSlice []int, secondSlice []int) []int {
	var res []int
	for i := 0; i < len(firstSlice); i++ {
		res = append(res, firstSlice[i]-secondSlice[i])
	}
	return res
}

func Subtraction(firstSlice []int, secondSlice []int) []int {
	var res []int
	for i := 0; i < len(firstSlice); i++ {
		res = append(res, firstSlice[i]-secondSlice[i])
	}
	return res
}

func Multiplication(firstSlice []int, secondSlice []int) []int {
	var res []int
	for i := 0; i < len(firstSlice); i++ {
		res = append(res, firstSlice[i]-secondSlice[i])
	}
	return res
}

func Division(firstSlice []int, secondSlice []int) []int {
	var res []int
	for i := 0; i < len(firstSlice); i++ {
		res = append(res, firstSlice[i]-secondSlice[i])
	}
	return res
}

func SumSlice(s Slice) Element {
	var sum Element
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func MapSlice(s Slice, op func(Element) Element) {
	for i := 0; i < len(s); i++ {
		s[i] = op(s[i])
	}

}

func FoldSlice(s Slice, op func(Element, Element) Element, init Element) Element {
	res := op(init, s[0])
	for i := 1; i < len(s); i++ {
		res = op(res, s[i])
	}
	return res
}
