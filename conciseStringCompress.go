package conciseStringCompress

import (
	"fmt"
	"math"
	"strings"
	"sync"

	bitutil "github.com/valorwing/ConciseStringCompress/internal/bitUtil"
	"github.com/valorwing/ConciseStringCompress/internal/constants"
)

type CompressorConfig struct {
	NetworkFixByteEnabled  bool
	TabOverAlphabetEnabled bool
}

type Compressor struct {
	alphabet     []rune
	alphabetLock sync.RWMutex
	alphabetMap  map[rune]uint8
	configLock   sync.RWMutex
	config       CompressorConfig
}

func NewCustomAlphabetCompressor(alphabet []rune, config CompressorConfig) *Compressor {
	retVal := &Compressor{
		alphabet:     nil,
		alphabetLock: sync.RWMutex{},
		alphabetMap:  nil,
		configLock:   sync.RWMutex{},
		config:       config,
	}

	retVal.SetAlphabet(alphabet)
	return retVal
}

func NewDefaultCompressor() *Compressor {
	return NewCustomAlphabetCompressor(constants.DefaultAlphabet, CompressorConfig{})
}

func (c *Compressor) SetAlphabet(alphabet []rune) error {
	if len(alphabet) != constants.AlphabetLength {
		return fmt.Errorf(constants.ErrInvalidAlphabetLengthFormat, constants.AlphabetLength, len(alphabet))
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

func (c *Compressor) GetConfig() CompressorConfig {
	c.configLock.RLock()
	defer c.configLock.RUnlock()
	return c.config
}

func (c *Compressor) SetConfig(cfg CompressorConfig) {
	c.configLock.Lock()
	c.config = cfg
	c.configLock.Unlock()
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

func (c *Compressor) CompressWithoutPack(input string) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}
	c.configLock.RLock()
	if c.config.TabOverAlphabetEnabled {
		input = strings.ReplaceAll(input, "\t", "    ")
	}
	c.configLock.RUnlock()
	retVal := make([]byte, len([]byte(input)))

	c.alphabetLock.RLock()
	defer c.alphabetLock.RUnlock()
	writeBuffByte := byte(0)
	for i, r := range input {
		if !c.isInAlphabet(r) {
			return nil, fmt.Errorf(constants.ErrInvalidStringFormat, string(r), constants.AlphabetLength)
		}

		writeBuffByte = c.getIndex(r)
		if writeBuffByte < constants.FiveBitsAlphabetPartLength {

			writeBuffByte = bitutil.ResetBit(writeBuffByte, 5)
		} else {

			writeBuffByte -= constants.FiveBitsAlphabetPartLength
			writeBuffByte = bitutil.SetBit(writeBuffByte, 6)
		}
		retVal[i] = writeBuffByte
	}
	return retVal, nil
}

func (c *Compressor) CompressString(input string) ([]byte, error) {
	if len(input) == 0 {
		return []byte{}, nil
	}
	c.configLock.RLock()
	if c.config.TabOverAlphabetEnabled {
		input = strings.ReplaceAll(input, "\t", "    ")
	}
	c.configLock.RUnlock()
	retVal := make([]byte, ((len([]byte(input))*7)/8)+1)
	retValPtr := &retVal

	byteOffset := uint64(0)
	bitOffset := uint8(0)
	writed6Bits := uint64(0)
	writed7Bits := uint64(0)
	c.alphabetLock.RLock()
	defer c.alphabetLock.RUnlock()

	for _, r := range input {
		if !c.isInAlphabet(r) {
			return nil, fmt.Errorf(constants.ErrInvalidStringFormat, string(r), constants.AlphabetLength)
		}

		index := c.getIndex(r)
		var writeBuffByte byte
		var writeBitsLen uint8
		writeBuffByte = index
		if index < constants.FiveBitsAlphabetPartLength {

			writeBitsLen = 6
			writeBuffByte = bitutil.ResetBit(writeBuffByte, 5)
		} else {
			writeBitsLen = 7
			writeBuffByte -= constants.FiveBitsAlphabetPartLength
			writeBuffByte = bitutil.SetBit(writeBuffByte, 6)
		}
		bitutil.WriteBits(writeBuffByte, writeBitsLen, retValPtr, byteOffset, bitOffset)
		if writeBitsLen == 6 {
			writed6Bits++
		} else {
			writed7Bits++
		}
		bitOffset += writeBitsLen
		if bitOffset > 7 {
			bitOffset -= 8
			byteOffset++
		}
	}

	outLenBits := writed6Bits*6 + writed7Bits*7
	outLen := int(math.Ceil(float64(outLenBits) / 8.0))
	retVal = retVal[:outLen]

	c.configLock.RLock()
	if c.config.NetworkFixByteEnabled {
		if retVal[len(retVal)-1]&1 == 0 || retVal[len(retVal)-1] == constants.NetworkFixByte {
			retVal = append(retVal, constants.NetworkFixByte)
		}
	}
	c.configLock.RUnlock()
	return retVal, nil
}

func (c *Compressor) UnpackedDecompress(input []byte) string {
	if len(input) == 0 {
		return ""
	}
	retVal := make([]rune, len(input))
	c.alphabetLock.RLock()
	defer c.alphabetLock.RUnlock()
	for i := range input {

		if !bitutil.ReadBit(input[i], 6) {
			retVal[i] = c.alphabet[input[i]]
		} else {
			tmp := bitutil.ResetBit(input[i], 6)
			retVal[i] = c.alphabet[tmp+constants.FiveBitsAlphabetPartLength]
		}
	}
	return string(retVal)
}

func (c *Compressor) DecompressString(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	retVal := make([]rune, 0, len(input)*2)
	inputLen := uint64(len(input))
	currentReadBitsOffset := uint8(0)
	currentReadByteOffset := uint64(0)

	c.configLock.RLock()
	if c.config.NetworkFixByteEnabled {
		if input[len(input)-1] == constants.NetworkFixByte {
			input = input[:len(input)-1]
		}
	}
	c.configLock.RUnlock()

	inputPtr := &input
	c.alphabetLock.RLock()
	bitByte := byte(0)
	readBitsOk := false
	for {
		bitByte = 0
		bitByte, readBitsOk = bitutil.ReadBits(inputPtr, 1, currentReadByteOffset, currentReadBitsOffset)
		if !readBitsOk {
			break
		}

		currentReadBitsOffset += 1
		if currentReadBitsOffset > 7 {
			currentReadByteOffset += 1
			currentReadBitsOffset -= 8
			if currentReadByteOffset == inputLen {
				break
			}
		}

		if !bitutil.ReadBit(bitByte, 0) {
			bitByte = 0
			bitByte, readBitsOk = bitutil.ReadBits(inputPtr, 5, currentReadByteOffset, currentReadBitsOffset)
			currentReadBitsOffset += 5
		} else {
			bitByte = 0
			bitByte, readBitsOk = bitutil.ReadBits(inputPtr, 6, currentReadByteOffset, currentReadBitsOffset)
			bitByte += constants.FiveBitsAlphabetPartLength
			currentReadBitsOffset += 6
		}
		if readBitsOk {
			retVal = append(retVal, c.alphabet[bitByte])
		} else {
			break
		}

		if currentReadBitsOffset > 7 {
			currentReadBitsOffset -= 8
			currentReadByteOffset += 1
			if currentReadByteOffset == inputLen {
				break
			}
		}
	}
	c.alphabetLock.RUnlock()

	retValStr := string(retVal)
	c.configLock.RLock()
	if c.config.TabOverAlphabetEnabled {
		retValStr = strings.ReplaceAll(retValStr, "    ", "\t")
	}
	c.configLock.RUnlock()
	return retValStr
}
