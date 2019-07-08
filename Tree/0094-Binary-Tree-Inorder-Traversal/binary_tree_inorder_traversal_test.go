package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("InorderTraversal(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, inorderTraversal(test.root))
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				inorderTraversal(test.root)
			}
		}
	})
}
