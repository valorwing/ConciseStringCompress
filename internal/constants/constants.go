package constants

const ErrInvalidString = " is not in alhabet. You can use package function get alphabet for get all supported charachters"
const ErrInvalidAlphabetLength = "alpabet length must be equal 64. not "

const NetworkFixByte uint8 = 0b10000001

var DefaultAlphabet = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'.', ',', '"', '\n', '\\', '/', ' ', '%', '@', '-', '+', '!',
}
