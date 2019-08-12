package algorithm

import (
	"log"
	"sort"
)

// Population is a collection of indviduals
type Population struct {
	Generation int
	members    []Individual
}

// NewPopulation of the given size
func NewPopulation(size int) *Population {
	m := make([]Individual, size)

	for i := range m {
		m[i] = *newIndividual()
	}

	return &Population{
		members: m,
	}
}

// Iterate the population by one generation
func (p *Population) Iterate(target string) {
	p.Generation++

	for _, i := range p.members {
		i.calculateFitness(target)
	}

	sort.Slice(p.members, func(a, b int) bool {
		return p.members[a].fitness > p.members[b].fitness
	})

	size := float64(len(p.members))

	// ignore the top 10% allowing them into the next generation

	// next 50% breed
	for i := int(size * 0.1); i < int(size*0.60); i++ {
		f1 := &p.members[i]
		i++
		f2 := &p.members[i]

		p.members = append(p.members, Individual{chromosome: f1.CombineWith(f2)})
	}

	// replace whatever is left with new individuals
	for i := int(size * 0.6); i < int(size); i++ {
		p.members[i] = *newIndividual()
	}
}

// Print the state of the population
func (p *Population) Print() {
	log.Printf("Generation: %v Fittest: %v", p.Generation, p.members[0].chromosome)
}
