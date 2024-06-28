package utils

/*
Get the index of an element in a given array.
Parameters:
	element - string, int
	array - []string []int
Returns int:
	-1: element is non-existent
	index: element's index
*/
func IndexOf[T comparable](element T, array []T) int {
	for i, v := range array {
		if element == v {
			return i
		}
	}
	return -1
}