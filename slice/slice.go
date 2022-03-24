//Utility packge provide developer general functions that not tied to this project
package slice

// Contains checks if a string is present in a slice
// While this function works well enough, it may be inefficient for larger slices.
// In such cases, consider using a map instead of a slice as you won’t have to iterate over the
// entire list just to check for the existence of a value.
func StringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
