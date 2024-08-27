package conciseStringCompress_test

import (
	"fmt"
	"testing"

	conciseStringCompress "github.com/valorwing/ConciseStringCompress"
)

var DefaultAlphabet = []rune{
	' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'r', 'h', 'l', 'd', 'c', 'u', 'm', 'f', 'p', 'g', 'w', 'b',
	'y', 'v', 'k', 'x', 'j', 'q', 'z', '.', ',', '-', '_', '!', '?', ':', ';', '"', '\'', '(', ')', '[',
	']', '{', '}', '<', '>', '/', '\\', '|', '@', '#', '$', '%', '^', '&', '*', '+', '=', '~', '`', '0',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K',
	'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '\n',
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
	if data[len(data)-1]&1 == 0 {
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
	if data[len(data)-1]&1 == 0 {
		t.Fail()
	}
	restored := compressor.DecompressString(data)
	if str != restored {
		t.Fail()
	}
}

func TestMediumText(t *testing.T) {
	str := `Gallia est omnis divisa in partes tres,
	quarum unam incolunt Belgae, aliam Aquitani, tertiam qui ipsorum lingua Celtae, nostra Galli appellantur.`

	compressor := conciseStringCompress.NewDefaultCompressor()
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

	str := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec tristique nisi, ut posuere dolor. Maecenas consectetur imperdiet nunc et interdum. Duis rutrum nisl non cursus volutpat. Proin cursus gravida pellentesque. Proin volutpat et felis non varius. Integer dictum consectetur rutrum. Nunc a libero iaculis, gravida sem non, accumsan neque. Proin sit amet felis placerat, vestibulum risus sit amet, sollicitudin tellus. Sed non posuere mauris. Pellentesque vitae odio ac tellus scelerisque rutrum. Mauris ut suscipit nisl. Proin rutrum venenatis eros, eu malesuada dolor. Vestibulum bibendum enim vitae viverra sodales. Morbi congue lacinia risus, quis posuere mi suscipit in. Donec nulla.`

	compressor := conciseStringCompress.NewDefaultCompressor()
	rawData := []byte(str)

	data, err := compressor.CompressString(str)
	if err != nil {
		t.Fail()
	}
	if data[len(data)-1]&1 == 0 {
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
