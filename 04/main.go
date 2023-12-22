package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Card is a struct that contains the card's ID, numbers and winning numbers
type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
	Score          int
	Matches        int
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

// extractID takes in a string and returns the card ID as an int
func extractID(input string) int {
	//Regex to extract the card ID is "Card (\d+):"
	r, _ := regexp.Compile(`(\d+)`)

	//Find the first match
	match := r.FindString(input)

	//return the first submatch as an int
	id, _ := strconv.Atoi(match)
	return id
}

// extractNumbers takes in a string and returns a slice of ints
func extractNumbers(input string) []int {
	//Regex to extract the numbers is "(\d+)"
	r, _ := regexp.Compile(`(\d+)`)

	//Find all matches
	matches := r.FindAllString(input, -1)

	//Create a slice of ints
	var numbers []int

	//Loop through matches and parse each match into an int
	for _, match := range matches {
		number, _ := strconv.Atoi(match)
		numbers = append(numbers, number)
	}

	//Return numbers
	return numbers
}

// calculateScore takes in numbers and winning numbers and returns the score, which is 0 if no numbers match, and everything after that is 2^(number of matches-1)
func calculateScore(numbers []int, winningNumbers []int) (int, int) {
	//Create a counter for the number of matches
	var matches int

	//Loop through numbers and winningNumbers and compare each number
	for _, number := range numbers {
		for _, winningNumber := range winningNumbers {
			if number == winningNumber {
				matches++
			}
		}
	}

	//If there are no matches, return 0
	if matches == 0 {
		return 0, 0
	}

	//Otherwise return 2^(matches-1)
	return 1 << (matches - 1), matches
}

// parseCards takes in a slice of strings and returns a slice of Card structs
func parseCards(inputs []string) []Card {
	//Each input has the following format: "Card 187: 69  1 12 25 31 26 35 94  8 73 | 89 25 58 94 11 69 78 35 73 80 17  1 43 91 88  7 97 12 44 70 26 31 33 57 99"
	//The colon splits the card ID from the numbers
	//The pipe splits the numbers from the winning numbers
	//Create a slice of Card structs
	var cards []Card

	//Loop through inputs
	for _, input := range inputs {
		var card Card
		//First split the input by colon
		cardSplit := strings.Split(input, ":")

		//Get the card ID, which is the first index of cardSplit
		card.ID = extractID(cardSplit[0])

		//Split based on | to get the numbers and winning numbers
		numbersSplit := strings.Split(cardSplit[1], "|")

		//Get the numbers, which is the first index of numbersSplit
		card.Numbers = extractNumbers(numbersSplit[0])

		//Get the winning numbers, which is the second index of numbersSplit
		card.WinningNumbers = extractNumbers(numbersSplit[1])

		//Calculate the score
		card.Score, card.Matches = calculateScore(card.Numbers, card.WinningNumbers)

		//append the card to cards
		cards = append(cards, card)
	}
	return cards
}

// calculateTotalScore takes in a slice of cards and returns the total score
func part1(cards []Card) int {
	var totalScore int
	for _, card := range cards {
		totalScore += card.Score
	}
	return totalScore
}

// part2 means that if card 5 has 6 matches, I get a copy of cards 6, 7, 8, 9, 10, 11. Loop through the cards and calculate the total number of cards I have, including the original cards
func part2(cards []Card) int {
	//Create a map to track how many cards I have for each card ID
	cardMap := make(map[int]int)
	//intialize the map with the one of each card
	for _, card := range cards {
		cardMap[card.ID] = 1
	}

	//Loop through cards. If card 1 has 3 matches, add 1 to cardMap[2], cardMap[3], cardMap[4]. If card 2 has 4 matches, add 1 to cardMap[3], cardMap[4], cardMap[5], cardMap[6]
	for _, card := range cards {
		for i := 1; i <= card.Matches; i++ {
			//If I have 2 copies of card 2, that had 2 matches, I will have 2 copies of card 3, 4, plus the original 2 copies of 3 and 4
			for j := 1; j <= cardMap[card.ID]; j++ {
				cardMap[card.ID+i]++
			}
		}
	}

	var totalCards int

	//Loop through cardMap and calculate the total number of cards I have
	for _, card := range cards {
		totalCards += cardMap[card.ID]
	}

	//Return totalCards
	return totalCards
}

func main() {
	inputs := loadInputs()

	//Parse inputs into cards
	cards := parseCards(inputs)

	//Calculate total score
	totalScore := part1(cards)
	//Print "Part 1: totalScore"
	log.Println("Part 1:", totalScore)

	//Part 2
	totalCards := part2(cards)
	//Print "Part 2: totalCards"
	log.Println("Part 2:", totalCards)
}
