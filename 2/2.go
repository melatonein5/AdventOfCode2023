package main

import (
	"bufio"
	"log"
	"os"
)

// alphaDigits is a slice of strings containing the digits 0-9 spelled out, in lowercase
var alphaDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// alphaDigitsMap maps the spelled out digits to their integer value
var alphaDigitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// searchLeft takes in a string and searches for the first digit from the left of the string and returns the digit as an int. If the current index is not a digit, it will check if it is a spelled out digit, if it is, it will return the digit as an int
func searchLeft(str string) int {
	for i := 0; i < len(str); i++ {
		//If current index is a digit, return the digit as an int
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}
		//Loop through alphaDigits and check if current index, plus the next len(alphaDigit) indexes, is equal to the current alphaDigit, if it is, return the the integer value of the current index
		for _, alphaDigit := range alphaDigits {
			if i+len(alphaDigit) <= len(str) && str[i:i+len(alphaDigit)] == alphaDigit {
				return alphaDigitsMap[alphaDigit]
			}
		}
	}
	return 0
}

// searchRight takes in a string and searches for the first digit from the right of the string and returns the digit as an int. If the current index is not a digit, it will check if it is a spelled out digit, if it is, it will return the digit as an int
func searchRight(str string) int {
	//If current index is a digit, return the digit as an int
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}
		//Loop through alphaDigits and check if current index, plus the next len(alphaDigit) indexes, is equal to the current alphaDigit, if it is, return the the integer value of the current index. This is the same as searchLeft, but in reverse
		for _, alphaDigit := range alphaDigits {
			if i+len(alphaDigit) <= len(str) && str[i:i+len(alphaDigit)] == alphaDigit {
				return alphaDigitsMap[alphaDigit]
			}
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
