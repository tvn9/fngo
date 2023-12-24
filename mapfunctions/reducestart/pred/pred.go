package pred

type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type (
	reduceFunc[A any] func(a1, a2 A) A
)

func ReduceWithStart[A any](input []A, start A, reducer reduceFunc[A]) A {
	if len(input) == 0 {
		return start
	}
	if len(input) == 1 {
		return reducer(start, input[0])
	}
	result := reducer(start, input[0])
	for _, element := range input[1:] {
		result = reducer(result, element)
	}
	return result
}
