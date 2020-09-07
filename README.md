# searching-algorithms

Searching Algorithms implemented in Go. In order to benchmark it was needed to add an idle timer of 1 microsecond, otherwise the function would run forever, e.g. [stackoverflow](https://stackoverflow.com/questions/37620251/golang-benchmarking-b-stoptimer-hangs-is-it-me).

## Linear Search

Time complexity: O(N)

### Benchmark

```
pkg: github.com/nahumsa/searching-algorithms/linear
BenchmarkLinear/100-4         	  192313	      6459 ns/op	       0 B/op	       0 allocs/op
BenchmarkLinear/200-4         	  165627	      7136 ns/op	       0 B/op	       0 allocs/op
BenchmarkLinear/400-4         	  152652	      8093 ns/op	       0 B/op	       0 allocs/op
BenchmarkLinear/800-4         	  106137	     11080 ns/op	       0 B/op	       0 allocs/op
BenchmarkLinear/1600-4        	   78976	     15583 ns/op	       0 B/op	       0 allocs/op
```

## Binary Search

Caveat: The list needs to be sorted in order to use binary search.

Time complexity: O(log N)

### Benchmark not sorted

```
pkg: github.com/nahumsa/searching-algorithms/binary
BenchmarkBinaryNotSorted/100-4         	   33446	     33955 ns/op	    5824 B/op	      99 allocs/op
BenchmarkBinaryNotSorted/200-4         	   23588	     49536 ns/op	   13440 B/op	     199 allocs/op
BenchmarkBinaryNotSorted/400-4         	   12111	     97784 ns/op	   30080 B/op	     399 allocs/op
BenchmarkBinaryNotSorted/800-4         	    5982	    169317 ns/op	   66688 B/op	     799 allocs/op
BenchmarkBinaryNotSorted/1600-4        	    3123	    386195 ns/op	  146944 B/op	    1599 allocs/op
```

### Benchmark Sorted

```
pkg: github.com/nahumsa/searching-algorithms/binary
BenchmarkBinarySorted/100-4         	  147075	      7158 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySorted/200-4         	  153988	      7780 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySorted/400-4         	  153265	      7578 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySorted/800-4         	  141639	      8347 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySorted/1600-4        	  130537	      9056 ns/op	       0 B/op	       0 allocs/op
```

## Interpolation Search

Caveat: The list needs to be sorted in order to use interpolation search.

Time complexity: O(log log N)

### Benchmark not sorted

```
pkg: github.com/nahumsa/searching-algorithms/interpolation
BenchmarkInterpolationNotSorted/100-4         	   36152	     33095 ns/op	    5824 B/op	      99 allocs/op
BenchmarkInterpolationNotSorted/200-4         	   25357	     48882 ns/op	   13440 B/op	     199 allocs/op
BenchmarkInterpolationNotSorted/400-4         	   10000	    104590 ns/op	   30080 B/op	     399 allocs/op
BenchmarkInterpolationNotSorted/800-4         	    6543	    201406 ns/op	   66688 B/op	     799 allocs/op
BenchmarkInterpolationNotSorted/1600-4        	    3337	    388910 ns/op	  146944 B/op	    1599 allocs/op
```

### Benchmark Sorted

```
pkg: github.com/nahumsa/searching-algorithms/interpolation
BenchmarkInterpolationSorted/100-4         	  199711	      5875 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterpolationSorted/200-4         	  197515	      6082 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterpolationSorted/400-4         	  212798	      5752 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterpolationSorted/800-4         	  210783	      5965 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterpolationSorted/1600-4        	  195752	      6033 ns/op	       0 B/op	       0 allocs/op
```