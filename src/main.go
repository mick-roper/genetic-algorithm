package main

import (
	"log"
	"time"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

const target = "Hello World"

func main() {
	pop := algorithm.NewPopulation(100, len(target))

	for pop.Fittest().Chromosome() != target {
		pop.Iterate(target)
		pop.Print()
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("Population completed its evolution at generation ", pop.Generation)
}
