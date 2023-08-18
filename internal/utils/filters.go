package utils

// Filter filters a slice based on a predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
    filtered := make([]T, 0)
    for _, item := range slice {
        if predicate(item) {
            filtered = append(filtered, item)
        }
    }
    return filtered
}
