package conciseStringCompress_test

import (
	"fmt"
	"testing"

	conciseStringCompress "github.com/valorwing/ConciseStringCompress"
)

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
	str := `Gallia est omnis divisa in partes tres, quarum unam incolunt Belgae, aliam Aquitani, tertiam qui ipsorum lingua Celtae, nostra Galli appellantur.`

	compressor := conciseStringCompress.NewDefaultCompressor()
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
