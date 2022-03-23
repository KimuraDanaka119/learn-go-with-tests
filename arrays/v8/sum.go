package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, 0, add)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, []int{}, sumTail)
}

func Reduce[A any](collection []A, initialValue A, f func(A, A) A) A {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
