package arr

// using go generic (go version ^1.18 )

// Map applies a function to each element of a slice and returns a new slice
func Map[T any, U any](input []T, f func(T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = f(v)
	}
	return output
}

// Filter returns a slice of elements matching a f
func Filter[T any](input []T, f func(T) bool) []T {
	output := make([]T, 0, len(input)) // allocate with capacity but 0 length

	for _, v := range input {
		if isTrue := f(v); isTrue {
			output = append(output, v)
		}
	}
	return output
}

// Find returns the first element matching a f, or zero value if not found
func Find[T any](input []T, f func(T) bool) (T, bool) {
	var zero T

	for _, v := range input {
		if f(v) {
			zero = v
			return zero, true
		}
	}

	return zero, false
}

func FindIndex[T any](input []T, f func(T) bool) (int, bool) {
	var index = -1

	return index, false
}

// Reduce reduces a slice to a single value using a reducer function
func Reduce[T any, U any](input []T, initial U, reducer func(U, T) U) U {
	return initial
}

// Some returns true if any element satisfies the f
func Some[T any](input []T, f func(T) bool) bool {
	return false
}

// Every returns true if all elements satisfy the f
func Every[T any](input []T, f func(T) bool) bool {
	return false
}

// Includes returns true if the slice contains the specified element
func Includes[T comparable](input []T, value T) bool {
	return false
}
