package algorithm

import (
	"log"
	"sort"
)

// Population is a collection of indviduals
type Population struct {
	Generation int
	members    []*Individual
	length     int
}

// NewPopulation of the given size
func NewPopulation(size, length int) *Population {
	m := make([]*Individual, size)

	for i := range m {
		m[i] = newIndividual(length)
	}

	return &Population{
		Generation: 0,
		length:     length,
		members:    m,
	}
}

// Iterate the population by one generation
func (p *Population) Iterate(target string) {
	for _, i := range p.members {
		i.calculateFitness(target)
	}

	// order the members
	sort.Slice(p.members, func(a, b int) bool {
		return p.members[a].fitness > p.members[b].fitness
	})

	// updated the members
	size := len(p.members)
	bMin := int(float32(size)*0.1) - 1
	bMax := int(float32(size)*0.5) - 1
	breeders := p.members[:bMax]
	bLength := len(breeders)

	for i := range p.members {
		if i <= bMin { // elites
			// allow these to survive to the nxt generation
		} else if i <= bMax { // breeders
			// randomly pick some parents
			a := breeders[r.Intn(bLength)]
			b := breeders[r.Intn(bLength)]
			p.members[i] = &Individual{
				chromosome: a.CombineWith(b),
			}
		} else { // new members
			p.members[i] = newIndividual(p.length)
		}
	}

	p.Generation++
}

// Print the state of the population
func (p *Population) Print() {
	log.Printf("Generation: %v Fittest: %v", p.Generation, p.members[0].chromosome)
}

// Fittest member of the population
func (p *Population) Fittest() *Individual {
	return p.members[0]
}
