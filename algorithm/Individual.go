package algorithm

import "strings"

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
		if target[n] == i.chromosome[n] {
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
	sb := strings.Builder{}
	l := len(i.chromosome)
	sb.Grow(l)

	for n := 0; n < l; n++ {
		p := r.Float64()

		if p < 0.45 { // take from this parent
			sb.WriteByte(i.chromosome[n])
		} else if p < 0.9 { // take from other parent
			sb.WriteByte(o.chromosome[n])
		} else { // mutate
			sb.WriteString(randomGene())
		}
	}

	return sb.String()
}

// Chromosome of the individual
func (i *Individual) Chromosome() string {
	return i.chromosome
}
