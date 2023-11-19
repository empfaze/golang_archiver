package vlc

import "strings"

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

func (et encodingTable) GetDecodingTree() DecodingTree {
	result := DecodingTree{}

	for char, code := range et {
		result.Add(code, char)
	}

	return result
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	currentNode := dt

	for _, char := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}

		switch char {
		case '0':
			currentNode = currentNode.Zero
		case '1':
			currentNode = currentNode.One
		}
	}

	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = dt
	}

	return buf.String()
}

func (dt *DecodingTree) Add(code string, value rune) {
	currentNode := dt

	for _, char := range code {
		switch char {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &DecodingTree{}
			}

			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &DecodingTree{}
			}

			currentNode = currentNode.One
		}
	}

	currentNode.Value = string(value)
}
