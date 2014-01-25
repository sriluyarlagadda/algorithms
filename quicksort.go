package main

import "fmt"

func main() {
	items := []int{1, 7, 3, 9, 10, 2, 15, 21, 12}
	quickSort(items, 0, len(items)-1)
	fmt.Println(items)
}

//quick sort
func quickSort(items []int, initial int, final int) {

	if initial < final {
		middlePosition := partition(items, initial, final)
		quickSort(items, initial, middlePosition-1)
		quickSort(items, middlePosition+1, final)
	}

}

func partition(items []int, initial int, final int) int {
	pivot := items[final]
	j := initial - 1
	for i := initial; i < final-1; i++ {
		if items[i] < pivot {
			j = j + 1
			swap(items, i, j)
		}
	}

	swap(items, j+1, final)
	return j + 1
}

func swap(items []int, first int, second int) {
	items[first], items[second] = items[second], items[first]
}
