package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingStatisticsInitialState(t *testing.T) {
	a := assert.New(t)

	stats := NewMovingStatistics()

	a.Equal(0.0, stats.sum)
	a.Equal(0, stats.count)
	a.True(math.IsNaN(stats.max))
	a.True(math.IsNaN(stats.min))

	a.Equal(0, stats.Count())
	a.Equal(0.0, stats.Avg())
	a.True(math.IsNaN(stats.Max()))
	a.True(math.IsNaN(stats.Min()))
}

func TestMovingStatisticsAccumulatesValuesCorrectly(t *testing.T) {
	a := assert.New(t)

	stats := NewMovingStatistics()

	stats.Add(1.0)
	a.Equal(1, stats.Count())
	a.Equal(1.0, stats.Avg())
	a.Equal(1.0, stats.Max())
	a.Equal(1.0, stats.Min())

	stats.Add(2.0)
	a.Equal(2, stats.Count())
	a.Equal(1.5, stats.Avg())
	a.Equal(2.0, stats.Max())
	a.Equal(1.0, stats.Min())

	stats.Add(3.0)
	a.Equal(3, stats.Count())
	a.Equal(2.0, stats.Avg())
	a.Equal(3.0, stats.Max())
	a.Equal(1.0, stats.Min())
}

func TestMovingStatisticsAccumulatesNegativeValuesCorrectly(t *testing.T) {
	a := assert.New(t)

	stats := NewMovingStatistics()

	stats.Add(1.0)
	a.Equal(1, stats.Count())
	a.Equal(1.0, stats.Avg())
	a.Equal(1.0, stats.Max())
	a.Equal(1.0, stats.Min())

	stats.Add(-1.0)
	a.Equal(2, stats.Count())
	a.Equal(0.0, stats.Avg())
	a.Equal(1.0, stats.Max())
	a.Equal(-1.0, stats.Min())
}
