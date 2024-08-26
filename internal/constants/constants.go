// Package constants provides error templates and immutable general variables
// used throughout the project.
package constants

// AlphabetLength represents the length of a composite alphabet with the first 32 runes
// using 5-bit encoding and the remaining 64 runes using 6-bit encoding.
// Encoding is done using a prefixed variable length: 6-bit encoded with 7-bit as `1xxxxxx`
// and 5-bit encoded with 6-bit as `0xxxxx`.
const AlphabetLength = 96

// DefaultAlphabet is an optimal alphabet sorted by popularity.
// It contains the most common 96 printable characters.
var DefaultAlphabet = []rune{
	' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'r', 'h', 'l', 'd', 'c', 'u', 'm', 'f', 'p', 'g', 'w', 'b',
	'y', 'v', 'k', 'x', 'j', 'q', 'z', '.', ',', '-', '_', '!', '?', ':', ';', '"', '\'', '(', ')', '[',
	']', '{', '}', '<', '>', '/', '\\', '|', '@', '#', '$', '%', '^', '&', '*', '+', '=', '~', '`', '0',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K',
	'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// ErrInvalidStringFormat is the error triggered when the input compressed string
// contains a rune not present in the alphabet.
// The format string uses two variables: the first %s is the incorrect rune,
// and the second %d is the allowed custom alphabet length.
const ErrInvalidStringFormat = "%s is not in the alphabet. You can use the package function to get the full supported characters, or you can define a custom %d-length alphabet."

// ErrInvalidAlphabetLengthFormat is the error triggered when attempting to set
// a custom alphabet with a length different from the constant AlphabetLength.
const ErrInvalidAlphabetLengthFormat = "Alphabet length must be equal to %d, not %d."

// NetworkFixByte is a magic byte used when the compressed vector ends with 0
// or when this byte is added to ensure correct decompression.
const NetworkFixByte uint8 = 0b10000001
