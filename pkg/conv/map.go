package conv

// MapToSlice convert a map to a slice based on iteratee function
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T {
	result := make([]T, 0, len(aMap))

	for k, v := range aMap {
		result = append(result, iteratee(k, v))
	}

	return result
}
