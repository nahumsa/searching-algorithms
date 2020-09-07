package linear

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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
	for _, size := range []int{100, 200, 400, 800, 1600} {
		b.Run(fmt.Sprintf("%v", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				search := rand.Intn(size)
				b.StartTimer()
				Find(search, list)
				time.Sleep(1 * time.Microsecond)
			}
		})
	}
}
