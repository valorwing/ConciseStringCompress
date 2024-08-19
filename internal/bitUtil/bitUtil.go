package bitutil

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
func ReadBits(source *[]byte, bitLength uint8, byteOffset uint64, bitOffset uint8) (byte, bool) {
	var result byte
	var sourceTarget = *source
	for i := uint8(0); i < bitLength; i++ {
		if bitOffset == 8 {
			byteOffset++
			if byteOffset >= uint64(len(*source)) {
				return 0, false
			}
			bitOffset = 0
		}
		bit := sourceTarget[byteOffset] >> bitOffset & 1
		result |= bit << i
		bitOffset++
	}
	return result, true
}
