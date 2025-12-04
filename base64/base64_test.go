package base64_test

import (
	"fmt"
	"strings"
	"testing"

	base64Encoding "github.com/ram-nad/go-utils/base64"
)

type encoding string

type testCase struct {
	str        string
	encodedStr string
}

const (
	base64    encoding = "base64"
	base64url encoding = "base64url"
)

//nolint:gochecknoglobals // Cannot make these constants, but they are not modified
var globalTests []testCase = []testCase{
	{
		"\x00\x01\x02\x03\x04\x05\x06\x07\b\t\n\x0B\f\r\x0E\x0F\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\x7F",
		"AAECAwQFBgcICQoLDA0ODxAREhMUFRYXGBkaGxwdHh8gISIjJCUmJygpKissLS4vMDEyMzQ1Njc4OTo7PD0+P0BBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWltcXV5fYGFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6e3x9fn8=",
	},
	{
		"",
		"",
	},
	{
		"a",
		"YQ==",
	},
	{
		"aa",
		"YWE=",
	},
	{
		"aaa",
		"YWFh",
	},
	{
		"foo\x00",
		"Zm9vAA==",
	},
	{
		"foo\x00\x00",
		"Zm9vAAA=",
	},
	{
		"f",
		"Zg==",
	},
	{
		"fo",
		"Zm8=",
	},
	{
		"foo",
		"Zm9v",
	},
	{
		"foob",
		"Zm9vYg==",
	},
	{
		"fooba",
		"Zm9vYmE=",
	},
	{
		"foobar",
		"Zm9vYmFy",
	},
	{
		"\xFF\xFF\xC0",
		"///A",
	},
	{
		"C\xFF\xFF",
		"Q///",
	},
	{
		"C\xEF\xBE",
		"Q+++",
	},
	{
		"\x00",
		"AA==",
	},
	{
		"\x00a",
		"AGE=",
	},
	{
		"undefined",
		"dW5kZWZpbmVk",
	},
	{
		"null",
		"bnVsbA==",
	},
	{
		"0",
		"MA==",
	},
	{
		"7",
		"Nw==",
	},
	{
		"1.5",
		"MS41",
	},
	{
		"true",
		"dHJ1ZQ==",
	},
	{
		"false",
		"ZmFsc2U=",
	},
	{
		"NaN",
		"TmFO",
	},
	{
		"-Infinity",
		"LUluZmluaXR5",
	},
	{
		"Infinity",
		"SW5maW5pdHk=",
	},
}

//nolint:gochecknoglobals // Cannot make these constants, but they are not modified
var globalTestsWithInvalidCharacters []testCase = []testCase{
	{
		"",
		"A$Y^",
	},
	{
		"",
		"\n{\r}*=",
	},
	{
		"",
		"A~==",
	},
	{
		"",
		"A*%=",
	},
}

func runEncodeTest(
	t *testing.T,
	enc encoding,
	str string,
	expectedEncodedResult string,
) {
	var encodingResult string
	switch enc {
	case base64:
		encodingResult = base64Encoding.EncodeBase64(str)
	case base64url:
		encodingResult = base64Encoding.EncodeBase64URL(str)
	default:
		t.Fatalf("invalid encoding %s", enc)
	}

	t.Logf("%s encoding of %q is %q", enc, str, encodingResult)

	if expectedEncodedResult != encodingResult {
		t.Errorf(
			"expected encoding result %q. got %q",
			expectedEncodedResult,
			encodingResult,
		)
	}
}

func runDecodeTest(
	t *testing.T,
	enc encoding,
	str string,
	expectError bool,
	expectedDecodedResult string,
) {
	var decodingResult string
	var err error
	switch enc {
	case base64:
		decodingResult, err = base64Encoding.DecodeBase64(str)
	case base64url:
		decodingResult, err = base64Encoding.DecodeBase64URL(str)
	default:
		t.Fatalf("invalid encoding %s", enc)
	}

	if err != nil {
		t.Logf("%s decoding of %q resulted in error %s", enc, str, err.Error())
	} else {
		t.Logf("%s decoding of %q is %q", enc, str, decodingResult)
	}

	if expectError == true {
		if err == nil {
			t.Errorf(
				"expected error while decoding %q. no error was raised, got decoding result %q",
				str,
				decodingResult,
			)
		}
	} else {
		if err != nil {
			t.Errorf("expected decoding result %q. got error %s", decodingResult, err.Error())
		} else if expectedDecodedResult != decodingResult {
			t.Errorf("expected decoding result %q. got %q", expectedDecodedResult, decodingResult)
		}
	}
}

