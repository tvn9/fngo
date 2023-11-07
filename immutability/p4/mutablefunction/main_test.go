package main

import "testing"

func BenchmarkMutablePerson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutableCreatePerson()
	}
}
