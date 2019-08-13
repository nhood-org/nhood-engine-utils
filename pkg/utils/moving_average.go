package utils

/*
MovingAverage maintains an data for running statistics of numbers added to it

*/
type MovingAverage struct {
	sum   float64
	count float64
}

/*
NewMovingAverage return an pristine instance of MovingAverage

*/
func NewMovingAverage() *MovingAverage {
	return &MovingAverage{
		sum:   0,
		count: 0,
	}
}

/*
Add adds number to running statistics

*/
func (r *MovingAverage) Add(v float64) {
	r.count++
	r.sum += v
}

/*
Count return a size of all numbers added

*/
func (r *MovingAverage) Count() float64 {
	return r.count
}

/*
Avg return an average of all numbers added

*/
func (r *MovingAverage) Avg() float64 {
	if r.count == 0 {
		return 0
	}
	return r.sum / r.count
}
