package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var expected float64
	expected = 3.5

	if actual := Average([]int{1, 2, 3, 4, 5, 6}); expected != actual {
		t.Errorf("expected=%f, actual=%f", expected, actual)
	}
	if actual := Average([]int32{1, 2, 3, 4, 5, 6}); expected != actual {
		t.Errorf("expected=%f, actual=%f", expected, actual)
	}
	if actual := Average([]int64{1, 2, 3, 4, 5, 6}); expected != actual {
		t.Errorf("expected=%f, actual=%f", expected, actual)
	}
	if actual := Average([]float32{1, 2, 3, 4, 5, 6}); expected != actual {
		t.Errorf("expected=%f, actual=%f", expected, actual)
	}
	if actual := Average([]float64{1, 2, 3, 4, 5, 6}); expected != actual {
		t.Errorf("expected=%f, actual=%f", expected, actual)
	}
}
