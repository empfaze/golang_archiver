package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/empfaze/archiver/lib"
	"github.com/spf13/cobra"
)

const PACKED_FILE_EXTENSION = "vlc"

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable length code",
	Run:   pack,
}

var ErrEmptyPath = errors.New("Path to file is missing")

func pack(_ *cobra.Command, args []string) {
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

	packed := lib.Encode(string(data))

	err = os.WriteFile(getPackedFileName(filePath), []byte(packed), 0644)
}

func getPackedFileName(path string) string {
	fileName := filepath.Base(path)
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)

	return baseName + "." + PACKED_FILE_EXTENSION
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
