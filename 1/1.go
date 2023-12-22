package main

import (
	"bufio"
	"log"
	"os"
)

// searchLeft takes in a string and searches for the first digit from the left of the string and returns the digit as an int
func searchLeft(str string) int {
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}
	}
	return 0
}

// searchRight takes in a string and searches for the first digit from the right of the string and returns the digit as an int
func searchRight(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}
	}
	return 0
}

// calculateNumber will take the leftmost digit found in a string and the rightmost digit found in a string and return a concatenated int
func calculateNumber(str string) int {
	return searchLeft(str)*10 + searchRight(str)
}

// calculateSum will take a slice of ints and return the sum of all ints in the slice
func calculateSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// loadInputs loads from a txt file and returns a slice of strings, each string is a line from the txt file. The txt file is called inputs.txt
func loadInputs() []string {
	var inputs []string
	//Open inputs.txt
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Close file when function ends
	defer file.Close()

	//Read file line by line and add each line to inputs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	//Return inputs
	return inputs
}

func main() {
	inputs := loadInputs()
	var numbers []int

	//Calculate numbers from inputs
	for _, input := range inputs {
		numbers = append(numbers, calculateNumber(input))
	}

	//Finally calculate sum of numbers
	sum := calculateSum(numbers)
	log.Println("Sum of numbers:", sum)
}
