package main

import (
	"log"

	"algorithm"
)

func main() {
	pop := algorithm.NewPopulation(100)

	for !pop.Converged {
		pop.Iterate()
	}

	log.Printf("Population converged at generation %v\n", pop.Generation)

}
