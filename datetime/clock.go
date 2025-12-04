package datetime

/*
The implementation of ClockTime is based on the time.Time type from the time package.

It uses unsafe.Pointer to convert between ClockTime and time.Time.

If internal representation of time.Time changes, the implementation of ClockTime will need to be updated.

Example:
	clockStart := datetime.NowClock()
	.... do something ....
	duration := datetime.SinceClock(clockStart)
	fmt.Prinf("Time taken: %v", duration)
*/

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

// ClockTime is used for measuring time
type ClockTime int64

func subClockTime(t1, t2 ClockTime) Duration {
	r := t1 - t2

	if r < 0 && t1 > t2 {
		return Duration(math.MaxInt64) // Overflow
	}

	if r > 0 && t1 < t2 {
		return Duration(math.MinInt64) // Underflow
	}

	return Duration(r)
}

func addClockTime(t1 ClockTime, d Duration) ClockTime {
	r := t1 + ClockTime(d)

	if r < t1 && d > 0 {
		return ClockTime(math.MaxInt64) // Overflow
	}

	if r > t1 && d < 0 {
		return ClockTime(math.MinInt64) // Underflow
	}

	return r
}

func NowClock() ClockTime {
	now := time.Now()
	//nolint:gosec // this depends on internals of time.Time to get access to monotonic clock
	mono := *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&now)) + unsafe.Sizeof(uint64(0))))

	return ClockTime(mono)
}

func SinceClock(t ClockTime) Duration {
	now := NowClock()

	return subClockTime(now, t)
}

func UntilClock(t ClockTime) Duration {
	now := NowClock()

	return subClockTime(t, now)
}

func (t ClockTime) Sub(d ClockTime) Duration {
	return subClockTime(t, d)
}

func (t ClockTime) Add(d Duration) ClockTime {
	return addClockTime(t, d)
}

// String returns Human readable string representation of monotonic clock time
func (t ClockTime) String() string {
	s := int64(t) / int64(time.Second)
	ns := int64(t) % int64(time.Second)
	if ns < 0 {
		ns = -ns
	}
	return fmt.Sprintf("m%+d.%09d", s, ns)
}

// GoString returns Go's representation of the monotonic clock time
func (t ClockTime) GoString() string {
	return fmt.Sprintf(
		"{s: %+d, nsec: %+d}",
		t/ClockTime(time.Second),
		t%ClockTime(time.Second),
	)
}
