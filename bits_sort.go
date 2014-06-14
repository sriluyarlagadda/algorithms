package main

import (
	"fmt"
	"github.com/sriluyarlagadda/datastructures"
)

func main() {
	binaryRadixTree := datastructures.NewBinaryRadixTree()
	binaryRadixTree.Insert("1011")
	binaryRadixTree.Insert("010")
	binaryRadixTree.Insert("01")
	binaryRadixTree.Insert("011")
	binaryRadixTree.Insert("0010")
	binaryRadixTree.Insert("1010")
	binaryRadixTree.Insert("1000")
	binaryRadixTree.Insert("0101")
	binaryRadixTree.Traverse()
	sortedStrings := binaryRadixTree.Sort()
	fmt.Println("sorted data:", sortedStrings)
}
