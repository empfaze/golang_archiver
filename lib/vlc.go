package lib

import (
	"strings"
	"unicode"
)

func Encode(str string) []byte {
	preparedText := getPreparedText(str)
	binaryString := getEncodedBinary(preparedText)
	splittedByChunksString := getSplittedByChunksString(binaryString, CHUNK_SIZE)

	return splittedByChunksString.ToBytes()
}

func getPreparedText(str string) string {
	var buf strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}

func getEncodedBinary(str string) string {
	var buf strings.Builder

	for _, char := range str {
		buf.WriteString(getBinaryNumber(char))
	}

	return buf.String()
}

func getBinaryNumber(char rune) string {
	table := getEncodingTable()

	result, ok := table[char]
	if !ok {
		panic("Unknown character: " + string(char))
	}

	return result
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

func Decode(encodedData []byte) string {
	binaryChunks := NewBinaryChunks(encodedData)
	binaryString := binaryChunks.Join()
	decodingTree := getEncodingTable().GetDecodingTree()

	return getExportedText(decodingTree.Decode(binaryString))
}

func getExportedText(str string) string {
	var buf strings.Builder

	isCapitalLetter := false

	for _, char := range str {
		if isCapitalLetter {
			buf.WriteString(string(unicode.ToUpper(char)))
			isCapitalLetter = false

			continue
		}

		if char == '!' {
			isCapitalLetter = true
		} else {
			buf.WriteString(string(char))
		}
	}

	return buf.String()
}
