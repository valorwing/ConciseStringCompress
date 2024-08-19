
# ConciseStringCompress

**ConciseStringCompress** is a Go library that demonstrates a string compression technique with a guaranteed minimum compression rate of 24%. The library operates with a time complexity of O(n) and utilizes a 6-bit (64-character) alphabet. However, users have the flexibility to specify their own custom alphabet.
## Features

-   **Guaranteed Compression**: Achieve a minimum of 24% compression on strings.
-   **Custom Alphabet Support**: Use the default 6-bit alphabet or define your own 64-character alphabet.
-   **Efficient Compression**: The library compresses and decompresses strings with O(n) complexity.
-  **Network Transmission Ready**: Library guarantees that the last byte and bit will not be zero
## Installation
To use ConciseStringCompress, you can install it via `go get`:
```bash
go get github.com/valorwing/ConciseStringCompress
```
## Usage

### Basic Example

Here's a simple example demonstrating how to compress and decompress a string using the default alphabet:

```go
package main

import (
    "fmt"
    "github.com/valorwing/ConciseStringCompress"
)

func main() {
    compressor := conciseStringCompress.NewDefaultCompressor()
    
    input := "Hello, World!"
    compressed, err := compressor.CompressString(input)
    if err != nil {
        panic(err)
    }

    fmt.Println("Compressed:", compressed)
    
    decompressed := compressor.DecompressString(compressed)
    fmt.Println("Decompressed:", decompressed)
}

```
### Using a Custom Alphabet

You can also specify your own custom 64-character alphabet:
```go
package main

import (
    "fmt"
    "github.com/valorwing/ConciseStringCompress"
)

func main() {
    customAlphabet :=[]rune{

	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',

	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',

	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',

	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',

	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', ' ', '\n',
	}
                             
    compressor := conciseStringCompress.NewCustomAlphabetCompressor(customAlphabet)
    
    input := "Custom Alphabet Test 1"
    compressed, err := compressor.CompressString(input)
    if err != nil {
        panic(err)
    }

    fmt.Println("Compressed:", compressed)
    
    decompressed := compressor.DecompressString(compressed)
    fmt.Println("Decompressed:", decompressed)
}

```
## Benchmarks
You can check  bechmark code in file conciseStringCompress_test.go

```bash
goos: linux
goarch: amd64
pkg: github.com/valorwing/ConciseStringCompress
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz

BenchmarkCompressString-8     	   51217	     23678 ns/op	     576 B/op	       1 allocs/op
BenchmarkDecompressString-8   	  112016	     10349 ns/op	    5568 B/op	       2 allocs/op
```

## License

This library is licensed under the MIT license. See the LICENSE file for more details.
