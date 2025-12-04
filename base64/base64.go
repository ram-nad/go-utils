// Package base64 provides support for working with base64 and base64url encoded strings
package base64

import (
	"encoding/base64"
	"unsafe"
)

// Unsafe conversion to skip expensive string to []byte conversion
// Only used when we can guarantee that the underlying bytes won't be modified
func bytesFromStr(str string) []byte {
	//nolint:gosec // this slice won't be modified
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

// Unsafe conversion to skip expensive []byte to string conversion
// Only used when we don't keep the []byte around
func strFromBytes(data []byte) string {
	//nolint:gosec // we won't use the slice after this conversion
	return unsafe.String(unsafe.SliceData(data), len(data))
}

/*
EncodeBase64Bytes encodes bytes to standard base64.

To be used for encoding binary data.
*/
func EncodeBase64Bytes(data []byte) string {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(buf, data)
	return strFromBytes(buf)
}

/*
DecodeBase64Bytes decodes base64 encoded string to bytes.

To be used when decoding to binary data.
*/
func DecodeBase64Bytes(str string) ([]byte, error) {
	bufLen := base64.StdEncoding.DecodedLen(len(str))
	dst := make([]byte, bufLen)
	n, err := base64.StdEncoding.Decode(dst, bytesFromStr(str))
	// N <= bufLen, because \r and \n are ignored when decoding
	return dst[:n], err
}

// EncodeBase64 encodes string to standard base64
func EncodeBase64(str string) string {
	return EncodeBase64Bytes(bytesFromStr(str))
}

// DecodeBase64 decodes base64 encoded string
func DecodeBase64(str string) (string, error) {
	decodedBytes, err := DecodeBase64Bytes(str)
	if err != nil {
		return "", err
	}

	return strFromBytes(decodedBytes), nil
}

/*
EncodeBase64URLBytes encodes bytes to base64url (without padding).

To be used for encoding binary data.
*/
func EncodeBase64URLBytes(data []byte) string {
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(data)))
	base64.RawURLEncoding.Encode(buf, data)
	return strFromBytes(buf)
}

/*
DecodeBase64URLBytes decodes base64url (without padding) encoded string to bytes.

To be used when decoding to binary data.
*/
func DecodeBase64URLBytes(str string) ([]byte, error) {
	bufLen := base64.RawURLEncoding.DecodedLen(len(str))
	dst := make([]byte, bufLen)
	// N <= bufLen, because \r and \n are ignored when decoding
	n, err := base64.RawURLEncoding.Decode(dst, bytesFromStr(str))
	return dst[:n], err
}

// EncodeBase64URL encodes string to base64url
func EncodeBase64URL(str string) string {
	return EncodeBase64URLBytes(bytesFromStr(str))
}

// DecodeBase64URL decodes base64url encoded string
func DecodeBase64URL(str string) (string, error) {
	decodedBytes, err := DecodeBase64URLBytes(str)
	if err != nil {
		return "", err
	}
	return strFromBytes(decodedBytes), nil
}
