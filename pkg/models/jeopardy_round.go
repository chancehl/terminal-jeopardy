package models

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type JeopardyRound struct {
	Name       string             `json:"name"`
	Categories []JeopardyCategory `json:"categories"`
}

// Prints the categories to the std out
func (r *JeopardyRound) PrintRoundCategories() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.DiscardEmptyColumns)

	var values []int
	var headers []string

	for _, category := range r.Categories {
		headers = append(headers, category.Name)

		for _, question := range category.Questions {
			values = append(values, question.Value)
		}
	}

	sort.Ints(values)

	fmt.Fprintln(writer, strings.Join(headers, "\t"))

	for _, value := range values {
		for i := 0; i < len(r.Categories); i++ {
			fmt.Fprintf(writer, "%d\t", value)
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

// Prints the current round to the console (not going to lie... ChatGPT wrote this 🙏)
func (r *JeopardyRound) PrintRound() {
	width := 50 // minimum width of 50 chars

	// Calculate padding on each side of the input string
	padding := (width - len(r.Name) - 2) / 2
	leftPadding := strings.Repeat(" ", padding)
	rightPadding := strings.Repeat(" ", width-len(r.Name)-padding-2)

	// Top and bottom decoration line
	decoration := strings.Repeat("*", width)

	// Print the decorated and centered output
	fmt.Println(decoration)
	fmt.Printf("*%s%s%s*\n", leftPadding, r.Name, rightPadding)
	fmt.Println(decoration)
}
