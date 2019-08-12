package algorithm

import "log"

// Population is a collection of indviduals
type Population struct {
	generation int
	members    []Individual
}

// NewPopulation of the given size
func NewPopulation(size int) *Population {
	m := make([]Individual, size)

	for i := range m {
		m[i] = *newIndividual()
	}

	return &Population{}
}

// Iterate the population by one generation
func (p *Population) Iterate(target string) {
	p.generation++
	var f1, f2, w *Individual

	for _, i := range p.members {
		i.calculateFitness(target)

		if f1 == nil || i.fitness > f1.fitness {
			f1 = &i
		} else if f2 == nil || i.fitness > f2.fitness {
			f2 = &i
		}

		if w == nil || i.fitness < w.fitness {
			w = &i
		}
	}

	// merge the fittest
	c := f1.CombineWith(f2)

	// replace the weakest chromosome
	w.chromosome = c

	log.Printf("Generation: %v Fittest: %v", p.generation, f1.chromosome)
}
