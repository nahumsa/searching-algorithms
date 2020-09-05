package linear

// Find returns true if a element is in the list
// and false if the element is not in the list.
func Find(findValue int, list []int) bool {
	for _, listValue := range list {
		if listValue == findValue {
			return true
		}
	}

	return false
}
