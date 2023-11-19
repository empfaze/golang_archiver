package lib

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

type HexChunk string
type HexChunks []HexChunk

const CHUNK_SIZE = 8
const SEPARATOR = " "

func (bc BinaryChunks) ToHex() HexChunks {
	result := make(HexChunks, 0, len(bc))

	for _, chunk := range bc {
		hexChunk := chunk.ToHex()
		result = append(result, hexChunk)
	}

	return result
}

func (bc BinaryChunk) ToHex() HexChunk {
	number, err := strconv.ParseUint(string(bc), 2, CHUNK_SIZE)
	if err != nil {
		panic("Can't parse binary chunk: " + err.Error())
	}

	hexNumber := fmt.Sprintf("%x", number)
	uppercasedHexNumber := strings.ToUpper(hexNumber)

	if len(uppercasedHexNumber) == 1 {
		return HexChunk("0" + uppercasedHexNumber)
	}

	return HexChunk(uppercasedHexNumber)
}

func (hc HexChunks) ToBinary() BinaryChunks {
	result := make(BinaryChunks, 0, len(hc))

	for _, chunk := range hc {
		binaryChunk := chunk.ToBinary()
		result = append(result, binaryChunk)
	}

	return result
}

func (hc HexChunk) ToBinary() BinaryChunk {
	number, err := strconv.ParseUint(string(hc), 16, CHUNK_SIZE)
	if err != nil {
		panic("Can't parse hex chunk: " + err.Error())
	}

	result := fmt.Sprintf("%08b", number)

	return BinaryChunk(result)
}

func (hc HexChunks) ToString() string {
	switch len(hc) {
	case 0:
		return ""
	case 1:
		return string(string(hc[0]))
	}

	var buf strings.Builder

	for index, chunk := range hc {
		buf.WriteString(string(chunk))

		if index != len(hc)-1 {
			buf.WriteString(SEPARATOR)
		}
	}

	return buf.String()
}

func (bc BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bc {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func GetHexChunks(str string) HexChunks {
	chunks := strings.Split(str, SEPARATOR)

	result := make(HexChunks, 0, len(chunks))

	for _, chunk := range chunks {
		result = append(result, HexChunk(chunk))
	}

	return result
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
