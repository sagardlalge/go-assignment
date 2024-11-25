package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"strconv"
	"log"
	"strings"
)

// Function to validate input
func validateInput(value float64) error {
	if value <= 0 {
		return errors.New("input must be a positive number greater than zero")
	}
	return nil
}

// Function to calculate profit
func calculateProfit(costPrice, sellingPrice float64) (float64, error) {
	// Validate inputs
	if err := validateInput(costPrice); err != nil {
		return 0, err
	}
	if err := validateInput(sellingPrice); err != nil {
		return 0, err
	}

	// Calculate profit
	profit := sellingPrice - costPrice
	return profit, nil
}

// Function to write profit to file
func writeProfitToFile(profit float64) error {
	// Format the profit to a string
	profitStr := fmt.Sprintf("Profit: %.2f", profit)

	// Open or create file
	file, err := os.Create("profit.txt")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString(profitStr)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// Function to read profit from the file
func readProfitFromFile() (float64, error) {
	// Open file
	file, err := os.Open("profit.txt")
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Use bufio.Scanner to read the full line from the file
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("error reading from file: %w", scanner.Err())
	}

	// Read the content
	content := scanner.Text()

	// Trim any leading or trailing spaces
	content = strings.TrimSpace(content)

	// Ensure content starts with "Profit: "
	if !strings.HasPrefix(content, "Profit: ") {
		return 0, fmt.Errorf("unexpected content format in file: %s", content)
	}

	// Extract profit value from the string and parse it as a float
	profitStr := content[len("Profit: "):]
	profit, err := strconv.ParseFloat(profitStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing profit value from file: %w", err)
