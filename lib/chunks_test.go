package lib

import (
	"reflect"
	"testing"
)

func Test_getSplittedByChunksString(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base name",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSplittedByChunksString(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSplittedByChunksString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bc   BinaryChunks
		want HexChunks
	}{
		{
			name: "base test",
			bc:   BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "base test",
			str:  "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "base test",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunks
		want BinaryChunks
	}{
		{
			name: "base test",
			hc:   HexChunks{"2F", "80"},
			want: BinaryChunks{"00101111", "10000000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bc   BinaryChunks
		want string
	}{
		{
			name: "base test",
			bc:   BinaryChunks{"01001111", "10000000"},
			want: "0100111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
