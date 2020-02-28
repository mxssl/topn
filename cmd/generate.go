package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// numberLines contains how much lines with random numbers will be generated
var numberLines string

// fileName contains name of generated file
var fileName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate file with random numbers",
	Long:  `Generate file with random numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		generateFile(numberLines, fileName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&numberLines,
		"lines",
		"l",
		"1000",
		"How much lines with random numbers will be generated")

	generateCmd.Flags().StringVarP(&fileName,
		"name",
		"n",
		"numbers.txt",
		"Name of generated file")
}

func generateFile(numberLines string, fileName string) {
	// Convert numberLines to int
	lines, err := strconv.Atoi(numberLines)
	if err != nil {
		fmt.Printf("Please enter correct integer number of lines. Error: %v\n", err)
		os.Exit(1)
	}

	// Create empty file
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Cannot create a file. Error: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	// Write N lines of random numbers
	w := bufio.NewWriter(f)
	for i := 0; i < lines; i++ {
		str := strconv.Itoa(rand.Int())
		_, err := w.WriteString(str + "\n")
		if err != nil {
			fmt.Printf("Cannot write to file. Error: %v\n", err)
			os.Exit(1)
		}
	}

	w.Flush()

	fmt.Printf("File %v with %v lines of random numbers was successfully generated!\n", fileName, numberLines)
}
