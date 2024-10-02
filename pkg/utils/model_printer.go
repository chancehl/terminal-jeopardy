package utils

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/chancehl/terminal-jeopardy/pkg/constants"
	"github.com/chancehl/terminal-jeopardy/pkg/models"
)

// Prints the categories to the std out
func PrintCategories(round models.JeopardyRound) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.DiscardEmptyColumns)

	var headers []string
	for _, category := range round.Categories {
		headers = append(headers, category.Name)
	}

	header := strings.Join(headers, "\t")

	fmt.Fprintln(writer, header)

	var values []int

	if round.Name == constants.Rounds.Jeopardy {
		values = []int{200, 400, 600, 800, 1000}
	} else if round.Name == constants.Rounds.DoubleJeopardy {
		values = []int{400, 800, 1200, 1600, 2000}
	} else {
		values = []int{-1} // TODO: implement wager
	}

	for _, value := range values {
		for i := 0; i < len(round.Categories); i++ {
			fmt.Fprintf(writer, "%d\t", value)
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

// Prints the current round to the console (not going to lie... ChatGPT wrote this 🙏)
func PrintRound(round models.JeopardyRound) {
	width := 50 // minimum width of 50 chars

	// Calculate padding on each side of the input string
	padding := (width - len(round.Name) - 2) / 2
	leftPadding := strings.Repeat(" ", padding)
	rightPadding := strings.Repeat(" ", width-len(round.Name)-padding-2)

	// Top and bottom decoration line
	decoration := strings.Repeat("*", width)

	// Print the decorated and centered output
	fmt.Println(decoration)
	fmt.Printf("*%s%s%s*\n", leftPadding, round.Name, rightPadding)
	fmt.Println(decoration)
}
