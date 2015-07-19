package clock

import (
	"fmt"
)

const TestVersion = 2

type Clock struct {
	hour   int
	minute int
}

func Time(hour, minute int) Clock {
	clock := Clock{hour, minute}
	clock = clock.Sanitize()
	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	clock = clock.Sanitize()
	return clock
}

func (clock Clock) SanitizeHour() Clock {
	clock.hour %= 24
	if clock.hour < 0 {
		clock.hour += 24
	}

	return clock
}

func (clock Clock) SanitizeMinute() Clock {
	additionalHour := clock.minute / 60
	clock.minute %= 60
	if clock.minute < 0 {
		clock.minute += 60
		additionalHour -= 1
	}
	clock.hour += additionalHour

	return clock
}

func (clock Clock) Sanitize() Clock {
	clock = clock.SanitizeMinute()
	clock = clock.SanitizeHour()

	return clock
}
