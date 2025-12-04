package stringhash_test

/*
ASCII tests are taken from https://github.com/darkskyapp/string-hash

Tests using non-ASCII code points are expected to generate
same hash as they would using the JS implementation.
*/

import (
	"fmt"
	"testing"

	"github.com/ram-nad/go-utils/stringhash"
)

type hashTestCase struct {
	str  string
	hash uint32
}

func runTest(t *testing.T, tests []hashTestCase) {
	for _, table := range tests {
		calculatedHash := stringhash.Hash(table.str)
		if calculatedHash != table.hash {
			t.Logf("calculated hash of %q is %d", table.str, calculatedHash)
			t.Errorf("expected hash %d. got %d", table.hash, calculatedHash)
		}
	}
}

func TestHash(t *testing.T) {
	tests := []hashTestCase{
		{"Mary had a little lamb.", 1766333550},
		{"Hello, world!", 343662184},

		{"Hello, 世界", 1861035601},
		{"A界𐐷", 1362180894},

		{"\x08\xc3\x01", 193382351},
		{"\xed\x80\x01", 193380297},
		{"\xd8\x01\xc3", 193308639},
		{"\xdc\x11\x02", 193379722},
	}

	t.Run("ASCII", func(t *testing.T) {
		runTest(t, tests[:2])
	})

	t.Run("Non-ASCII", func(t *testing.T) {
		runTest(t, tests[2:4])
	})

	t.Run("Invalid Unicode", func(t *testing.T) {
		runTest(t, tests[4:8])
	})

	t.Run("Large/String", func(t *testing.T) {
		largeStr := "Hello This is a very large string, supposed to be more than 256 characters: abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ... Wait.. I can write a song: Hey Jude, don't make it bad. Take a sad song and make it better. Remember to let her into your heart, Then you can start to make it better. Hey Jude, don't be afraid, You were made to go out and get her, The minute you let her under your skin Then you begin to make it better. Na Na Na Na Na Na Na Na Na Na Na.. Hey Jude"
		runTest(t, []hashTestCase{{largeStr, 1610438000}})
	})
}

func ExampleHash() {
	hash := stringhash.Hash("Hello")
	_, _ = fmt.Println(hash)
	// Output: 181379975
}

func BenchmarkHashASCII(b *testing.B) {
	for b.Loop() {
		stringhash.Hash("Hello, World, from inside a benchmark")
	}
}

func BenchmarkHashNonASCII(b *testing.B) {
	for b.Loop() {
		stringhash.Hash("कल करे सो आज कर, आज करे सो अब")
	}
}
