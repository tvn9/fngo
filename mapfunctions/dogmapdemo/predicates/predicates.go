package predicates

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

type (
	Name   string
	Breed  int
	Gender int
)

const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

const (
	Male Gender = iota
	Female
)

type MapFunc[A any] func(A) A

func Map[A any](input []A, m MapFunc[A]) []A {
	output := make([]A, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}
