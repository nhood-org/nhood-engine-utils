package relevance

type Vector struct {
	weight int
	vector []int
}

func NewVector(size int) *Vector {
	return &Vector{
		weight: 0,
		vector: make([]int, size),
	}
}

func (t *Vector) increment(i int) {
	relevance := t.vector[i]
	relevance++
	t.vector[i] = relevance
	t.weight++
}
