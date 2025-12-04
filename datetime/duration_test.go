package datetime_test

import (
	"fmt"
	"testing"

	"github.com/ram-nad/go-utils/datetime"
)

func TestDuration(t *testing.T) {
	t.Run("Nanoseconds", func(t *testing.T) {
		v := datetime.Nanoseconds(123)

		if int64(v) != 123 {
			t.Errorf("Expected int64 of duration to be 123, got %d", int64(v))
		}

		if v.String() != "123ns" {
			t.Errorf(
				"Expected string representation of duration to be \"123ns\", got %q",
				v.String(),
			)
		}
	})

	t.Run("Microseconds", func(t *testing.T) {
		v := datetime.Microseconds(-113)

		if int64(v) != -113*1000 {
			t.Errorf("Expected int64 of duration to be -113000, got %d", int64(v))
		}

		if v.String() != "-113µs" {
			t.Errorf(
				"Expected string representation of duration to be \"-113µs\", got %q",
				v.String(),
			)
		}
	})

	t.Run("Milliseconds", func(t *testing.T) {
		v := datetime.Milliseconds(453)

		if int64(v) != 453*1000*1000 {
			t.Errorf("Expected int64 of duration to be 453000000, got %d", int64(v))
		}

		if v.String() != "453ms" {
			t.Errorf(
				"Expected string representation of duration to be \"453ms\", got %q",
				v.String(),
			)
		}
	})

	t.Run("Seconds", func(t *testing.T) {
		v := datetime.Seconds(-12)

		if int64(v) != -12*1000*1000*1000 {
			t.Errorf("Expected int64 of duration to be -12000000000, got %d", int64(v))
		}

		if v.String() != "-12s" {
			t.Errorf(
				"Expected string representation of duration to be \"-12s\", got %q",
				v.String(),
			)
		}
	})

	t.Run("Minutes", func(t *testing.T) {
		v := datetime.Minutes(2)

		if int64(v) != 2*60*1000*1000*1000 {
			t.Errorf("Expected int64 of duration to be 120000000000, got %d", int64(v))
		}

		if v.String() != "2m0s" {
			t.Errorf(
				"Expected string representation of duration to be \"2m0s\", got %q",
				v.String(),
			)
		}
	})

	t.Run("Hours", func(t *testing.T) {
		v := datetime.Hours(1)

		if int64(v) != 60*60*1000*1000*1000 {
			t.Errorf("Expected int64 of duration to be 3600000000000, got %d", int64(v))
		}

		if v.String() != "1h0m0s" {
			t.Errorf(
				"Expected string representation of duration to be \"1h0m0s\", got %q",
				v.String(),
			)
		}
	})
}

