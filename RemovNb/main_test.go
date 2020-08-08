package main

import (
	"fmt"
	"testing"
)

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// sumAllExpectTwoNumbers(100)
		RemovNb(26)
	}
}

func TestSumAllExpectTwoNumbers(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := uint64(351)
	result := sumAllExpectTwoNumbers(26)
	if expected != result {
		t.Error("Failed")
	}
}
