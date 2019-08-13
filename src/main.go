package main

import (
	"log"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

const target = "Hello World"

func main() {
	pop := algorithm.NewPopulation(100, len(target))

	for true {
		pop.Iterate(target)
		pop.Print()
	}

	log.Printf("Population converged at generation %v\n", pop.Generation)
}
