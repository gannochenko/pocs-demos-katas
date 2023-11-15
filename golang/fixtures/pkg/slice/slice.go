package slice

func Merge[T comparable](slice1, slice2 []T) []T {
	// Combine the slices
	mergedSlice := append(slice1, slice2...)

	// Create a map to track unique elements
	uniqueElements := make(map[T]bool)

	// Create a result slice for unique elements
	var result []T

	// Iterate through the merged slice
	for _, value := range mergedSlice {
		// If the element is not in the map, add it to the result slice and map
		if _, ok := uniqueElements[value]; !ok {
			uniqueElements[value] = true
			result = append(result, value)
		}
	}

	return result
}
