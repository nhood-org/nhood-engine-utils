package utils

import "testing"

func TestMovingAvergeInitialState(t *testing.T) {
	avr := NewMovingAverage()

	if avr.sum != 0 {
		t.Errorf("Expected avr.sum to be 0, while it was %f", avr.sum)
	}
	if avr.count != 0 {
		t.Errorf("Expected avr.count to be 0, while it was %f", avr.count)
	}
	if avr.Count() != 0 {
		t.Errorf("Expected avr.Count() to return 0, while it was %f", avr.Count())
	}
	if avr.Avg() != 0 {
		t.Errorf("Expected avr.Avg() to return 0, while it was %f", avr.Avg())
	}
}

func TestMovingAverageCumulatesValuesCorrectly(t *testing.T) {
	avr := NewMovingAverage()

	avr.Add(1.0)
	if avr.Count() != 1 {
		t.Errorf("Expected avr.Count() to return 1, while it was %f", avr.Count())
	}
	if avr.Avg() != 1 {
		t.Errorf("Expected avr.Avg() to return 1, while it was %f", avr.Avg())
	}

	avr.Add(2.0)
	if avr.Count() != 2 {
		t.Errorf("Expected avr.Count() to return 2, while it was %f", avr.Count())
	}
	if avr.Avg() != 1.5 {
		t.Errorf("Expected avr.Avg() to return 1.5, while it was %f", avr.Avg())
	}

	avr.Add(3.0)
	if avr.Count() != 3 {
		t.Errorf("Expected avr.Count() to return 3, while it was %f", avr.Count())
	}
	if avr.Avg() != 2 {
		t.Errorf("Expected avr.Avg() to return 2, while it was %f", avr.Avg())
	}
}

func TestMovingAverageCumulatesNegativeValuesCorrectly(t *testing.T) {
	avr := NewMovingAverage()

	avr.Add(1.0)
	if avr.Count() != 1 {
		t.Errorf("Expected avr.Count() to return 1, while it was %f", avr.Count())
	}
	if avr.Avg() != 1 {
		t.Errorf("Expected avr.Avg() to return 1, while it was %f", avr.Avg())
	}

	avr.Add(-1.0)
	if avr.Count() != 2 {
		t.Errorf("Expected avr.Count() to return 2, while it was %f", avr.Count())
	}
	if avr.Avg() != 0 {
		t.Errorf("Expected avr.Avg() to return 0, while it was %f", avr.Avg())
	}
}
