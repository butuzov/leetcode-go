package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("NextPermutation(%d)", i), func(t *testing.T) {
			nextPermutation(test.nums)
			assert.Equal(t, test.expected, test.nums)
		})
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				nextPermutation(test.nums)
			}
		}
	})
}
