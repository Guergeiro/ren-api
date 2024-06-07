package interval

import (
	"slices"
	"time"
)

type Interval struct {
	start time.Time
	stop  time.Time
}

func NewIntervals(start, stop time.Time) []Interval {
	maximumStop := start.AddDate(0, 1, 0)
	isSmallerThanMaximum := maximumStop.Compare(stop) >= 0
	if isSmallerThanMaximum {
		return []Interval{
			NewInterval(start, stop),
		}
	}

	newStart := maximumStop.AddDate(0, 0, 1)
	remainder := stop.Sub(newStart)
	newStop := newStart.Add(remainder)
	return slices.Concat(
		NewIntervals(start, maximumStop),
		NewIntervals(newStart, newStop),
	)
}

func NewInterval(start, stop time.Time) Interval {
	return Interval{
		start,
		stop,
	}
}

func (i Interval) StartTime() time.Time {
	return i.start
}
func (i Interval) StopTime() time.Time {
	return i.stop
}

func (i1 Interval) Equal(i2 Interval) bool {
	return i1.Compare(i2) == 0
}

func (i1 Interval) Compare(i2 Interval) int {
	if cmp := i1.start.Compare(i2.start); cmp != 0 {
		return cmp
	}
	return i1.stop.Compare(i2.stop)
}
