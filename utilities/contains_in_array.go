package utilities

func ContainsInArray(from []string, search string) bool {
	for _, i := range from {
		if i == search {
			return true
		}
	}
	return false
}