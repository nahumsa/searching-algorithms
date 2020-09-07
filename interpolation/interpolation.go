package interpolation

import (
	"github.com/nahumsa/sorting-algorithms/merge"
)

// Find returns true if a element is in the list
// and false if the element is not in the list.
// List needs to be sorted before the interpolation search
// therefore we use mergesort to sort the list.
// Thus the time complexity is O(N logN), because of the sort.
func Find(findValue int, list []int, sorted bool) bool {
	if !sorted {
		merge.SortInt(list)
	}

	low := 0
	high := len(list) - 1

	for list[low] < findValue && list[high] > findValue {
		mid := interpolate(list, low, high, findValue)

		if list[mid] < findValue {
			low = mid + 1
		} else if list[mid] > findValue {
			high = mid - 1
		} else {
			return true
		}

	}

	// Test edge cases when the value is on the border of the list

	if list[low] == findValue {
		return true
	} else if list[high] == findValue {
		return true
	} else {
		return false
	}
}

func interpolate(list []int, low, high, findValue int) int {
	return low + ((findValue-list[low])*(high-low))/(list[high]-list[low])
}
