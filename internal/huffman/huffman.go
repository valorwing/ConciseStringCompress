package huffman

import (
	"container/heap"
)

type Node struct {
	char        uint16
	frequency   float64
	left, right *Node
}

type HuffmanHeap []*Node

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BuildTree(frequencies map[uint16]float64) *Node {
	h := &HuffmanHeap{}
	heap.Init(h)

	for char, freq := range frequencies {
		heap.Push(h, &Node{char: char, frequency: freq})
	}

	for h.Len() > 1 {
		left := heap.Pop(h).(*Node)
		right := heap.Pop(h).(*Node)

		merged := &Node{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}

		heap.Push(h, merged)
	}

	return heap.Pop(h).(*Node)
}

func GenerateCodes(node *Node, prefix string, codes map[uint16]string) {
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil {
		codes[node.char] = prefix
	}

	GenerateCodes(node.left, prefix+"0", codes)
	GenerateCodes(node.right, prefix+"1", codes)
}

func Encode(input []byte, codes map[uint16]string) []byte {
	var output []byte
	var currentByte byte
	var bitCount uint8

	var previosByte byte = 0
	var previosByteIsEmpty = true

	var lastRawByte = false
	var inputLen = uint64(len(input))
	var lastByteIndex = inputLen - 1

	var code string
	for i := uint64(0); i < inputLen; i++ {
		if previosByteIsEmpty {
			previosByteIsEmpty = false
			previosByte = input[i]
		} else {
			previosByteIsEmpty = true
			code = codes[uint16(previosByte)|uint16(input[i])<<8]
			for _, bit := range code {
				if bit == '1' {
					currentByte |= 1 << (7 - bitCount)
				}
				bitCount++
				if bitCount == 8 {
					output = append(output, currentByte)
					currentByte = 0
					bitCount = 0
				}
			}
		}

	}

	for char := range input {

		code := codes[input[char]]

	}

	if bitCount > 0 {
		output = append(output, currentByte)
	}

	return output
}

func Decode(encoded []byte, root *Node) []byte {
	var decoded []byte
	node := root
	var bitMask byte

	for _, b := range encoded {
		for bitMask = 0x80; bitMask > 0; bitMask >>= 1 {
			if b&bitMask > 0 {
				node = node.right
			} else {
				node = node.left
			}

			if node.left == nil && node.right == nil {
				decoded = append(decoded, node.char)
				node = root
			}
		}
	}

	return decoded
}
