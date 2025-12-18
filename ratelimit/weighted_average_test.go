package ratelimit

import "testing"

/*
Unit tests for weighted average
*/

func TestWeightedAverage_Success(t *testing.T) {
	values := []float64{80, 90, 100}
	weights := []float64{0.2, 0.3, 0.5}

	result, err := WeightedAverage(values, weights)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 93.0
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestWeightedAverage_LengthMismatch(t *testing.T) {
	_, err := WeightedAverage([]float64{1, 2}, []float64{1})
	if err == nil {
		t.Fatal("expected error for mismatched lengths")
	}
}

func TestWeightedAverage_NegativeWeight(t *testing.T) {
	_, err := WeightedAverage(
		[]float64{10, 20},
		[]float64{1, -1},
	)
	if err == nil {
		t.Fatal("expected error for negative weight")
	}
}

func TestWeightedAverage_ZeroTotalWeight(t *testing.T) {
	_, err := WeightedAverage(
		[]float64{10, 20},
		[]float64{0, 0},
	)
	if err == nil {
		t.Fatal("expected error for zero total weight")
	}
}
