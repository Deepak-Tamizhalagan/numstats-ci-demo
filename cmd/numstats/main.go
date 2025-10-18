package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Deepak-Tamilzhagan/numstats-ci-demo/internal/stats"
)

func usage() {
	fmt.Println("Usage: numstats <num1> <num2> ...")
}

func parseArgs(args []string) ([]float64, error) {
	out := make([]float64, 0, len(args))
	for _, a := range args {
		v, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %q", a)
		}
		out = append(out, v)
	}
	return out, nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	nums, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Printf("Count: %d\n", len(nums))
	fmt.Printf("Sum: %g\n", stats.Sum(nums))
	fmt.Printf("Mean: %g\n", stats.Mean(nums))
	fmt.Printf("Median: %g\n", stats.Median(nums))
	fmt.Printf("Min: %g\n", stats.Min(nums))
	fmt.Printf("Max: %g\n", stats.Max(nums))
}
