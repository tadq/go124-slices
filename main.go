package main

import (
	"fmt"
	"slices"
)

func main() {
	repeat()
}

func getAll() {
	vals := []string{"one", "two", "three"}

	for i, v := range slices.All(vals) {
		fmt.Println(i, v)
	}

	for i, v := range slices.Backward(vals) {
		fmt.Println(i, v)
	}
}

func appendSeq() {
	seq := func(yield func(int) bool) {
		for i := range 4 {
			if !yield(i) {
				return
			}
		}
	}

	s := slices.AppendSeq([]int{1, 2, 3}, seq)
	fmt.Println(s)
}

func makeSlice() []string {
	return []string{"one", "two", "three", "four"}
}

// Perform binary search.
func binarySearch() {
	s := slices.Clone(makeSlice())
	slices.Sort(s)
	n, found := slices.BinarySearch(s, "four")
	fmt.Println(n, found)
	nn, nfound := slices.BinarySearch(s, "zero")
	fmt.Println(nn, nfound)
}

func clip() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:4:10]
	fmt.Println(cap(s))
	clip := slices.Clip(s)
	fmt.Println(cap(clip))
	fmt.Println(clip)
}

func iterTry() {
	seq := func(yield func(i int) bool) {
		for i := range 8 {
			if !yield(i) {
				return
			}
		}
	}

	for n := range seq {
		fmt.Println(n)
	}
}

func compact() {
	items := []int{1, 1, 1, 2, 3, 4, 5, 5, 5}
	s := slices.Compact(items)
	fmt.Println(s)
	fmt.Println(cap(s))
	s = slices.Clip(s)
	fmt.Println(cap(s))

	ret := slices.Concat(items, s)
	fmt.Println(ret)
}

func repeat() {
	items := []string{"a", "b", "c"}
	s := slices.Repeat(items, 3)
	fmt.Println(s)
}
