package binary

import (
	"fmt"
	"math/rand"
	"testing"
)

// intList creates a ordered list with n elements
func intList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] += i
	}
	return list
}

func TestBinary(t *testing.T) {
	want := []bool{true, false, true, false, true}
	ints := []int{5, 10, 15, 20, 30}

	for i := 0; i < len(ints); i++ {
		// create list
		list := rand.Perm(ints[i])

		var search int
		if !want[i] {
			search = ints[i] + 5
		} else {
			search = rand.Intn(ints[i])
		}
		fmt.Println(search)
		result := Find(search, list, false)

		if result != want[i] {
			t.Errorf("Got = %v; want %v", result, want[i])
		}

	}
}

func BenchmarkBinaryNotSorted(b *testing.B) {
	for _, size := range []int{100, 200, 400, 800, 1600} {
		b.Run(fmt.Sprintf("%v", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				list := rand.Perm(size)
				search := rand.Intn(size)
				Find(search, list, false)
			}
		})
	}
}

func BenchmarkBinarySorted(b *testing.B) {
	for _, size := range []int{100, 200, 400, 800, 1600} {
		b.Run(fmt.Sprintf("%v", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				list := intList(size)
				search := rand.Intn(size)
				Find(search, list, true)
			}
		})
	}
}
