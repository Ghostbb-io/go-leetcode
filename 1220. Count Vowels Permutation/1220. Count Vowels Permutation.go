package main

import "fmt"

func main() {
	fmt.Println(countVowelPermutation(5))
}

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	m := 1000000007

	for c := 2; c <= n; c++ {
		a, e, i, o, u = (e+i+u)%m, (a+i)%m, (e+o)%m, i, (i+o)%m
	}

	return (a + e + i + o + u) % m
}
