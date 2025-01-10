package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DailyRevenue struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

type RevenueByState struct {
	State   string  `json:"state"`
	Revenue float64 `json:"revenue"`
}

func main() {

	// Question #1 - Index SUM
	indexK()
	// End of Question #1 - Index SUM

	// Question #2 - Fibonacci
	isFibonacci()
	// End of Question #2 - Fibonacci

	// Question #3 - Daily Revenue
	dailyRevenue()
	// End of Question #3 - Daily Revenue

	// Question #4 - Revenue percentage representation by state.
	revenuePercentage()
	// End of Question #4 - Revenue percentage representation by state.

	// Question #5 - Reverse String
	revertString("Hello, World!")
	// End of Question #5 - Reverse String
}

// indexK calculates the sum of the first n natural numbers running on an index set on 13.
func indexK() {
	println("#1 - Index SUM")
	var index, k, sum int
	index = 13
	sum = 0
	k = 0

	for k < index {
		k = k + 1
		sum = sum + k
		println("Result of the sum: ", sum)
	}
}

// isFibonacci checks if a number is a Fibonacci number or not.
func isFibonacci() {
	println("#2 - Fibonacci")
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	a, b := 0, 1
	for b <= n {
		if b == n {
			println(n, "is a Fibonacci number.")
		}
		a, b = b, a+b
	}
	println(n, "is not a Fibonacci number.")
}

// dailyRevenue reads a JSON file containing daily revenues and calculates the minimum, maximum, and average revenue.
func dailyRevenue() {
	println("#3 - Daily Revenue")
	file, err := os.Open("dados.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}(file)

	byteValue, _ := ioutil.ReadAll(file)

	var dailyRevenues []DailyRevenue
	err = json.Unmarshal(byteValue, &dailyRevenues)
	if err != nil {
		log.Fatalf("failed to parse JSON: %s", err)
	}

	var validRevenues []float64
	for _, r := range dailyRevenues {
		if r.Valor > 0 {
			validRevenues = append(validRevenues, r.Valor)
		}
	}

	if len(validRevenues) == 0 {
		fmt.Println("No valid revenue found.")
		return
	}

	// Calculate min, max, and average revenue
	minRevenue := validRevenues[0]
	maxRevenue := validRevenues[0]
	totalRevenue := 0.0

	for _, revenue := range validRevenues {
		if revenue < minRevenue {
			minRevenue = revenue
		}
		if revenue > maxRevenue {
			maxRevenue = revenue
		}
		totalRevenue += revenue
	}

	averageRevenue := totalRevenue / float64(len(validRevenues))

	// Count days with revenue above average
	daysAboveAverage := 0
	for _, revenue := range validRevenues {
		if revenue > averageRevenue {
			daysAboveAverage++
		}
	}

	// Print results
	fmt.Printf("Minimum revenue: %.2f\n", minRevenue)
	fmt.Printf("Maximum revenue: %.2f\n", maxRevenue)
	fmt.Printf("Number of days with revenue above average: %d\n", daysAboveAverage)
}

// revenuePercentage reads a JSON file containing revenue by state and calculates the revenue percentage.
func revenuePercentage() {
	println("#4 - Revenue Percentage Representation by State")
	// Read JSON file
	file, err := os.Open("revenue-by-state.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}(file)

	byteValue, _ := ioutil.ReadAll(file)
	var revenueByStates []RevenueByState
	err = json.Unmarshal(byteValue, &revenueByStates)
	if err != nil {
		log.Fatalf("failed to parse JSON: %s", err)
	}
	copyValues := make([]RevenueByState, len(revenueByStates))
	copy(copyValues, revenueByStates)

	// Calculate total revenue
	totalRevenue := 0.0
	for _, r := range revenueByStates {
		totalRevenue += r.Revenue
	}

	// Calculate revenue percentage
	for i, r := range revenueByStates {
		revenueByStates[i].Revenue = (r.Revenue / totalRevenue) * 100
	}

	// Print results
	for i, r := range revenueByStates {
		fmt.Printf("%s: R$ %.2f ≅ %.2f%%\n", r.State, copyValues[i].Revenue, r.Revenue)
	}
}

// revertString reverses a given string and prints the result.
func revertString(s string) {
	println("#5 - Reverse String")
	println("Original string:", s)
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	println("Reversed string:", result)
}
