package main

import (
	"fmt"
	"os"
	"strconv"
)

var _ = os.O_APPEND
var _ = strconv.IntSize
var _ = fmt.Println

/*func main() {
	inputListString := os.Args[1:]

	var inputListInt []int

	for _, i := range inputListString {
		value, _ := strconv.ParseInt(i, 0, 64)
		inputListInt = append(inputListInt, int(value))
	}

	fmt.Println("value", len(inputListInt))

	intputList := sort(inputListInt)
	fmt.Println(intputList)
}*/

//sort using insertion sort
func sort(inputList []int) []int {

	for i := 1; i < len(inputList); i++ {
		k := i
		for j := i - 1; j >= 0; j-- {
			if inputList[k] < inputList[j] {
				temp := inputList[k]
				inputList[k] = inputList[j]
				inputList[j] = temp

			} else {
				break
			}
			k = j

		}
	}

	return inputList
}