func base64ToBase64URL(str string) string {
	stripPaddingAndSpaceCharacters := strings.TrimRight(str, "\n\r=")

	replacePlus := strings.ReplaceAll(stripPaddingAndSpaceCharacters, "+", "-")
	replaceSlash := strings.ReplaceAll(replacePlus, "/", "_")

	return replaceSlash
}

func TestEncodeBase64(t *testing.T) {
	t.Run("Octets", func(t *testing.T) {
		runEncodeTest(t, base64, globalTests[0].str, globalTests[0].encodedStr)
	})

	t.Run("Core", func(t *testing.T) {
		for i := 1; i <= 17; i++ {
			runEncodeTest(t, base64, globalTests[i].str, globalTests[i].encodedStr)
		}
	})

	t.Run("JS", func(t *testing.T) {
		for i := 18; i <= 27; i++ {
			runEncodeTest(t, base64, globalTests[i].str, globalTests[i].encodedStr)
		}
	})
}

func TestEncodeBase64URL(t *testing.T) {
	t.Run("Octets", func(t *testing.T) {
		runEncodeTest(
			t,
			base64url,
			globalTests[0].str,
			base64ToBase64URL(globalTests[0].encodedStr),
		)
	})

	t.Run("Core", func(t *testing.T) {
		for i := 1; i <= 17; i++ {
			runEncodeTest(
				t,
				base64url,
				globalTests[i].str,
				base64ToBase64URL(globalTests[i].encodedStr),
			)
		}
	})

	t.Run("JS", func(t *testing.T) {
		for i := 18; i <= 27; i++ {
			runEncodeTest(
				t,
				base64url,
				globalTests[i].str,
				base64ToBase64URL(globalTests[i].encodedStr),
			)
		}
	})
}

func TestDecodeBase64(t *testing.T) {
	t.Run("Octets", func(t *testing.T) {
		runDecodeTest(t, base64, globalTests[0].encodedStr, false, globalTests[0].str)
	})

	t.Run("Core", func(t *testing.T) {
		for i := 1; i <= 17; i++ {
			runDecodeTest(
				t,
				base64,
				globalTests[i].encodedStr,
				false,
				globalTests[i].str,
			)
		}
	})

	t.Run("JS", func(t *testing.T) {
		for i := 18; i <= 27; i++ {
			runDecodeTest(
				t,
				base64,
				globalTests[i].encodedStr,
				false,
				globalTests[i].str,
			)
		}
	})

	t.Run("Invalid Characters", func(t *testing.T) {
		for i := range globalTestsWithInvalidCharacters {
			runDecodeTest(
				t,
				base64,
				globalTestsWithInvalidCharacters[i].encodedStr,
				true,
				globalTestsWithInvalidCharacters[i].str,
			)
		}
		runDecodeTest(t, base64, "AB__", true, "")
		runDecodeTest(t, base64, "A-BC", true, "")
		runDecodeTest(t, base64, "AB_-", true, "")
	})

	t.Run("Invalid Padding/Length", func(t *testing.T) {
		runDecodeTest(t, base64, "A", true, "")
		runDecodeTest(t, base64, "A=", true, "")
		runDecodeTest(t, base64, "A==", true, "")
		runDecodeTest(t, base64, "A===", true, "")
		runDecodeTest(t, base64, "AB", true, "")
		runDecodeTest(t, base64, "AB=", true, "")
		runDecodeTest(t, base64, "ABC", true, "")
	})

	t.Run("Non Strict Behavior", func(t *testing.T) {
		runDecodeTest(t, base64, "AA==", false, "\x00")
		runDecodeTest(t, base64, "AB==", false, "\x00")
		runDecodeTest(t, base64, "AAA=", false, "\x00\x00")
		runDecodeTest(t, base64, "AAB=", false, "\x00\x00")
	})
}

