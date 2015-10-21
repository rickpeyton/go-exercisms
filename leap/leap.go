package leap

// Declaring the version this submission should be tested against
const TestVersion = 1

// IsLeapYear takes in a 4 digit year and returns true if given year is a leap
// year or false if given year is not
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
