/*
Package stringhash computes fast hash of the given string.

See https://www.npmjs.com/package/string-hash.
*/
package stringhash

/*
Base Implementation: https://github.com/darkskyapp/string-hash

JavaScript uses UTF-16 to encode strings, while Go uses UTF-8. We need to
convert Go Strings to UTF-16 and then compute the hash. Also, JS does
bitwise operations on signed 32 bit integer. But for the computations we are doing
it will give the same result if we use unsigned 32 bit integer.

Rigorously tested to ensure 100% compatibility with the JavaScript implementation.

References:
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charCodeAt
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/length

https://en.wikipedia.org/wiki/UTF-16
*/

import (
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

const stackAllocateSizeLimit = 1024 / 4 // 1 KB, 4 bytes per uint32

// Hash returns hash of the string
func Hash(str string) uint32 {
	var hash uint32 = 5381

	/*
		Strings would be usually small so, it makes sense to
		convert all of them at once before computation. This
		can be changed if over usage of memory is a concern.
	*/

	lenStr := uint(len(str))
	//nolint:gosec // this is a safe conversion, as we are not modifying the bytes
	strDataPtr := unsafe.StringData(str)

	// Avoid costly string to []byte conversion
	//nolint:gosec // this is a safe conversion, as we are not modifying the bytes
	strBytes := unsafe.Slice(strDataPtr, lenStr)

	// Temporary array to store UTF-16 code points
	// Will heap allocate if this doesn't suffice
	utf16arr := [stackAllocateSizeLimit]uint32{}
	var utf16Slice []uint32

	if lenStr <= stackAllocateSizeLimit {
		// Use stack allocated array
		utf16Slice = utf16arr[:]
	} else {
		// Heap allocate
		utf16Slice = make([]uint32, lenStr)
	}

	// Index for utf16Slice
	var j uint = 0

	_ = strBytes[lenStr-1]   // Bounds check hint to the compiler
	_ = utf16Slice[lenStr-1] // Bounds check hint to the compiler

	// Convert string to UTF-16
	for i := uint(0); i < lenStr && j <= i; {
		utf8Rune, runeWidth := utf8.DecodeRune(
			strBytes[i:],
		) // `runeWidth` cannot be 0, as we don't pass empty slice to DecodeRune
		utf16CodePointsCount := utf16.RuneLen(
			utf8Rune,
		) // `utf8Rune` is always a valid rune or utf8.RuneError

		switch utf16CodePointsCount {
		case 1:
			if utf8Rune == utf8.RuneError {
				utf16Slice[j] = uint32(
					strBytes[i],
				) // If we have an invalid UTF-8 sequence, we just store the byte as a code point
			} else {
				utf16Slice[j] = uint32(utf8Rune)
			}
			j += 1
		//nolint:mnd // Case when we have a surrogate pair
		case 2:
			r1, r2 := utf16.EncodeRune(utf8Rune)
			utf16Slice[j] = uint32(r1)
			utf16Slice[j+1] = uint32(r2)
			j += 2
		}
		i += uint(runeWidth)
	}

	// Resize the slice to the actual length
	utf16CodePoints := utf16Slice[:j]

	for i := len(utf16CodePoints) - 1; i >= 0; i -= 1 {
		//nolint:mnd // Magic number is used in the algorithm
		hash = (hash * 33) ^ uint32(utf16CodePoints[i])
	}

	return hash
}
