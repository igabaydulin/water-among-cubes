package main

import (
	"testing"
)

func TestFillPool(t *testing.T) {
	input := [][]float64{
		{78, 16, 94, 36},
		{87, 93, 50, 22},
		{63, 28, 91, 60},
		{64, 27, 41, 27},
		{73, 37, 12, 69},
		{68, 30, 83, 31},
		{63, 24, 68, 36}}
	expected := 44
	actual := fillPool(input)

	if expected != actual {
		t.Fail()
	}
}
