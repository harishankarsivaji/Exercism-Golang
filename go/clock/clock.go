package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

//New creates a new clock
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 60 * 24
	if m < 0 {
		m += 24 * 60
	}
	return Clock{hour: m / 60, minute: m % 60}
}

//Add - adds minute to existing time
func (m Clock) Add(minute int) Clock {
	return New(m.hour, m.minute+minute)
}

//Subtract - minus the input from current time
func (m Clock) Subtract(minute int) Clock {
	return New(m.hour, m.minute-minute)
}

//String returns the current time as string
func (m Clock) String() string {
	return fmt.Sprintf("%02d:%02d", m.hour, m.minute)
}