func TestDurationMethods(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		x := datetime.Hours(
			1,
		) + datetime.Minutes(
			3,
		) + datetime.Seconds(
			40,
		) + datetime.Milliseconds(
			50,
		) + datetime.Microseconds(
			121,
		) + datetime.Nanoseconds(
			5,
		)

		if x.String() != "1h3m40.050121005s" {
			t.Errorf("Expected String 1h3m40.050121002s, got %q", x.String())
		}

		x = datetime.Hours(
			-1,
		) + datetime.Minutes(
			-3,
		) + datetime.Seconds(
			-10,
		) + datetime.Milliseconds(
			-30,
		) + datetime.Microseconds(
			-121,
		) + datetime.Nanoseconds(
			-2,
		)

		if x.String() != "-1h3m10.030121002s" {
			t.Errorf("Expected String -1h3m10.030021002s, got %q", x.String())
		}
	})

	t.Run("GoString", func(t *testing.T) {
		x := datetime.Hours(
			1,
		) + datetime.Minutes(
			3,
		) + datetime.Seconds(
			40,
		) + datetime.Milliseconds(
			50,
		) + datetime.Microseconds(
			121,
		) + datetime.Nanoseconds(
			2,
		)

		if x.GoString() != "1h3m40.050121002s" {
			t.Errorf("Expected String 1h3m40.050121002s, got %q", x.String())
		}

		x = datetime.Hours(
			-1,
		) + datetime.Minutes(
			-3,
		) + datetime.Seconds(
			-10,
		) + datetime.Milliseconds(
			-30,
		) + datetime.Microseconds(
			-21,
		) + datetime.Nanoseconds(
			-2,
		)

		if x.GoString() != "-1h3m10.030021002s" {
			t.Errorf("Expected String -1h3m10.030021002s, got %q", x.String())
		}
	})

	t.Run("Abs", func(t *testing.T) {
		x := datetime.Hours(
			-1,
		) + datetime.Minutes(
			-3,
		) + datetime.Seconds(
			-40,
		) + datetime.Milliseconds(
			-50,
		) + datetime.Microseconds(
			-121,
		) + datetime.Nanoseconds(
			-2,
		)
		x = x.Abs()

		if x != datetime.Hours(
			1,
		)+datetime.Minutes(
			3,
		)+datetime.Seconds(
			40,
		)+datetime.Milliseconds(
			50,
		)+datetime.Microseconds(
			121,
		)+datetime.Nanoseconds(
			2,
		) {
			t.Errorf("Expected positive duration, got %v", x)
		}
	})

	t.Run("Truncate", func(t *testing.T) {
		x := datetime.Hours(
			2,
		) + datetime.Minutes(
			2,
		) + datetime.Seconds(
			13,
		) + datetime.Milliseconds(
			140,
		) + datetime.Microseconds(
			121,
		) + datetime.Nanoseconds(
			102,
		)
		x = x.Truncate(datetime.Milliseconds(100))

		if x != datetime.Hours(
			2,
		)+datetime.Minutes(
			2,
		)+datetime.Seconds(
			13,
		)+datetime.Milliseconds(
			100,
		) {
			t.Errorf("Expected truncated duration, got %v", x)
		}

		y := datetime.Hours(
			-2,
		) + datetime.Minutes(
			-2,
		) + datetime.Seconds(
			-13,
		) + datetime.Milliseconds(
			-130,
		) + datetime.Microseconds(
			-120,
		) + datetime.Nanoseconds(
			-102,
		)
		y = y.Truncate(datetime.Seconds(4))

		if y != datetime.Hours(-2)+datetime.Minutes(-2)+datetime.Seconds(-12) {
			t.Errorf("Expected truncated duration, got %v", y)
		}

		z := datetime.Hours(2) + datetime.Minutes(2) + datetime.Seconds(32)
		z = z.Truncate(datetime.Seconds(-100))

		if z != datetime.Hours(2)+datetime.Minutes(1)+datetime.Seconds(40) {
			t.Errorf("Expected truncated duration, got %v", z)
		}
	})

	t.Run("Round", func(t *testing.T) {
		x := datetime.Hours(
			2,
		) + datetime.Minutes(
			2,
		) + datetime.Seconds(
			13,
		) + datetime.Milliseconds(
			140,
		) + datetime.Microseconds(
			127,
		) + datetime.Nanoseconds(
			102,
		)
		x = x.Round(datetime.Milliseconds(50))

		if x != datetime.Hours(
			2,
		)+datetime.Minutes(
			2,
		)+datetime.Seconds(
			13,
		)+datetime.Milliseconds(
			150,
		) {
			t.Errorf("Expected truncated duration, got %v", x)
		}

		y := datetime.Hours(2) + datetime.Minutes(2) + datetime.Seconds(30)
		y = y.Round(datetime.Seconds(-100))

		if y != datetime.Hours(2)+datetime.Minutes(3)+datetime.Seconds(20) {
			t.Errorf("Expected truncated duration, got %v", y)
		}
	})
}

func TestDurationAccessMethods(t *testing.T) {
	x := datetime.Hours(
		1,
	) + datetime.Minutes(
		3,
	) + datetime.Seconds(
		40,
	) + datetime.Milliseconds(
		50,
	) + datetime.Microseconds(
		121,
	) + datetime.Nanoseconds(
		2,
	)

	t.Run("Nanoseconds", func(t *testing.T) {
		ns := x.Nanoseconds()

		if ns != (60*60*1000*1000*1000)+(3*60*1000*1000*1000)+(40*1000*1000*1000)+(50*1000*1000)+(121*1000)+2 {
			t.Errorf("Didn't get expected value of ns")
		}
	})

	t.Run("Microseconds", func(t *testing.T) {
		us := x.Microseconds()

		if us != (60*60*1000*1000)+(3*60*1000*1000)+(40*1000*1000)+(50*1000)+121 {
			t.Errorf("Didn't get expected value of µs")
		}
	})

	t.Run("Milliseconds", func(t *testing.T) {
		ms := x.Milliseconds()

		if ms != (60*60*1000)+(3*60*1000)+(40*1000)+50 {
			t.Errorf("Didn't get expected value of ms")
		}
	})

	t.Run("Seconds", func(t *testing.T) {
		s := x.Seconds()

		if s != (60*60)+(3*60)+40 {
			t.Errorf("Didn't get expected value of s")
		}
	})

	t.Run("Minutes", func(t *testing.T) {
		m := x.Minutes()

		if m != (60)+3 {
			t.Errorf("Didn't get expected value of m")
		}
	})

	t.Run("Hours", func(t *testing.T) {
		h := x.Hours()

		if h != 1 {
			t.Errorf("Didn't get expected value of h")
		}
	})
}

