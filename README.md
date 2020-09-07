# searching-algorithms

Searching Algorithms implemented in Go.

## Linear Search

Time complexity: O(N)

### Benchmark

```
pkg: github.com/nahumsa/searching-algorithms/linear
BenchmarkLinear/100-4         	  288826	      3797 ns/op	     896 B/op	       1 allocs/op
BenchmarkLinear/200-4         	  151480	      7355 ns/op	    1792 B/op	       1 allocs/op
BenchmarkLinear/400-4         	   78327	     15895 ns/op	    3200 B/op	       1 allocs/op
BenchmarkLinear/800-4         	   37466	     32543 ns/op	    6528 B/op	       1 allocs/op
BenchmarkLinear/1600-4        	   18969	     62736 ns/op	   13568 B/op	       1 allocs/op
```

## Binary Search

Caveat: The list needs to be sorted in order to use binary search.

Time complexity: O(log N)

### Benchmark not sorted

```
pkg: github.com/nahumsa/searching-algorithms/binary
BenchmarkBinaryNotSorted/100-4         	   77901	     17500 ns/op	    6720 B/op	     100 allocs/op
BenchmarkBinaryNotSorted/200-4         	   38032	     38535 ns/op	   15232 B/op	     200 allocs/op
BenchmarkBinaryNotSorted/400-4         	   15457	     76798 ns/op	   33280 B/op	     400 allocs/op
BenchmarkBinaryNotSorted/800-4         	    6958	    160245 ns/op	   73216 B/op	     800 allocs/op
BenchmarkBinaryNotSorted/1600-4        	    3697	    336715 ns/op	  160512 B/op	    1600 allocs/op
```

### Benchmark Sorted

```
pkg: github.com/nahumsa/searching-algorithms/binary
BenchmarkBinarySorted/100-4         	  118857	      9702 ns/op	    6720 B/op	     100 allocs/op
BenchmarkBinarySorted/200-4         	   62858	     20363 ns/op	   15232 B/op	     200 allocs/op
BenchmarkBinarySorted/400-4         	   28528	     41331 ns/op	   33280 B/op	     400 allocs/op
BenchmarkBinarySorted/800-4         	   13070	     95804 ns/op	   73216 B/op	     800 allocs/op
BenchmarkBinarySorted/1600-4        	    6045	    193689 ns/op	  160512 B/op	    1600 allocs/op
```