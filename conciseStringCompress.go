package conciseStringCompress

import (
	"errors"
	"fmt"
	"math"
	"sync"

	bitutil "github.com/valorwing/ConciseStringCompress/internal/bitUtil"
	"github.com/valorwing/ConciseStringCompress/internal/constants"
)

var zeros []bool
var ones []bool

type Compressor struct {
	alphabet     []rune
	alphabetLock sync.RWMutex
	alphabetMap  map[rune]uint8
}

func NewCustomAlphabetCompressor(alphabet []rune) *Compressor {
	retVal := &Compressor{
		alphabet:     nil,
		alphabetLock: sync.RWMutex{},
		alphabetMap:  nil,
	}

	retVal.SetAlphabet(alphabet)
	return retVal
}

func NewDefaultCompressor() *Compressor {

	return NewCustomAlphabetCompressor(constants.DefaultAlphabet)
}

func (c *Compressor) SetAlphabet(alphabet []rune) error {

	if len(alphabet) != 64 {
		return errors.New(constants.ErrInvalidAlphabetLength + fmt.Sprint(len(alphabet)))
	}

	newAlphabetMap := map[rune]uint8{}

	for i := range alphabet {
		newAlphabetMap[alphabet[i]] = uint8(i)
	}

	c.alphabetLock.Lock()
	c.alphabet = alphabet
	c.alphabetMap = newAlphabetMap
	c.alphabetLock.Unlock()
	return nil
}

func (c *Compressor) GetAlphabet() []rune {

	c.alphabetLock.RLock()
	retVal := c.alphabet
	c.alphabetLock.RUnlock()

	return retVal
}

func (c *Compressor) isInAlphabet(r rune) bool {
	c.alphabetLock.RLock()
	_, ok := c.alphabetMap[r]
	c.alphabetLock.RUnlock()
	return ok
}

func (c *Compressor) getIndex(r rune) uint8 {

	v, ok := c.alphabetMap[r]

	if !ok {
		panic("not in alphabet")
	}
	return v
}

func (c *Compressor) CompressString(input string) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}

	var retVal []byte = make([]uint8, ((len([]byte(input))*6)/8)+1)
	retValPtr := &retVal

	byteOffset := uint64(0)
	bitOffset := uint8(0)
	writedRunes := 0
	c.alphabetLock.RLock()
	for _, r := range input {

		if !c.isInAlphabet(r) {
			return nil, errors.New(string(r) + constants.ErrInvalidString)
		}
		bitutil.WriteBits(byte(c.getIndex(r)), 6, retValPtr, byteOffset, bitOffset)
		writedRunes++
		bitOffset += 6
		if bitOffset >= 8 {
			bitOffset -= 8
			byteOffset += 1
		}
	}
	c.alphabetLock.RUnlock()
	outLenBits := writedRunes * 6

	outLen := int(math.Floor(float64(outLenBits) / float64(8)))

	if outLenBits%8 != 0 {
		outLen += 1
	}

	retVal = retVal[:outLen]

	if retVal[len(retVal)-1]&1 == 0 || retVal[len(retVal)-1] == constants.NetworkFixByte {
		retVal = append(retVal, constants.NetworkFixByte)
	}

	return retVal, nil
}

func (c *Compressor) DecompressString(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	retVal := make([]rune, 0, len(input)*2)

	currentReadBitsOffset := uint8(0)
	currentReadByteOffset := uint64(0)

	if input[len(input)-1] == constants.NetworkFixByte {
		input = input[:len(input)-1]
	}

	inputPtr := &input
	c.alphabetLock.RLock()
	for {

		bitByte, readBitsOk := bitutil.ReadBits(inputPtr, 6, currentReadByteOffset, currentReadBitsOffset)
		if readBitsOk {
			retVal = append(retVal, c.alphabet[bitByte])
		} else {
			break
		}

		stringByte := fmt.Sprintf("%08b", bitByte)[2:]

		fmt.Println(stringByte + " " + stringByte[:1] + " " + stringByte[5:] + " ")

		currentReadBitsOffset += 6
		if currentReadBitsOffset >= 8 {
			currentReadBitsOffset -= 8
			currentReadByteOffset += 1
			if currentReadByteOffset == uint64(len(input)) {
				break
			}
		}
	}
	c.alphabetLock.RUnlock()
	return string(retVal)
}
