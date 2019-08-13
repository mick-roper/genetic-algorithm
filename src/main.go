package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/mick-roper/genetic-algorithm/src/algorithm"
)

var target = flag.String("target", "", "the target string we trying to reach")
var delayMs = flag.Int("delay", 25, "the amount of delay between iterations in milliseconds")
var popSize = flag.Int("population", 100, "the size of the population you want to use")

func main() {
	flag.Parse()

	if *target == "" {
		log.Fatal("a target must be provided")
	}

	go func() {
		for {
			printMemUsage()
			time.Sleep(5 * time.Second)
		}
	}()

	interval := time.Duration(*delayMs) * time.Millisecond
	pop := algorithm.NewPopulation(*popSize, len(*target))

	for pop.Fittest().Chromosome() != *target {
		pop.Iterate(*target)
		pop.Print()

		time.Sleep(interval)
	}

	log.Println("Population completed its evolution at generation ", pop.Generation)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
