// Package bitutil provides utilities for reading and writing a specified number
// of bits to and from a byte slice.
package bitutil

// WriteBits writes `writeBitLength` bits from `objectToWrite` to `targetArray`
// starting at the specified `initialByteOffset` and `initialBitOffset` in the target array.
// If necessary, the function automatically continues writing into the next byte.
// WARNING: The function does not manage the length of `targetArray`; you must ensure
// the array has sufficient space for the written bits.
func WriteBits(objectToWrite byte, writeBitLength uint8, targetArray *[]byte, initialByteOffset uint64, initialBitOffset uint8) {
	target := *targetArray
	source := objectToWrite

	for i := uint8(0); i < writeBitLength; i++ {
		if initialBitOffset == 8 {
			initialByteOffset++
			initialBitOffset = 0
		}
		bit := (source >> i) & 1
		target[initialByteOffset] = (target[initialByteOffset] &^ (1 << initialBitOffset)) | (bit << initialBitOffset)
		initialBitOffset++
	}
}

// ReadBits reads `bitLength` bits from `source`, starting at `byteOffset` and `bitOffset`
// in the specified byte. The function returns the read byte and a boolean value indicating
// whether the read operation stayed within the bounds of the source array.
// If necessary, the function automatically continues reading from the next byte.
func ReadBits(source *[]byte, bitLength uint8, byteOffset uint64, bitOffset uint8) (byte, bool) {
	var result byte
	sourceTarget := *source
	sourceLen := uint64(len(*source))
	for i := uint8(0); i < bitLength; i++ {
		if bitOffset == 8 {
			byteOffset++
			if byteOffset >= sourceLen {
				return 0, false
			}
			bitOffset = 0
		}
		bit := (sourceTarget[byteOffset] >> bitOffset) & 1
		result |= bit << i
		bitOffset++
	}
	return result, true
}
