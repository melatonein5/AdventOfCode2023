package boilerplate

import (
	"bufio"
	"log"
	"os"
)

//Contains the boilerplate code for the solutions, copy and paste this into the solution files
//At some point it should probably be its own package, but I'm not sure how to do that yet

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

// calculateSum will take a slice of ints and return the sum of all ints in the slice
func calculateSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
