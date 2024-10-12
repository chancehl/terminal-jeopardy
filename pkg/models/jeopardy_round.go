package models

import (
	"fmt"
	"maps"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/chancehl/terminal-jeopardy/pkg/constants"
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

	fmt.Fprintln(writer, strings.Join(headers, "\t"))

	uniqueValues := getUniqueValues(values)

	for i := 0; i < len(uniqueValues); i++ {
		for j := 0; j < constants.Rounds.CategoriesPerRound; j++ {
			fmt.Fprintf(writer, "%d\t", uniqueValues[i])
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

// Prints the current round to the console (not going to lie... ChatGPT wrote this 🙏)
func (r *JeopardyRound) PrintDecoratedRoundName() {
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

func getUniqueValues(values []int) []int {
	uniqueValues := make([]int, 0)

	valueMap := make(map[int]bool)

	for _, value := range values {
		valueMap[value] = true
	}

	for key := range maps.Keys(valueMap) {
		uniqueValues = append(uniqueValues, key)
	}

	sort.Ints(uniqueValues)

	return uniqueValues
}
