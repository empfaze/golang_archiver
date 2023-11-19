package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/empfaze/archiver/lib/compression"
	"github.com/empfaze/archiver/lib/compression/vlc"
	"github.com/spf13/cobra"
)

const UNPACKED_FILE_EXTENSION = "txt"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("Unknown method")
	}

	filePath := args[0]

	file, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		handleError(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(getUnpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
}

func getUnpackedFileName(path string) string {
	fileName := filepath.Base(path)
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)

	return baseName + "." + UNPACKED_FILE_EXTENSION
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression methods")

	err := unpackCmd.MarkFlagRequired("method")
	if err != nil {
		panic(err)
	}
}
