// Package bitutil provides utilities for reading and writing a specified number
// of bits to and from a byte slice.
package bitutil

// SetdBit selective set `1` bit state by index
func SetBit(b byte, bitPosition uint8) byte {
	retVal := b
	switch bitPosition {
	case 0:
		retVal = b | 0b00000001
	case 1:
		retVal = b | 0b00000010
	case 2:
		retVal = b | 0b00000100
	case 3:
		retVal = b | 0b00001000
	case 4:
		retVal = b | 0b00010000
	case 5:
		retVal = b | 0b00100000
	case 6:
		retVal = b | 0b01000000
	case 7:
		retVal = b | 0b10000000
	default:
		panic("set invalid bit position")
	}
	return retVal
}

// ReadBit selective set `0` bit state by index
func ResetBit(b byte, bitPosition uint8) byte {
	retVal := b
	switch bitPosition {
	case 0:
		retVal = b & 0b11111110
	case 1:
		retVal = b & 0b11111101
	case 2:
		retVal = b & 0b11111011
	case 3:
		retVal = b & 0b11110111
	case 4:
		retVal = b & 0b11101111
	case 5:
		retVal = b & 0b11011111
	case 6:
		retVal = b & 0b10111111
	case 7:
		retVal = b & 0b01111111
	default:
		panic("reset invalid bit position")
	}
	return retVal
}

// ReadBit selective detect bit state by index
// transformed to bool `0` - false `1` - true
func ReadBit(b byte, bitPosition uint8) bool {
	retVal := b
	switch bitPosition {
	case 0:
		retVal = b & 0b00000001
	case 1:
		retVal = b & 0b00000010
	case 2:
		retVal = b & 0b00000100
	case 3:
		retVal = b & 0b00001000
	case 4:
		retVal = b & 0b00010000
	case 5:
		retVal = b & 0b00100000
	case 6:
		retVal = b & 0b01000000
	case 7:
		retVal = b & 0b10000000
	default:
		panic("read invalid bit position")
	}
	return retVal != 0
}

// WriteBits writes `writeBitLength` bits from `objectToWrite` to `targetArray`
// starting at the specified `initialByteOffset` and `initialBitOffset` in the target array.
// If necessary, the function automatically continues writing into the next byte.
// WARNING: The function does not manage the length of `targetArray`; you must ensure
// the array has sufficient space for the written bits.
func WriteBits(objectToWrite byte, writeBitLength uint8, targetArray *[]byte, initialByteOffset uint64, initialBitOffset uint8) {
	target := *targetArray
	source := objectToWrite

	writeBitLength--
	for i := writeBitLength; i != 255; i-- {
		if initialBitOffset > 7 {
			initialByteOffset++
			initialBitOffset -= 8
		}
		bit := source&(1<<i) != 0
		if bit {
			target[initialByteOffset] |= (1 << initialBitOffset)
		} else {
			target[initialByteOffset] &= ^(1 << initialBitOffset)
		}
		initialBitOffset++
	}
}

// ReadBits reads `bitLength` bits from `source`, starting at `byteOffset` and `bitOffset`
// in the specified byte. The function returns the read byte and a boolean value indicating
// whether the read operation stayed within the bounds of the source array.
// If necessary, the function automatically continues reading from the next byte.
func ReadBits(source *[]byte, bitLength uint8, byteOffset uint64, bitOffset uint8) (byte, bool) {
	var result byte = 0
	sourceTarget := *source
	sourceLen := uint64(len(*source))
	bitLength--
	for i := bitLength; i != 255; i-- {
		if bitOffset > 7 {
			byteOffset++
			if byteOffset >= sourceLen {
				return 0, false
			}
			bitOffset -= 8
		}
		bit := sourceTarget[byteOffset]&(1<<bitOffset) != 0
		if bit {
			result |= (1 << i)
		} else {
			result &= ^(1 << i)
		}

		bitOffset++
	}
	return result, true
}
