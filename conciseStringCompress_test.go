package conciseStringCompress_test

import (
	"fmt"
	"testing"

	conciseStringCompress "github.com/valorwing/ConciseStringCompress"
	"github.com/valorwing/ConciseStringCompress/internal/constants"
)

var DefaultAlphabet = constants.DefaultAlphabet

func TestRawWithoutPack(t *testing.T) {
	str := string(DefaultAlphabet)

	compressor := conciseStringCompress.NewDefaultCompressor()
	data, err := compressor.CompressWithoutPack(str)
	fmt.Println()
	fmt.Printf("Readed first : %08b \n", data[0])
	fmt.Printf("Readed second: %08b \n", data[1])
	if err != nil {
		t.Fail()
	}
	restored := compressor.UnpackedDecompress(data)
	if str != restored {
		t.Fail()
	}

}

func TestAlhpabet(t *testing.T) {

	str := string(DefaultAlphabet)

	compressor := conciseStringCompress.NewDefaultCompressor()
	data, err := compressor.CompressString(str)
	fmt.Println()
	fmt.Printf("Readed first : %08b \n", data[0])
	fmt.Printf("Readed second: %08b \n", data[1])
	if err != nil {
		t.Fail()
	}
	restored := compressor.DecompressString(data)
	if str != restored {
		t.Fail()
	}
}

func TestZeroOnly(t *testing.T) {

	compressor := conciseStringCompress.NewDefaultCompressor()
	str := "a"
	data, err := compressor.CompressString(str)
	if err != nil {
		t.Fail()
	}
	restored := compressor.DecompressString(data)
	if str != restored {
		t.Fail()
	}
}

func TestTabOverAlphabetAndNetworkFix(t *testing.T) {
	str := `Gallia est omnis divisa in partes tres,
	quarum unam incolunt Belgae, aliam Aquitani, tertiam qui ipsorum lingua Celtae, nostra Galli appellantur.`

	compressor := conciseStringCompress.NewCustomAlphabetCompressor(constants.DefaultAlphabet,
		conciseStringCompress.CompressorConfig{
			NetworkFixByteEnabled:  true,
			TabOverAlphabetEnabled: true,
		})
	data, err := compressor.CompressString(str)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if data[len(data)-1]&1 == 0 {
		t.Fail()
	}
	restored := compressor.DecompressString(data)
	if str != restored {
		fmt.Println("Fail: testData: ", str, " restored: ", restored)
		t.Fail()
	}
	rawData := []byte(str)
	if len(data) > len(rawData) {
		t.Fail()
	}
}

func TestBaseCompressAndDecompress(t *testing.T) {

	str := `In the vast expanse of the universe, amidst the countless stars and galaxies, lies a small blue planet, teeming with life. This planet, known as Earth, is home to a multitude of species, each uniquely adapted to its environment. From the towering mountains to the deep blue oceans, every corner of this world is filled with wonder and mystery.

The story of Earth is a story of evolution, adaptation, and survival. Over billions of years, life has evolved in countless ways, from simple single-celled organisms to complex beings capable of thought and emotion. Humans, one of the most recent arrivals on this ancient planet, have developed remarkable technologies, explored vast landscapes, and even ventured into the depths of space.

Yet, despite all our advancements, there is still so much we do not understand. The mysteries of the universe continue to elude us, and the more we learn, the more questions we have. What lies beyond the edge of the observable universe? Are we alone in this vast cosmos, or are there other intelligent beings out there, looking up at their skies and wondering the same thing?

As we look to the future, we must remember the importance of exploration, curiosity, and the pursuit of knowledge. Our journey has only just begun, and there are infinite possibilities waiting to be discovered.`
	compressor := conciseStringCompress.NewDefaultCompressor()
	rawData := []byte(str)

	data, err := compressor.CompressString(str)
	if err != nil {
		t.Fail()
	}
	restored := compressor.DecompressString(data)
	if str != restored {
		t.Fail()
	}

	if len(data) > len(rawData) {
		t.Fail()
	}

	fmt.Printf("Raw data len: %d compressed data len: %d, compress ratio: %f \n", len(rawData), len(data), float64(float64(len(data))/float64(len(rawData))))
}

const testString = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec tristique nisi, ut posuere dolor. Maecenas consectetur imperdiet nunc et interdum. Duis rutrum nisl non cursus volutpat. Proin cursus gravida pellentesque. Proin volutpat et felis non varius. Integer dictum consectetur rutrum. Nunc a libero iaculis, gravida sem non, accumsan neque. Proin sit amet felis placerat, vestibulum risus sit amet, sollicitudin tellus. Sed non posuere mauris. Pellentesque vitae odio ac tellus scelerisque rutrum. Mauris ut suscipit nisl. Proin rutrum venenatis eros, eu malesuada dolor. Vestibulum bibendum enim vitae viverra sodales. Morbi congue lacinia risus, quis posuere mi suscipit in. Donec nulla.`

func BenchmarkCompressString(b *testing.B) {

	compressor := conciseStringCompress.NewDefaultCompressor()
	for i := 0; i < b.N; i++ {
		data, err := compressor.CompressString(testString)
		if err != nil {
			b.Fatalf("CompressString failed: %v", err)
		}
		_ = data
	}
}

func BenchmarkDecompressString(b *testing.B) {

	compressor := conciseStringCompress.NewDefaultCompressor()

	data, err := compressor.CompressString(testString)
	if err != nil {
		b.Fatalf("CompressString failed: %v", err)
	}

	for i := 0; i < b.N; i++ {
		restored := compressor.DecompressString(data)
		if testString != restored {
			b.Fatalf("DecompressString failed: expected %s, got %s", testString, restored)
		}
	}
}
