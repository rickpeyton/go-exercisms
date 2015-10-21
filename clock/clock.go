package clock

import "fmt"

// Declare version to test submission against
const TestVersion = 2

// Clock represents the time in minutes on a 24 hour clock
type Clock int

// Time converts hours and minutes into minutes on a 24 hour clock
func Time(hours, minutes int) Clock {
	// Takes in hours and minutes and returns minutes
	time := hours*60 + minutes
	for time >= 1440 {
		time -= 1440
	}
	for time < 0 {
		time += 1440
	}
	return Clock(time)
}

// Add or subtract minutes from an existing clock to update 24 hour clock
func (clock Clock) Add(minutes int) Clock {
	updateClock := clock + Clock(minutes)
	for updateClock >= 1440 {
		updateClock -= 1440
	}
	for updateClock < 0 {
		updateClock += 1440
	}
	return updateClock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}
