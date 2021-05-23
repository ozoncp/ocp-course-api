package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ozoncp/ocp-course-api/internal/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %v <filename> <count>\n", os.Args[0])
		os.Exit(1)
	}
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "The value '%v' is not a number!\n", os.Args[2])
		os.Exit(2)
	}
	utils.RepeatedlyRead(os.Args[1], count)
}
