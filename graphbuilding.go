package main

import (
	"fmt"
	"github.com/sriluyarlagadda/datastructures"
	"io/ioutil"
	"log"
)

var _ = fmt.Println

func main() {

	data, error := ioutil.ReadFile("graph.data")
	if error != nil {
		log.Fatal("error reading data")
	}
	graph, error := datastructures.NewGraph(10, false)
	if error != nil {
		log.Fatal("unable to create graph")
	}

	error = graph.ReadGraph(data)
	if error != nil {
		fmt.Println("error:", error)
	}

	graph.DisplayMatrix()

	graph.TraverseBFS(func(edge *(datastructures.Edge)) {
		fmt.Println("vertex:", edge.Y)
	})

	graph.FindPath(4)

}
