package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var bytesCmd = &cobra.Command{
	Use:   "bytes",
	Short: "Display number of bytes in the file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a file path")
		}

		filePath := args[0]
		file, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("File does not exist")
			return
		}

		fmt.Println(file.Size())
	},
}

func init() {
	rootCmd.AddCommand(bytesCmd)
}
