// Package datetime makes it easy to work with time and clocks
package datetime

/*
Adding a separate datetime package to make sure all time operations are in UTC.

Additionally avoid any confusion with time.Time operations by separately exposing Wall and Monotonic Clocks.
*/

import (
	"fmt"
	"time"
)

// Time stores the date and time in UTC
type Time time.Time

// These are aliases to time package types
type (
	Month   = time.Month
	Weekday = time.Weekday
)

// PrintFormat is used to specify format when printing time
type PrintFormat string

const (
	RFC822      PrintFormat = time.RFC822
	RFC822Z     PrintFormat = time.RFC822Z
	RFC1123     PrintFormat = time.RFC1123
	RFC1123Z    PrintFormat = time.RFC1123Z
	RFC3339     PrintFormat = time.RFC3339
	RFC3339Nano PrintFormat = time.RFC3339Nano
	DateOnly    PrintFormat = time.DateOnly
	TimeOnly    PrintFormat = time.TimeOnly
)

func Date(year, month, day, hour, minute, second, nsec int) Time {
	// This will always return without monotonic clock
	return Time(
		time.Date(year, Month(month), day, hour, minute, second, nsec, time.UTC),
	)
}

func Now() Time {
	now := time.Now()
	// Make sure monotonic clock is stripped
	return Time(time.Unix(now.Unix(), int64(now.Nanosecond())).UTC())
}

func Unix(sec, nsec int64) Time {
	// This will always return without monotonic clock
	// Additionally calling UTC strips monotonic clock from the time
	return Time(time.Unix(sec, nsec).UTC())
}

func UnixMilli(msec int64) Time {
	// This will always return without monotonic clock
	return Time(time.UnixMilli(msec).UTC())
}

func Since(t Time) Duration {
	return Duration(time.Since(time.Time(t)))
}

func Until(t Time) Duration {
	return Duration(time.Until(time.Time(t)))
}

func (t Time) AddDate(years int, months int, days int) Time {
	return Time(time.Time(t).AddDate(years, months, days))
}

func (t Time) Add(d Duration) Time {
	return Time(time.Time(t).Add(time.Duration(d)))
}

func (t Time) Sub(u Time) Duration {
	return Duration(time.Time(t).Sub(time.Time(u)))
}

func (t Time) Equal(o Time) bool {
	return time.Time(t).Equal(time.Time(o))
}

func (t Time) Before(o Time) bool {
	return time.Time(t).Before(time.Time(o))
}

func (t Time) After(o Time) bool {
	return time.Time(t).After(time.Time(o))
}

// Truncate returns the result of rounding t down to a multiple of d (towards the zero time).
func (t Time) Truncate(d Duration) Time {
	return Time(time.Time(t).Truncate(time.Duration(d).Abs()))
}

// Round returns the result of rounding t to the nearest multiple of d. The rounding behavior for halfway values is to round up.
func (t Time) Round(d Duration) Time {
	return Time(time.Time(t).Round(time.Duration(d).Abs()))
}

// Year component of the time
func (t Time) Year() int {
	return time.Time(t).Year()
}

// Month component of the time
func (t Time) Month() Month {
	return time.Time(t).Month()
}

// Day component of the time
func (t Time) Day() int {
	return time.Time(t).Day()
}

// Weekday returns weekday of the time
func (t Time) Weekday() Weekday {
	return time.Time(t).Weekday()
}

// YearDay returns day of the year for the time
func (t Time) YearDay() int {
	return time.Time(t).YearDay()
}

// Hour component of the time
func (t Time) Hour() int {
	return time.Time(t).Hour()
}

// Minute component of the time
func (t Time) Minute() int {
	return time.Time(t).Minute()
}

// Second component of the time
func (t Time) Second() int {
	return time.Time(t).Second()
}

// Nanosecond component of the time
func (t Time) Nanosecond() int {
	return time.Time(t).Nanosecond()
}

// Unix returns the time as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.
func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

/*
Format is used to print time in a given format

Example:

	t := Date(2021, 1, 1, 0, 0, 0, 0)
	fmt.Println(t.Format(datetime.RFC3339))
	Output: 2021-01-01T00:00:00Z

	t := Date(2021, 1, 1, 0, 0, 1, 23)
	fmt.Println(t.Format(datetime.RFC3339Nano))
	Output: 2021-01-01T00:00:01.000000023Z

	t := Date(2021, 1, 1, 0, 0, 0, 0)
	fmt.Println(t.Format(datetime.DateOnly))
	Output: 2021-01-01

	t := Date(2021, 1, 1, 23, 11, 0, 0)
	fmt.Println(t.Format(datetime.TimeOnly))
	Output: 23:11:00
*/
func (t Time) Format(format PrintFormat) string {
	return time.Time(t).Format(string(format))
}

// ISOString returns the time in ISO 8601 format
func (t Time) ISOString() string {
	return t.Format(RFC3339)
}

// ISOStringNano returns the time in ISO 8601 format with nanoseconds
func (t Time) ISOStringNano() string {
	return t.Format(RFC3339Nano)
}

// String returns the human readable string representation of the time
func (t Time) String() string {
	return t.Format(RFC1123)
}

// GoString returns the Go's representation of the time
func (t Time) GoString() string {
	return fmt.Sprintf("{s: %d, nsec: %d}", t.Unix(), t.Nanosecond())
}
