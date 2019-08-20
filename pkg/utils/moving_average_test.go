package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAvergeInitialState(t *testing.T) {
	assert := assert.New(t)

	avr := NewMovingAverage()

	assert.Equal(avr.sum, 0.0)
	assert.Equal(avr.count, 0.0)
	assert.Equal(avr.Count(), 0.0)
	assert.Equal(avr.Avg(), 0.0)
}

func TestMovingAverageCumulatesValuesCorrectly(t *testing.T) {
	assert := assert.New(t)

	avr := NewMovingAverage()

	avr.Add(1.0)
	assert.Equal(avr.Count(), 1.0)
	assert.Equal(avr.Avg(), 1.0)

	avr.Add(2.0)
	assert.Equal(avr.Count(), 2.0)
	assert.Equal(avr.Avg(), 1.5)

	avr.Add(3.0)
	assert.Equal(avr.Count(), 3.0)
	assert.Equal(avr.Avg(), 2.0)
}

func TestMovingAverageCumulatesNegativeValuesCorrectly(t *testing.T) {
	assert := assert.New(t)

	avr := NewMovingAverage()

	avr.Add(1.0)
	assert.Equal(avr.Count(), 1.0)
	assert.Equal(avr.Avg(), 1.0)

	avr.Add(-1.0)
	assert.Equal(avr.Count(), 2.0)
	assert.Equal(avr.Avg(), 0.0)
}
