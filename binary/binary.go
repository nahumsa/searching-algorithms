package binary

import (
	"sort"

	"github.com/nahumsa/sorting-algorithms/merge"
)

// Find returns true if a element is in the list
// and false if the element is not in the list.
// List needs to be sorted before the binary search
// therefore we use mergesort to sort the list.
// Thus the time complexity is O(N logN), because of the sort.
func Find(findValue int, list []int) bool {
	merge.SortInt(list)
	i := sort.Search(len(list), func(i int) bool { return list[i] >= findValue })
	if i < len(list) && list[i] == findValue {
		return true
	}
	return false
}
