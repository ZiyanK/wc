package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lineCountCmd represents the lc command
var lineCountCmd = &cobra.Command{
	Use:   "lc",
	Short: "Display the number of lines in the file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a file path")
		}

		filePath := args[0]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("File does not exist")
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		lineCount := 0
		for {
			_, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			lineCount++
		}

		fmt.Println(lineCount)
	},
}

func init() {
	rootCmd.AddCommand(lineCountCmd)
}
