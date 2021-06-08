package main

import (
	"flag"
	"log"
	"time"

	"github.com/mick-roper/genetic-algorithm/algorithm"
)

var (
	target  string
	delayMs int
	popSize int
)

func init() {
	flag.StringVar(&target, "target", "", "the target strign we are trying to generate")
	flag.IntVar(&delayMs, "delay", 20, "the amount of delay between iterations in milliseconds")
	flag.IntVar(&popSize, "population", 20, "the size of the population you want to use")

	flag.Parse()
}

func main() {
	flag.Parse()

	if target == "" {
		log.Fatal("a target must be provided")
	}

	interval := time.Duration(delayMs) * time.Millisecond
	population := algorithm.NewPopulation(popSize, len(target))

	for population.Fittest().Chromosome() != target {
		population.Iterate(target)
		population.Print()

		time.Sleep(interval)
	}

	log.Println("Population completed its evolution at generation ", population.Generation)
}
