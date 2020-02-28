package cmd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	testFile := "../testdata/test.txt"
	n := "5"
	expected := []int{100, 90, 80, 70, 60}
	output := findTop(testFile, n)

	if !reflect.DeepEqual(expected, output[1:]) {
		t.Errorf("Expected: %v, Output: %v", expected, output[1:])
	}
	fmt.Println("expected: ", expected)
	fmt.Println("output: ", output[1:])
}
