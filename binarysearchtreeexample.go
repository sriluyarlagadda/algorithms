package main

import (
	"errors"
	"fmt"
	"github.com/sriluyarlagadda/datastructures"
)

var _ = fmt.Printf

type ComparableInt int

func (c ComparableInt) Compare(y datastructures.Comparer) (bool, error) {
	value, ok := y.(ComparableInt)
	if ok {
		if int(c) > int(value) {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, errors.New("The comparables are not of the same type")

}

func main() {
	tree := datastructures.NewTree()
	tree.Insert(ComparableInt(20))
	tree.Insert(ComparableInt(67))
	tree.Insert(ComparableInt(3))
	tree.Insert(ComparableInt(12))
	tree.Insert(ComparableInt(11))
	tree.Insert(ComparableInt(13))
	tree.Insert(ComparableInt(66))
	tree.Insert(ComparableInt(68))
	tree.Insert(ComparableInt(65))

	tree.Traverse()

	node := tree.Search(ComparableInt(67))
	if node != nil {
		print(node.NodeValue)
	}

	tree.Delete(node)
	tree.Traverse()

}
