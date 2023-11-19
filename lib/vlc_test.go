package lib

import (
	"testing"
)

func Test_getPreparedText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPreparedText(tt.str); got != tt.want {
				t.Errorf("getPreparedText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEncodedBinary(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!ted",
			want: "001000100110100101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncodedBinary(tt.str); got != tt.want {
				t.Errorf("getEncodedBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
