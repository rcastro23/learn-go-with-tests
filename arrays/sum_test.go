package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	/*
		Two options when adding this new test:

		* Break the existing API by changing the argument to Sum to be a slice rather
		than an array. When we do this, we will potentially ruin someone's day because
		our other test will no longer compile!

		* Create a new function
	*/

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
