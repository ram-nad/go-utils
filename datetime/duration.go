package datetime

import "time"

/*
Duration can be used to represent an amount of time. It can be positive or negative.

Internally it is stored as number of nanoseconds.
*/
type Duration int64

func ParseDuration(s string) (Duration, error) {
	d, err := time.ParseDuration(s)
	return Duration(int64(d)), err
}

func Nanoseconds(ns int64) Duration {
	return Duration(ns)
}

func Microseconds(us int64) Duration {
	return Duration(us * int64(time.Microsecond))
}

func Milliseconds(ms int64) Duration {
	return Duration(ms * int64(time.Millisecond))
}

func Seconds(s int64) Duration {
	return Duration(s * int64(time.Second))
}

func Minutes(m int64) Duration {
	return Duration(m * int64(time.Minute))
}

func Hours(h int64) Duration {
	return Duration(h * int64(time.Hour))
}

// Nanoseconds returns the duration as a number of nanoseconds (truncate towards zero)
func (d Duration) Nanoseconds() int64 {
	return int64(d)
}

// Microseconds returns the duration as a number of microseconds (truncate towards zero)
func (d Duration) Microseconds() int64 {
	return int64(d) / int64(time.Microsecond)
}

// Milliseconds returns the duration as a number of milliseconds (truncate towards zero)
func (d Duration) Milliseconds() int64 {
	return int64(d) / int64(time.Millisecond)
}

// Seconds returns the duration as a number of seconds (truncate towards zero)
func (d Duration) Seconds() int64 {
	return int64(d) / int64(time.Second)
}

// Minutes returns the duration as a number of minutes (truncate towards zero)
func (d Duration) Minutes() int64 {
	return int64(d) / int64(time.Minute)
}

// Hours returns the duration as a number of hours (truncate towards zero)
func (d Duration) Hours() int64 {
	return int64(d) / int64(time.Hour)
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) GoString() string {
	return time.Duration(d).String()
}

// Abs Returns the absolute value of duration
func (d Duration) Abs() Duration {
	return Duration(time.Duration(d).Abs())
}

// Truncate converts duration to a multiple of m by truncating any additional duration. Absolute value of m is used for conversions
func (d Duration) Truncate(m Duration) Duration {
	return Duration(time.Duration(d).Truncate(time.Duration(m).Abs()))
}

// Round d to a multiple of m. Round up in case of halfway values. Absolute value of m is used for conversions
func (d Duration) Round(m Duration) Duration {
	return Duration(time.Duration(d).Round(time.Duration(m).Abs()))
}
