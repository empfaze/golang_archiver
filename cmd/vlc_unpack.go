package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/empfaze/archiver/lib"
	"github.com/spf13/cobra"
)

const UNPACKED_FILE_EXTENSION = "txt"

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable length code",
	Run:   unpack,
}

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
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

	packed := lib.Decode(string(data))

	err = os.WriteFile(getUnpackedFileName(filePath), []byte(packed), 0644)
}

func getUnpackedFileName(path string) string {
	fileName := filepath.Base(path)
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)

	return baseName + "." + UNPACKED_FILE_EXTENSION
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
