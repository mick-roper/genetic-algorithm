package main

import (
	"flag"
	"log"
	"time"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

var target = flag.String("target", "", "the target string we trying to reach")
var delayMs = flag.Int("delay", 25, "the amount of delay between iterations in milliseconds")
var popSize = flag.Int("population", 20, "the size of the population you want to use")

func main() {
	flag.Parse()

	if *target == "" {
		log.Fatal("a target must be provided")
	}

	interval := time.Duration(*delayMs) * time.Millisecond
	pop := algorithm.NewPopulation(*popSize, len(*target))

	for pop.Fittest().Chromosome() != *target {
		pop.Iterate(*target)
		pop.Print()

		time.Sleep(interval)
	}

	log.Println("Population completed its evolution at generation ", pop.Generation)
}
