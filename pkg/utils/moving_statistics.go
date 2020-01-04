package utils

import "math"

/*
MovingStatistics maintains an data for running statistics of numbers added to it

*/
type MovingStatistics struct {
	sum   float64
	count int
	max   float64
	min   float64
}

/*
NewMovingStatistics return an pristine instance of MovingStatistics

*/
func NewMovingStatistics() *MovingStatistics {
	return &MovingStatistics{
		sum:   0,
		count: 0,
		max:   math.NaN(),
		min:   math.NaN(),
	}
}

/*
Add adds number to running statistics

*/
func (r *MovingStatistics) Add(v float64) {
	r.count++
	r.sum += v
	if math.IsNaN(r.max) || v > r.max {
		r.max = v
	}
	if math.IsNaN(r.min) || v < r.min {
		r.min = v
	}
}

/*
Count returns a size of all numbers added

*/
func (r *MovingStatistics) Count() int {
	return r.count
}

/*
Max returns maximum value all numbers added

*/
func (r *MovingStatistics) Max() float64 {
	return r.max
}

/*
Min returns minimum value all numbers added

*/
func (r *MovingStatistics) Min() float64 {
	return r.min
}

/*
Avg return an average of all numbers added

*/
func (r *MovingStatistics) Avg() float64 {
	if r.count == 0 {
		return 0
	}
	return r.sum / float64(r.count)
}
