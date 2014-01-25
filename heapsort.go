package main

import (
	"fmt"
	"github.com/sriluyarlagadda/datastructures"
)

func main() {
	fmt.Println("size", 20)
	queue := datastructures.NewPriorityQueue(20)
	queue.Insert(4)
	queue.Insert(20)
	queue.Insert(30)
	queue.Insert(17)
	queue.Insert(21)
	queue.Insert(28)
	fmt.Println("queue size", queue.Capacity())
	capacity := queue.Size()
	for i := 0; i < capacity; i++ {
		value := queue.GetMinimum()
		fmt.Println(value)
	}
	queue.GetMinimum()

}