func TestParseDuration(t *testing.T) {
	t.Run("ParseDuration", func(t *testing.T) {
		x, err := datetime.ParseDuration("2h2m13.140121102s")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if x != datetime.Hours(
			2,
		)+datetime.Minutes(
			2,
		)+datetime.Seconds(
			13,
		)+datetime.Milliseconds(
			140,
		)+datetime.Microseconds(
			121,
		)+datetime.Nanoseconds(
			102,
		) {
			t.Errorf("Didn't get correct parsed duration, got %v", x)
		}

		y, err := datetime.ParseDuration("2h2m13s32ms12us1ns")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if y != datetime.Hours(
			2,
		)+datetime.Minutes(
			2,
		)+datetime.Seconds(
			13,
		)+datetime.Milliseconds(
			32,
		)+datetime.Microseconds(
			12,
		)+datetime.Nanoseconds(
			1,
		) {
			t.Errorf("Didn't get correct parsed duration, got %v", y)
		}

		z, err := datetime.ParseDuration("2h2m64s")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if z != datetime.Hours(2)+datetime.Minutes(3)+datetime.Seconds(4) {
			t.Errorf("Didn't get correct parsed duration, got %v", z)
		}
	})
}

func ExampleDuration() {
	x := datetime.Hours(
		2,
	) + datetime.Minutes(
		2,
	) + datetime.Seconds(
		13,
	) + datetime.Milliseconds(
		140,
	) + datetime.Microseconds(
		121,
	) + datetime.Nanoseconds(
		2,
	)
	_, _ = fmt.Println(x.String())
	// Output: 2h2m13.140121002s
}

func ExampleParseDuration() {
	x, err := datetime.ParseDuration("2h2m13s")
	if err == nil {
		_, _ = fmt.Printf(
			"Hour: %d, Minute: %d, Second: %d",
			x.Hours(),
			x.Minutes(),
			x.Seconds(),
		)
	}
	// Output: Hour: 2, Minute: 122, Second: 7333
}

func ExampleParseDuration_second() {
	x, err := datetime.ParseDuration("2h2m13.10235s")
	if err == nil {
		_, _ = fmt.Printf(
			"Hour: %d, Minute: %d, Second: %d, Milliseconds: %d, Microseconds: %d, Nanoseconds: %d",
			x.Hours(),
			x.Minutes(),
			x.Seconds(),
			x.Milliseconds(),
			x.Microseconds(),
			x.Nanoseconds(),
		)
	}
	// Output: Hour: 2, Minute: 122, Second: 7333, Milliseconds: 7333102, Microseconds: 7333102350, Nanoseconds: 7333102350000
}

func ExampleDuration_Abs() {
	x := datetime.Hours(
		-1,
	) + datetime.Minutes(
		-3,
	) + datetime.Seconds(
		-40,
	) + datetime.Milliseconds(
		-50,
	) + datetime.Microseconds(
		-121,
	) + datetime.Nanoseconds(
		-2,
	)
	x = x.Abs()
	_, _ = fmt.Println(x.String())
	// Output: 1h3m40.050121002s
}

func ExampleDuration_Truncate() {
	x := datetime.Hours(
		2,
	) + datetime.Minutes(
		2,
	) + datetime.Seconds(
		13,
	) + datetime.Milliseconds(
		140,
	) + datetime.Microseconds(
		121,
	) + datetime.Nanoseconds(
		102,
	)
	x = x.Truncate(datetime.Milliseconds(100))
	_, _ = fmt.Println(x.String())
	// Output: 2h2m13.1s
}

func ExampleDuration_Round() {
	x := datetime.Hours(2) + datetime.Minutes(2) + datetime.Seconds(33)
	x = x.Round(datetime.Seconds(100))
	_, _ = fmt.Println(x.String())
	// Output: 2h3m20s
}
