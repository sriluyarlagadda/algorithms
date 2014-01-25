package main

import (
	"fmt"
	"github.com/sriluyarlagadda/datastructures"
)

func main() {
	queue := datastructures.NewQueue(10)
	queue.Enque(1)
	queue.Enque(2)
	queue.Enque("stringValue")

	fmt.Println("value:", queue.Deque())
	fmt.Println("value:", queue.Deque())
	fmt.Println("value:", queue.Deque())
}
