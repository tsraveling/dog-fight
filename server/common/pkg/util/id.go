package util

// IsValidCUID checks that the id is 25 characters long and starts with 'c'.
func IsValidCUID(id string) bool {
	return len(id) == 25 && id[0] == 'c'
}