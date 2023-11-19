package lib

import (
	"reflect"
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

func Test_Encode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	tests := []struct {
		name        string
		encodedData []byte
		want        string
	}{
		{
			name:        "base test",
			encodedData: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
