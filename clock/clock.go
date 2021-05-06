package clock

import "fmt"

const OneDay = 24 * 60

type Clock struct {
	min int
}

func (c Clock) Add(m int) Clock {
	return New(0, c.min+m)
}
func (c Clock) Subtract(m int) Clock {
	return New(0, c.min-m)
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

func New(hour, min int) Clock {
	totalMinutes := (hour*60 + min) % OneDay
	if totalMinutes < 0 {
		totalMinutes = totalMinutes + OneDay
	}
	return Clock{min: totalMinutes}
}
