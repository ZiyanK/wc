package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// characterCountCmd represents the cc command
var characterCountCmd = &cobra.Command{
	Use:   "cc",
	Short: "Display the number of characters in the file",
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
		characterCount := 0

		for {
			line, _ := reader.ReadString('\n')
			for _, _ = range line {
				characterCount++
			}
			if line == "" {
				break
			}
		}

		fmt.Println(characterCount)
	},
}

func init() {
	rootCmd.AddCommand(characterCountCmd)
}
