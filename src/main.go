package main

import (
	"flag"
	"log"
	"time"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

var target = flag.String("target", "", "the target string we trying to reach")

func main() {
	flag.Parse()

	if *target == "" {
		log.Fatal("a target must be provided")
	}

	pop := algorithm.NewPopulation(100, len(*target))

	for pop.Fittest().Chromosome() != *target {
		pop.Iterate(*target)
		pop.Print()
		time.Sleep(20 * time.Millisecond)
	}

	log.Println("Population completed its evolution at generation ", pop.Generation)
}
