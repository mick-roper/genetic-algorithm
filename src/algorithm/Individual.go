package algorithm

// Individual member of a population
type Individual struct {
	fitness    int
	chromosome string
}

func newIndividual(genes int) *Individual {
	return &Individual{
		fitness:    0,
		chromosome: createGenome(genes),
	}
}

func (i *Individual) calculateFitness(target string) {
	i.fitness = 0
	l := len(target)

	for n := 0; n < l; n++ {
		if target[n] != i.chromosome[n] {
			i.fitness++
		}
	}
}

// IsFitterThan tests if an individual is fittern than another individual
func (i *Individual) IsFitterThan(other *Individual) bool {
	return i.fitness > other.fitness
}

// CombineWith another individual to create a child that shares their attributes
func (i *Individual) CombineWith(o *Individual) string {
	var childChromosome string
	var l int
	il, ol := len(i.chromosome), len(o.chromosome)

	if l = il; ol > il {
		l = ol
	}

	for n := 0; n < l; n++ {
		p := r.Float64()

		if p < 0.45 && n < il { // take from this parent
			childChromosome += string(i.chromosome[n])
		} else if p < 0.9 && n < ol { // take from other parent
			childChromosome += string(o.chromosome[n])
		} else { // mutate
			childChromosome += randomGene()
		}
	}

	return childChromosome
}

// Chromosome of the individual
func (i *Individual) Chromosome() string {
	return i.chromosome
}
