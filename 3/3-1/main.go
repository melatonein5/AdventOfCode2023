package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
EXAMPLE ENGINE SCHEMATIC

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
*/

// I will need to load this into a 2D array, loop through the array and check for any character that is not a period or a digit.
// If the character is not a period or a digit, I need to check adjacent characters to see if they are digits. If they are, I need too look either side to find the periods, which mark the start and end of the number.
// I will then need to add this number to a slice of ints, the sum of which will be the answer.
// NOTE THAT ADJACENT ALSO INCLUDES DIAGONAL!!!

// Engine is a 2D slice of characters
type Engine [][]rune

// LoadEngine will load the engine from a txt file and return the engine as a 2D slice of characters
func LoadEngine() Engine {
	//Create a 2D slice of characters
	var engine Engine

	//Open engine.txt
	file, err := os.Open("engine.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Close file when function ends
	defer file.Close()

	//Read file line by line split each line into a slice of runes and add each line to engine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		engine = append(engine, []rune(scanner.Text()))
	}

	//Return engine
	return engine
}

// findSpecialCharacters will take in an engine and return a slice of indexes of special characters
func findSpecialCharacters(engine Engine) [][]int {
	var specialCharacters [][]int
	for i := 0; i < len(engine); i++ {
		for j := 0; j < len(engine[i]); j++ {
			if engine[i][j] != '.' {
				//Check if engine[i][j] is not a digit
				if engine[i][j] < '0' || engine[i][j] > '9' {
					specialCharacters = append(specialCharacters, []int{i, j})
				}
			}
		}
	}
	return specialCharacters
}

// findAdjacentDigits will take in an engine and a slice of indexes of special characters, find an adjacent digit, search either side for periods, and return the number as an integer
func findAdjacentDigits(engine Engine, specialCharacters []int) int {
	//First, define the return value as an int equal to zero
	var number int = 0

	//Adjacent digits will be from [i-1, j-1] to [i+1, j+1]
	//Perform a search for digits in this range, when one is found, return the position of the digit
	var digitPosition []int
	//Search engine bewteen [i-1, j-1] and [i+1, j+1]
	for i := specialCharacters[0] - 1; i <= specialCharacters[0]+1; i++ {
		for j := specialCharacters[1] - 1; j <= specialCharacters[1]+1; j++ {
			//Check if engine[i][j] is a digit
			if engine[i][j] >= '0' && engine[i][j] <= '9' {
				digitPosition = []int{i, j}
				break
			}
		}
	}

	//Check if digitPosition is empty, if it is, return 0
	if len(digitPosition) == 0 {
		return 0
	}

	//Now we only have to search on the j axis, as we know the i axis is the same as digitPosition[1]
	//Search left for a non-digit, and record the position of the non-digit + 1 (This is the start of the number)
	var start int
	for j := digitPosition[1]; j >= 0; j-- {
		//Check if engine[digitPosition[0]][j] is a non-digit, which can include special characters
		if engine[digitPosition[0]][j] < '0' || engine[digitPosition[0]][j] > '9' {
			start = j + 1
			break
		} else if j == 0 {
			start = 0
		}
	}

	//Search right for a non-digit, and record the position of the non-digit - 1 (This is the end of the number)
	var end int
	for j := digitPosition[1]; j < len(engine[digitPosition[0]]); j++ {
		if engine[digitPosition[0]][j] < '0' || engine[digitPosition[0]][j] > '9' {
			end = j - 1
			break
		} else if j == len(engine[digitPosition[0]])-1 {
			end = len(engine[digitPosition[0]]) - 1
		}
	}

	//The number is located between engine[digitPosition[0]][start] and engine[digitPosition[0]][end], convert this to an int
	number, _ = strconv.Atoi(string(engine[digitPosition[0]][start : end+1]))

	//Return the number
	return number
}

func main() {
	//Load engine
	engine := LoadEngine()

	//Find special characters
	specialCharacters := findSpecialCharacters(engine)

	//Loop through special characters and find adjacent numbers and add them to a slice of ints
	var numbers []int
	for _, specialCharacter := range specialCharacters {
		numbers = append(numbers, findAdjacentDigits(engine, specialCharacter))
	}

	//Sum the numbers
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	//Print the sum
	fmt.Println(sum)
}
