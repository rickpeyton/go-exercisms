package clock

import "fmt"

// Declare version to test submission against
const (
	TestVersion   = 2
	hourlyMinutes = 60
	dailyMinutes  = 1440
)

// Clock represents the time in minutes on a 24 hour clock
type Clock int

// Time converts hours and minutes into minutes on a 24 hour clock
func Time(hours, minutes int) Clock {
	// Takes in hours and minutes and returns positive minutes
	// Ensures that time is greater than 0 and less than 1440
	time := (((hours*hourlyMinutes + minutes) % dailyMinutes) + dailyMinutes) % dailyMinutes
	return Clock(time)
}

// Add or subtract minutes from an existing clock to update 24 hour clock
func (clock Clock) Add(minutes int) Clock {
	return Time(0, int(clock)+minutes)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/hourlyMinutes, clock%hourlyMinutes)
}