func TestDecodeBase64URL(t *testing.T) {
	t.Run("Octets", func(t *testing.T) {
		runDecodeTest(
			t,
			base64url,
			base64ToBase64URL(globalTests[0].encodedStr),
			false,
			globalTests[0].str,
		)
	})

	t.Run("Core", func(t *testing.T) {
		for i := 1; i <= 17; i++ {
			runDecodeTest(
				t,
				base64url,
				base64ToBase64URL(globalTests[i].encodedStr),
				false,
				globalTests[i].str,
			)
		}
	})

	t.Run("JS", func(t *testing.T) {
		for i := 18; i <= 27; i++ {
			runDecodeTest(
				t,
				base64url,
				base64ToBase64URL(globalTests[i].encodedStr),
				false,
				globalTests[i].str,
			)
		}
	})

	t.Run("Invalid Characters", func(t *testing.T) {
		for i := range globalTestsWithInvalidCharacters {
			runDecodeTest(
				t,
				base64url,
				base64ToBase64URL(globalTestsWithInvalidCharacters[i].encodedStr),
				true,
				globalTestsWithInvalidCharacters[i].str,
			)
		}
		runDecodeTest(t, base64url, "AB//", true, "")
		runDecodeTest(t, base64url, "A+BC", true, "")
		runDecodeTest(t, base64url, "AB/+", true, "")
	})

	t.Run("Invalid Padding/Length", func(t *testing.T) {
		runDecodeTest(t, base64url, "A", true, "")
		runDecodeTest(t, base64url, "A=", true, "")
		runDecodeTest(t, base64url, "A==", true, "")
		runDecodeTest(t, base64url, "A===", true, "")
		runDecodeTest(t, base64url, "AB=", true, "")
		runDecodeTest(t, base64url, "AB==", true, "")
		runDecodeTest(t, base64url, "ABC=", true, "")
	})

	t.Run("Non Strict Behavior", func(t *testing.T) {
		runDecodeTest(t, base64url, "AA", false, "\x00")
		runDecodeTest(t, base64url, "AB", false, "\x00")
		runDecodeTest(t, base64url, "AAA", false, "\x00\x00")
		runDecodeTest(t, base64url, "AAB", false, "\x00\x00")
	})
}

func ExampleEncodeBase64() {
	encodedString := base64Encoding.EncodeBase64("Hello, World!")
	_, _ = fmt.Println(encodedString)
	// Output: SGVsbG8sIFdvcmxkIQ==
}

func ExampleDecodeBase64() {
	decodedString, err := base64Encoding.DecodeBase64("SGVsbG8sIFdvcmxkIQ==")
	if err != nil {
		_, _ = fmt.Printf("Error: %s\n", err.Error())
	}
	_, _ = fmt.Println(decodedString)
	// Output: Hello, World!
}

func ExampleEncodeBase64URL() {
	encodedString := base64Encoding.EncodeBase64URL("Hello, World!")
	_, _ = fmt.Println(encodedString)
	// Output: SGVsbG8sIFdvcmxkIQ
}

func ExampleDecodeBase64URL() {
	decodedString, err := base64Encoding.DecodeBase64URL("SGVsbG8sIFdvcmxkIQ")
	if err != nil {
		_, _ = fmt.Printf("Error: %s\n", err.Error())
	}
	_, _ = fmt.Println(decodedString)
	// Output: Hello, World!
}

func BenchmarkEncodeBase64(b *testing.B) {
	for b.Loop() {
		base64Encoding.EncodeBase64(
			"Test encoding {\"key\": \"value\", \"key2\": \"value2\"}.{\"key3\": \"value3\", \"key4\": \"\"}{\n\"sub\": \"1234567890\",\n\r\"name\":\n \"John Doe\",\n\"iat\": 1516239022}. I am obviously testin JWT",
		)
	}
}

func BenchmarkEncodeBase64URL(b *testing.B) {
	for b.Loop() {
		base64Encoding.EncodeBase64URL(
			"Test encoding {\"key\": \"value\", \"key2\": \"value2\"}.{\"key3\": \"value3\", \"key4\": \"\"}{\n\"sub\": \"1234567890\",\n\r\"name\":\n \"John Doe\",\n\"iat\": 1516239022}. I am obviously testin JWT",
		)
	}
}

func BenchmarkDecodeBase64(b *testing.B) {
	for b.Loop() {
		_, err := base64Encoding.DecodeBase64(
			"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlzcyI6Imh0dHBzOi8vYmlnLWp3dC1pc3N1ZXIuZGV2In0tWX8L77aFUh+pEBKv3CdinhExUkWvQEgGMxJal5apNI",
		)
		if err != nil {
			b.Fatalf("Error: %s", err.Error())
		}
	}
}

func BenchmarkDecodeBase64URL(b *testing.B) {
	for b.Loop() {
		_, err := base64Encoding.DecodeBase64URL(
			"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlzcyI6Imh0dHBzOi8vYmlnLWp3dC1pc3N1ZXIuZGV2In0tWX8L77aFUh-pEBKv3CdinhExUkWvQEgGMxJal5apNI",
		)
		if err != nil {
			b.Fatalf("Error: %s", err.Error())
		}
	}
}
