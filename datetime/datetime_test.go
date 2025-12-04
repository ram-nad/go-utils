package datetime_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ram-nad/go-utils/datetime"
)

func TestDate(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 3, 0, 0)
	if dt.ISOString() != "2021-01-01T00:03:00Z" {
		t.Errorf("Expected \"2021-01-01T00:03:00Z\", but got %q", dt.ISOString())
	}
}

func TestNow(t *testing.T) {
	now := datetime.Now()
	if !strings.HasSuffix(now.ISOString(), "Z") {
		t.Errorf("Expected time to end with Z, but got %q", now.ISOString())
	}
}

func TestEqual(t *testing.T) {
	t1 := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	t2 := datetime.Date(2021, 1, 1, 0, 0, 0, 0)

	if !t1.Equal(t2) {
		t.Errorf("Expected %v to be equal to %v", t1, t2)
	}
}

func TestBefore(t *testing.T) {
	t1 := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	t2 := datetime.Date(2021, 1, 2, 0, 0, 0, 0)

	if !t1.Before(t2) {
		t.Errorf("Expected %v to be before %v", t1, t2)
	}
}

func TestAfter(t *testing.T) {
	t1 := datetime.Date(2021, 2, 1, 0, 0, 0, 0)
	t2 := datetime.Date(2021, 1, 4, 0, 0, 0, 0)

	if !t1.After(t2) {
		t.Errorf("Expected %v to be after %v", t1, t2)
	}
}

func TestUnix(t *testing.T) {
	expected := datetime.Unix(1633072800, 0)
	actual := datetime.Date(2021, 10, 1, 7, 20, 0, 0)

	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestUnixMilli(t *testing.T) {
	expected := datetime.UnixMilli(1633072800020)
	actual := datetime.Date(2021, 10, 1, 7, 20, 0, 20000000)

	if !expected.Equal(actual) {
		t.Errorf(
			"Expected %s, but got %s",
			expected.Format(datetime.RFC3339Nano),
			expected.Format(datetime.RFC3339Nano),
		)
	}
}

func TestAddDate(t *testing.T) {
	initial := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	expected := datetime.Date(2022, 2, 2, 0, 0, 0, 0)

	actual := initial.AddDate(1, 1, 1)
	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestAdd(t *testing.T) {
	initial := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	expected := datetime.Date(2021, 1, 1, 1, 0, 0, 0)

	actual := initial.Add(datetime.Hours(1))
	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSub(t *testing.T) {
	t1 := datetime.Date(2021, 1, 1, 1, 0, 0, 0)
	t2 := datetime.Date(2021, 1, 1, 0, 0, 0, 0)

	expected := datetime.Duration(time.Hour)
	actual := t1.Sub(t2)

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSince(t *testing.T) {
	dt := datetime.Now()
	s := datetime.Since(dt)

	if s < 0 {
		t.Errorf("Expected positive duration, but got %v", s)
	}
}

func TestUntil(t *testing.T) {
	dt := datetime.Now()
	u := datetime.Until(dt)

	if u > 0 {
		t.Errorf("Expected negative duration, but got %v", u)
	}
}

func TestTruncate(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 6, 1, 1, 1)
	expected := datetime.Date(2021, 1, 1, 4, 0, 0, 0)

	actual := dt.Truncate(datetime.Hours(4))
	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestRound(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 6, 1, 1, 1)
	expected := datetime.Date(2021, 1, 1, 8, 0, 0, 0)

	actual := dt.Round(datetime.Hours(-4))
	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestFormat(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 0)

	expected := "2021-01-01"
	actual := dt.Format(datetime.DateOnly)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}

	expected = "00:00:00"
	actual = dt.Format(datetime.TimeOnly)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestISOString(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	expected := "2021-01-01T00:00:00Z"
	actual := dt.ISOString()
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestISOStringNano(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 120)
	expected := "2021-01-01T00:00:00.00000012Z"
	actual := dt.ISOStringNano()
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestString(t *testing.T) {
	dt := datetime.Date(2021, 1, 2, 0, 0, 0, 0)
	expected := "Sat, 02 Jan 2021 00:00:00 UTC"
	actual := dt.String()
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestGoString(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	expected := "{s: 1609459200, nsec: 0}"
	actual := dt.GoString()
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestComponents(t *testing.T) {
	dt := datetime.Date(2021, 1, 5, 12, 3, 45, 1467)

	if dt.Year() != 2021 {
		t.Errorf("Expected Year 2021, but got %d", dt.Year())
	}

	if dt.Month() != 1 {
		t.Errorf("Expected Month 1, but got %d", dt.Month())
	}

	if dt.Day() != 5 {
		t.Errorf("Expected Day 5, but got %d", dt.Day())
	}

	if dt.Weekday() != time.Tuesday {
		t.Errorf("Expected Weekday Tuesday, but got %v", dt.Weekday())
	}

	if dt.YearDay() != 5 {
		t.Errorf("Expected YearDay 5, but got %d", dt.YearDay())
	}

	if dt.Hour() != 12 {
		t.Errorf("Expected Hour 12, but got %d", dt.Hour())
	}

	if dt.Minute() != 3 {
		t.Errorf("Expected Minute 3, but got %d", dt.Minute())
	}

	if dt.Second() != 45 {
		t.Errorf("Expected Second 45, but got %d", dt.Second())
	}

	if dt.Nanosecond() != 1467 {
		t.Errorf("Expected Nanosecond 1467, but got %d", dt.Nanosecond())
	}
}

func TestGetUnixValue(t *testing.T) {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	expected := int64(1609459200)
	actual := dt.Unix()

	if expected != actual {
		t.Errorf("Expected Unix %d, but got %d", expected, actual)
	}
}

func ExampleDate() {
	dt := datetime.Date(2021, 1, 1, 0, 0, 0, 0)
	_, _ = fmt.Println(dt.ISOString())
	// Output: 2021-01-01T00:00:00Z
}

func ExampleDate_second() {
	dt := datetime.Date(2021, 1, 5, 11, 3, 10, 0o1)
	_, _ = fmt.Printf("Year: %d ", dt.Year())
	_, _ = fmt.Printf("Month: %d (%s) ", dt.Month(), dt.Month().String())
	_, _ = fmt.Printf("Day: %d ", dt.Day())
	_, _ = fmt.Printf("Weekday: %d ", dt.Weekday())
	_, _ = fmt.Printf("YearDay: %d ", dt.YearDay())
	_, _ = fmt.Printf("Hour: %d ", dt.Hour())
	_, _ = fmt.Printf("Minute: %d ", dt.Minute())
	_, _ = fmt.Printf("Second: %d ", dt.Second())
	_, _ = fmt.Printf("Nanosecond: %d ", dt.Nanosecond())
	// Output: Year: 2021 Month: 1 (January) Day: 5 Weekday: 2 YearDay: 5 Hour: 11 Minute: 3 Second: 10 Nanosecond: 1
}
