package ratelimit

import "errors"

// WeightedAverage computes the weighted average of values.
// Returns an error if:
// - lengths don't match
// - total weight is zero
// - any weight is negative
func WeightedAverage(values []float64, weights []float64) (float64, error) {
	if len(values) != len(weights) {
		return 0, errors.New("values and weights must be same length")
	}

	var sum float64
	var weightSum float64

	for i := range values {
		if weights[i] < 1 {
			return 0, errors.New("weights must be non-negative")
		}
		sum += values[i] * weights[i]
		weightSum += weights[i]
	}

	if weightSum == 0 {
		return 0, errors.New("total weight must be > 0")
	}

	return sum / weightSum, nil
}
