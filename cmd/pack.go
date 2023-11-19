package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/empfaze/archiver/lib/compression"
	"github.com/empfaze/archiver/lib/compression/vlc"
	"github.com/spf13/cobra"
)

const PACKED_FILE_EXTENSION = "vlc"

var ErrEmptyPath = errors.New("Path to file is missing")

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(getPackedFileName(filePath), packed, 0644)
	if err != nil {
		handleError(err)
	}
}

func getPackedFileName(path string) string {
	fileName := filepath.Base(path)
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)

	return baseName + "." + PACKED_FILE_EXTENSION
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression methods")

	err := packCmd.MarkFlagRequired("method")
	if err != nil {
		panic(err)
	}
}
