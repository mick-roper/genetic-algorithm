package main

import (
	"flag"
	"log"
	"time"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

var target = flag.String("target", "", "the target string we trying to reach")
var delayMs = flag.Int("delay", 25, "the amount of delay between iterations in milliseconds")
var popSize = flag.Int("population", 100, "the size of the population you want to use")

var start, diff int64

func main() {
	flag.Parse()

	if *target == "" {
		log.Fatal("a target must be provided")
	}

	interval := int64(*delayMs)
	pop := algorithm.NewPopulation(*popSize, len(*target))

	for pop.Fittest().Chromosome() != *target {
		start = time.Now().UnixNano()
		pop.Iterate(*target)
		pop.Print()

		diff = (time.Now().UnixNano() - start) / 100000

		<-time.After(time.Duration(interval-diff) * time.Nanosecond)
	}

	log.Println("Population completed its evolution at generation ", pop.Generation)
}
