package utils

import (
	"fmt"
)

func NotePMIndices(input []string) [][]int {
	var subIndice []int
	var outPut [][]int
	for i, rn := range input {
		if rn == "'" {
			subIndice = append(subIndice, i)
			if len(subIndice) == 2 {
				outPut = append(outPut, subIndice)
				subIndice = []int{}
			}
		}
	}
	return outPut
}

func FixPunctMarks(input []string, indices [][]int) {
	clean := input
	indexLen := len(indices)
	inputLen := len(input)

	for i := 0; i < indexLen; i++ {
		currentSlice := indices[i]
		var index int
		var firstMid, secondMid int

		first, second := IndicesValues(currentSlice)

		if second-first == 2 {
			index = GetIndex(second)
		} else if second-first > 2 {
			firstMid, secondMid = GetIndexForMultipleMid(first, second)
		}

		fmt.Println(index)

		for i := inputLen - 1; i >= 1; i-- {
			if second-first == 2 {
				if i == index {
					ModifyMid(&clean[i])
					continue
				}
			} else if second-first > 2 {
				switch i {
				case firstMid:
					ModifyFirstMid(&clean[i])
					continue
				case secondMid:
					ModifySecondMid(&clean[i])
					continue
				}
			}
		}
	}

	var output []string

	for i := 0; i < len(clean); i++ {
		if clean[i] != "'" {
			output = append(output, clean[i])
		}
	}

	fmt.Println(clean)
	fmt.Println(output)
}

func GetIndex(second int) int {
	mid := second - 1

	return mid
}

func GetIndexForMultipleMid(first, second int) (int, int) {
	firstMid := first + 1
	secondMid := second - 1

	return firstMid, secondMid
}

func IndicesValues(input []int) (int, int) {
	var a, b int

	for i, val := range input {
		if i == 0 {
			a = val
		} else {
			b = val
		}
	}
	return a, b
}

func ModifyMid(input *string) {
	temp := *input

	temp = "'" + temp + "'"

	*input = temp
}

func ModifyFirstMid(input *string) {
	temp := *input

	temp = "'" + temp

	*input = temp
}

func ModifySecondMid(input *string) {
	temp := *input

	temp += "'"

	*input = temp
}
