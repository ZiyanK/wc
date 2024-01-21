package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Display all the results",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a file path")
		}

		filePath := args[0]
		fileStat, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("File does not exist")
			return
		}
		bytes := fileStat.Size()

		lineCount := 0
		wordCount := 0
		characterCount := 0

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("File does not exist")
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _ := reader.ReadString('\n')
			if err != nil {
				break
			}
			lineCount++
			characterCount += len(line)

			fields := strings.Fields(line)
			for _, _ = range fields {
				wordCount++
			}
			if line == "" {
				break
			}
		}

		// Since we count an extra line
		lineCount--

		fmt.Printf("%v %v %v %v", bytes, lineCount, wordCount, characterCount)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
