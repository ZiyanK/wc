/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// wordCountCmd represents the wc command
var wordCountCmd = &cobra.Command{
	Use:   "wc",
	Short: "Display the number of words in the file",
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
		wordCount := 0

		for {
			line, _ := reader.ReadString('\n')
			fields := strings.Fields(line)
			for _, _ = range fields {
				wordCount++
			}
			if line == "" {
				break
			}
		}

		fmt.Println(wordCount)
	},
}

func init() {
	rootCmd.AddCommand(wordCountCmd)
}
