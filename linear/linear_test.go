package linear

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// intList creates a ordered list with n elements
func intList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] += i
	}
	return list
}

// shuffle Shuffles a list using Fisher-Yates shuffle
func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())

	for i := len(list) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func TestLinear(t *testing.T) {
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
		result := Find(search, list)

		if result != want[i] {
			t.Errorf("Got = %v; want %v", result, want[i])
		}

	}
}

func BenchmarkLinear(b *testing.B) {
	for _, size := range []int{
		100, 200, 400, 800, 1600, 3200,
	} {
		fmt.Println("Size:", size)
		for n := 0; n < b.N; n++ {
			b.StopTimer()

			list := rand.Perm(size)
			search := rand.Intn(size)

			b.StartTimer()
			_ = Find(search, list)
		}
	}
}
