package conv

// SliToMap convert a slice or an array of structs to a map based on iteratee function
func SliToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = v
	}
	return result
}
