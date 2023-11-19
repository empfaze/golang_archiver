package vlc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string

type BinaryChunk string
type BinaryChunks []BinaryChunk

const CHUNK_SIZE = 8
const SEPARATOR = " "

func NewBinaryChunks(data []byte) BinaryChunks {
	result := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		result = append(result, NewBinaryChunk(code))
	}

	return result
}

func NewBinaryChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs BinaryChunks) ToBytes() []byte {
	result := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		result = append(result, bc.ToByte())
	}

	return result
}

func (bc BinaryChunk) ToByte() byte {
	number, err := strconv.ParseUint(string(bc), 2, CHUNK_SIZE)
	if err != nil {
		panic("Cannot convert binary chunk: " + err.Error())
	}

	return byte(number)
}

func (bc BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bc {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func getSplittedByChunksString(binaryString string, chunkSize int) BinaryChunks {
	stringLength := utf8.RuneCountInString(binaryString)

	numberOfChunks := math.Ceil(float64(stringLength / chunkSize))

	binaryChunksSlice := make(BinaryChunks, 0, int(numberOfChunks))

	var buf strings.Builder

	for index, char := range binaryString {
		buf.WriteString(string(char))

		if (index+1)%chunkSize == 0 {
			binaryChunksSlice = append(binaryChunksSlice, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		binaryChunksSlice = append(binaryChunksSlice, BinaryChunk(lastChunk))
	}

	return binaryChunksSlice
}
