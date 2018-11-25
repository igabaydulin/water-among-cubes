package main

import (
	"testing"
)

func TestFillPoolFromExample(t *testing.T) {
	input := [][]float64{
		{3, 3, 4, 4, 4, 2},
		{3, 1, 3, 2, 1, 4},
		{7, 3, 1, 6, 4, 1}}
	expected := 5
	actual := fillPool(input)

	if expected != actual {
		t.Fail()
	}
}

func TestFillPoolSimple(t *testing.T) {
	input := [][]float64{
		{2, 2, 2},
		{2, 1, 2},
		{2, 2, 2}}
	expected := 1
	actual := fillPool(input)

	if expected != actual {
		t.Fail()
	}
}

func TestFillPoolEmpty(t *testing.T) {
	input := [][]float64{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2}}
	expected := 0
	actual := fillPool(input)

	if expected != actual {
		t.Fail()
	}
}

func TestFillPoolHighGround(t *testing.T) {
	input := [][]float64{
		{2, 2, 2},
		{2, 3, 2},
		{2, 2, 2}}
	expected := 0
	actual := fillPool(input)

	if expected != actual {
		t.Fail()
	}
}

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
