package main

import (
	"fmt"
	"testing"
)

func BenchmarkCariPrime(t *testing.B) {
	fmt.Println(t.N)
	for i := 0; i < t.N; i++ {
		cariPrime(i)
	}
}
