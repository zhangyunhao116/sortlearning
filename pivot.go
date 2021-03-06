package sortlearning

import "golang.org/x/exp/constraints"

// choosePivotV1 chooses a pivot in `v` and returns the index and `true` if the slice is likely already sorted.
//
// Elements in `v` might be reordered in the process.
//
// [0,8): chooses a static pivot.
// [8,ShortestNinther): uses the simple median-of-three method.
// [ShortestNinther,∞): uses the Tukey’s ninther method.
func choosePivotV1[T constraints.Ordered](v []T) (pivotidx int) {
	const (
		// shortestNinther is the minimum length to choose the Tukey’s ninther method.
		// Shorter slices use the simple median-of-three method.
		shortestNinther = 50
		// maxSwaps is the maximum number of swaps that can be performed in this function.
		maxSwaps = 4 * 3
	)

	l := len(v)

	var (
		// Counts the total number of swaps we are about to perform while sorting indices.
		swaps int
		// Three indices near which we are going to choose a pivot.
		a = l / 4 * 1
		b = l / 4 * 2
		c = l / 4 * 3
	)

	if l >= 8 {
		if l >= shortestNinther {
			// Find medians in the neighborhoods of `a`, `b`, and `c`.
			a = sortAdjacent(v, a, &swaps)
			b = sortAdjacent(v, b, &swaps)
			c = sortAdjacent(v, c, &swaps)
		}
		// Find the median among `a`, `b`, and `c`.
		b = sort3(v, a, b, c, &swaps)
	}
	return b
}

// sort2 swaps `a` `b` so that `v[a] <= v[b]`.
func sort2[T constraints.Ordered](v []T, a, b int, swaps *int) (int, int) {
	if v[b] < v[a] {
		*swaps++
		return b, a
	}
	return a, b
}

// sort3 swaps `a` `b` `c` so that `v[a] <= v[b] <= v[c]`, then return `b`.
func sort3[T constraints.Ordered](v []T, a, b, c int, swaps *int) int {
	a, b = sort2(v, a, b, swaps)
	b, c = sort2(v, b, c, swaps)
	a, b = sort2(v, a, b, swaps)
	return b
}

// sortAdjacent finds the median of `v[a - 1], v[a], v[a + 1]` and stores the index into `a`.
func sortAdjacent[T constraints.Ordered](v []T, a int, swaps *int) int {
	return sort3(v, a-1, a, a+1, swaps)
}

func reverseRange[T constraints.Ordered](v []T) {
	i := 0
	j := len(v) - 1
	for i < j {
		v[i], v[j] = v[j], v[i]
		i++
		j--
	}
}
