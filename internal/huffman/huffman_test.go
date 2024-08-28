package huffman_test

import (
	"fmt"
	"testing"

	conciseStringCompress "github.com/valorwing/ConciseStringCompress"
	bitutil "github.com/valorwing/ConciseStringCompress/internal/bitUtil"
	"github.com/valorwing/ConciseStringCompress/internal/constants"
	"github.com/valorwing/ConciseStringCompress/internal/huffman"
)

func TestBase(t *testing.T) {

	frequencies := map[byte]float64{}

	for i := range constants.DefaultAlphabet {

		if i < 32 {
			frequencies[byte(i)] = 610.0 - float64(i)
		} else {
			tmp := byte(i)
			tmp -= 32
			tmp = bitutil.SetBit(tmp, 6)
			frequencies[tmp] = 97.0 - float64(i)
		}
	}

	root := huffman.BuildTree(frequencies)

	codes := make(map[byte]string)
	huffman.GenerateCodes(root, "", codes)

	//fmt.Println("Huffman Codes:", codes)

	testString := `In the heart of the bustling city, amidst the tall skyscrapers and busy streets, lies a park of tranquility. This green oasis provides a respite from the urban rush, offering a serene escape where people can relax and rejuvenate. The park features sprawling lawns, picturesque ponds, and a variety of flora and fauna that attract both locals and visitors alike.

Amidst the park's natural beauty, one can find winding paths for leisurely strolls, vibrant flowerbeds, and shaded benches for quiet contemplation. Children can be seen playing on the swings and slides, their laughter echoing through the air. The park also hosts various community events and gatherings, making it a vibrant hub of social activity.

As the seasons change, the park transforms with the landscape. In spring, it bursts into color with blooming flowers. Summer brings lush greenery and warm sunlight. Autumn paints the leaves in shades of red and gold, while winter covers the park in a serene blanket of snow. Each season offers its own unique charm, making the park a beloved retreat throughout the year.`
	compressor := conciseStringCompress.NewDefaultCompressor()

	testFisrtStepCompressed, err := compressor.CompressWithoutPack(testString)
	if err != nil {
		t.Fail()
	}

	encoded := huffman.Encode(testFisrtStepCompressed, codes)
	fmt.Println("Encoded:", encoded)

	decoded := huffman.Decode(encoded, root)
	fmt.Println("Decoded:", []byte(decoded))

	fmt.Printf("Raw data len: %d, compressed data len: %d \n", len([]byte(testString)), len(encoded))
	fmt.Println("Compress rate: ", float64(len(encoded))/float64(len([]byte(testString))))

	decodedFinal := compressor.UnpackedDecompress(decoded)

	if testString != decodedFinal {
		t.Fail()
	}

}
