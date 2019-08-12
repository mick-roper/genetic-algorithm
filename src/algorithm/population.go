package algorithm

import "log"

// Population is a collection of indviduals
type Population struct {
	Generation int
	Fitness    int
	Converged  bool
	members    []Individual
	fittest    *Individual
	secFittest *Individual
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
	p.Generation++
	var w *Individual

	for n, i := range p.members {
		i.calculateFitness(target)

		if p.fittest == nil || i.fitness > p.fittest.fitness {
			p.fittest = &i
		} else if p.secFittest == nil || i.fitness > p.secFittest.fitness {
			p.secFittest = &i
		}

		if w == nil || i.fitness < w.fitness {
			w = &i
		}
	}

	// check for convergence
	if p.Fitness == p.fittest.fitness {
		p.Converged = true
	}
	p.Fitness = p.fittest.fitness

	// merge the fittest
	c := p.fittest.CombineWith(p.secFittest)

	// replace the weakest chromosome
	w.chromosome = c
}

// Print the state of the population
func (p *Population) Print() {
	log.Printf("Generation: %v Fittest: %v", p.Generation, p.fittest.chromosome)
}
