package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverageInitialState(t *testing.T) {
	a := assert.New(t)

	avr := NewMovingAverage()

	a.Equal(avr.sum, 0.0)
	a.Equal(avr.count, 0.0)
	a.Equal(avr.Count(), 0.0)
	a.Equal(avr.Avg(), 0.0)
}

func TestMovingAverageAccumulatesValuesCorrectly(t *testing.T) {
	a := assert.New(t)

	avr := NewMovingAverage()

	avr.Add(1.0)
	a.Equal(avr.Count(), 1.0)
	a.Equal(avr.Avg(), 1.0)

	avr.Add(2.0)
	a.Equal(avr.Count(), 2.0)
	a.Equal(avr.Avg(), 1.5)

	avr.Add(3.0)
	a.Equal(avr.Count(), 3.0)
	a.Equal(avr.Avg(), 2.0)
}

func TestMovingAverageAccumulatesNegativeValuesCorrectly(t *testing.T) {
	a := assert.New(t)

	avr := NewMovingAverage()

	avr.Add(1.0)
	a.Equal(avr.Count(), 1.0)
	a.Equal(avr.Avg(), 1.0)

	avr.Add(-1.0)
	a.Equal(avr.Count(), 2.0)
	a.Equal(avr.Avg(), 0.0)
}
