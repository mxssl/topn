package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// topNumbers contains how much sorted numbers user want to print
var topNumbers string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Find top N numbers in file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := findTop(fileName, topNumbers)
		fmt.Println(result[1:])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&fileName,
		"file",
		"f",
		"numbers.txt",
		"File with numbers")

	runCmd.Flags().StringVarP(&topNumbers,
		"top",
		"t",
		"10",
		"Show top N numbers")
}

func findTop(fileName string, topNumbers string) []int {
	// Convert topNumbers to int
	n, err := strconv.Atoi(topNumbers)
	if err != nil {
		fmt.Printf("Please enter correct integer number. Error: %v\n", err)
		os.Exit(1)
	}

	// Open file with numbers
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Cannot open file. Error: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	// Slice for top N ints. Length + 1 for new numbers from file
	top := make([]int, n+1)

	// Scan file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// If line is empty just continue loop
		if text == "" {
			continue
		}

		// Convert string to int
		i, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("Cannot convert string to integer. Error: %v\n", err)
		}

		// Add new number to slice on the 0 position
		top[0] = i

		// Slice sorting. Default sorting algo for int is QuickSort
		sort.Ints(top)
	}

	// Last slice sorting for large to small output
	sort.Sort(sort.Reverse(sort.IntSlice(top[1:])))

	// Return sorted slice
	return top
}
