package datetime_test

import (
	"math"
	"testing"

	"github.com/ram-nad/go-utils/datetime"
)

func TestNowClock(t *testing.T) {
	clockNow := datetime.NowClock()
	clockNext := datetime.NowClock()

	if clockNext < clockNow {
		t.Errorf("Expected clockNext to be greater than clockNow")
	}
}

func TestSubClockTime(t *testing.T) {
	t1 := datetime.ClockTime(1000)
	t2 := datetime.ClockTime(500)
	expected := datetime.Duration(500)

	if got := t1.Sub(t2); got != expected {
		t.Errorf("%v.Sub(%v) = %v, want %v", t1, t2, got, expected)
	}

	// Test overflow
	t1 = datetime.ClockTime(math.MaxInt64 - 678)
	t2 = datetime.ClockTime(-700)
	expected = datetime.Duration(math.MaxInt64)

	if got := t1.Sub(t2); got != expected {
		t.Errorf("%v.Sub(%v) = %v, want %v", t1, t2, got, expected)
	}

	// Test underflow
	t1 = datetime.ClockTime(math.MinInt64 + 200)
	t2 = datetime.ClockTime(205)
	expected = datetime.Duration(math.MinInt64)

	if got := t1.Sub(t2); got != expected {
		t.Errorf("%v.Sub(%v) = %v, want %v", t1, t2, got, expected)
	}
}

func TestAddClockTime(t *testing.T) {
	t1 := datetime.ClockTime(1000)
	d := datetime.Duration(500)
	expected := datetime.ClockTime(1500)

	if got := t1.Add(d); got != expected {
		t.Errorf("%v.Add(%v) = %v, want %v", t1, d, got, expected)
	}

	// Test overflow
	t1 = datetime.ClockTime(math.MaxInt64)
	d = datetime.Duration(1)
	expected = datetime.ClockTime(math.MaxInt64)

	if got := t1.Add(d); got != expected {
		t.Errorf("%v.Add(%v) = %v, want %v", t1, d, got, expected)
	}

	// Test underflow
	t1 = datetime.ClockTime(math.MinInt64)
	d = datetime.Duration(-1)
	expected = datetime.ClockTime(math.MinInt64)

	if got := t1.Add(d); got != expected {
		t.Errorf("%v.Add(%v) = %v, want %v", t1, d, got, expected)
	}
}

func TestClockString(t *testing.T) {
	tc := datetime.ClockTime(24242424191000)
	expected := "m+24242.424191000"

	if got := tc.String(); got != expected {
		t.Errorf("%v.String() = %v, want %v", tc, got, expected)
	}

	tc = datetime.ClockTime(-15242444181000)
	expected = "m-15242.444181000"

	if got := tc.String(); got != expected {
		t.Errorf("%v.String() = %v, want %v", tc, got, expected)
	}
}

func TestClockGoString(t *testing.T) {
	tc := datetime.ClockTime(-1545434181023)
	expected := "{s: -1545, nsec: -434181023}"

	if got := tc.GoString(); got != expected {
		t.Errorf("%v.String() = %v, want %v", tc, got, expected)
	}
}

func TestSinceClock(t *testing.T) {
	start := datetime.NowClock()
	duration := datetime.SinceClock(start)

	if duration < 0 {
		t.Errorf("SinceClock(%v) = %v, want >= 0", start, duration)
	}
}

func TestUntilClock(t *testing.T) {
	start := datetime.NowClock()
	duration := datetime.UntilClock(start)

	if duration > 0 {
		t.Errorf("UntilClock(%v) = %v, want <= 0", start, duration)
	}
}
