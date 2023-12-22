package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Three variables to hold the number of maximum red, green, and blue marbles that can be taken from the bag (12, 13, 14)
var maxRed int = 12
var maxGreen int = 13
var maxBlue int = 14

// Turn is a struct that contains the how many red, green, and blue marbles taken from the bag on a turn
type Turn struct {
	red   int
	green int
	blue  int
}

// Game is a struct that contains the game ID, number of turns and a slice of turns
type Game struct {
	id    int
	turns int
	turn  []Turn
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

// parseTurn takes in a turn string and returns a Turn struct
func parseTurn(turn string) Turn {

	//split the string by comma space
	colours := strings.Split(turn, ", ")

	//create a turn struct
	var t Turn

	//Loop through colours and parse each colour into the turn struct
	for _, colour := range colours {
		//Split the colour by space
		colourSplit := strings.Split(colour, " ")
		//Switch on the second index of colourSplit, which is the colour, the first index is the number of marbles
		switch colourSplit[1] {
		case "red":
			//Parse the number of marbles, which is the first index of colourSplit, into an int and assign it to t.red
			t.red, _ = strconv.Atoi(colourSplit[0])
		case "green":
			//Parse the number of marbles, which is the first index of colourSplit, into an int and assign it to t.green
			t.green, _ = strconv.Atoi(colourSplit[0])
		case "blue":
			//Parse the number of marbles, which is the first index of colourSplit, into an int and assign it to t.blue
			t.blue, _ = strconv.Atoi(colourSplit[0])
		}
	}

	//return the turn struct
	return t
}

// parseInputs takes in a slice of strings and returns a slice of games
func parseInputs(inputs []string) []Game {
	//Create a slice of games
	var games []Game
	//Loop through inputs and parse each line into a game
	for i := 0; i < len(inputs); i++ {
		//Create a game
		var game Game
		//Here is a sample line: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		//Split the line into 2 parts, the first part is the game ID, the second part is the turns
		gameIDAndTurns := strings.Split(inputs[i], ": ")
		//Parse the game ID, which will start at index 5 to the end of the string
		game.id, _ = strconv.Atoi(gameIDAndTurns[0][5:])

		//Split the turns when there is a semicolon
		turnsStrings := strings.Split(gameIDAndTurns[1], "; ")

		//Parse the number of turns, which is the length of turnsStrings
		game.turns = len(turnsStrings)

		//Loop through turnsStrings and return a slice of turns using parseTurn
		for _, turnString := range turnsStrings {
			game.turn = append(game.turn, parseTurn(turnString))
		}

		//Append the game to games
		games = append(games, game)
	}

	return games
}

// Check games will loop through a slice of games and check any turn exceeds the maximum number of marbles that can be taken from the bag per colour. Retursn a slice of ints, each int is the game ID of a game that has a turn that does not exceed the maximum number of marbles that can be taken from the bag per colour
func checkGames(games []Game) int {
	//Create a slice of ints
	sum := 0

	//Loop through games
	for _, game := range games {
		flag := true
		//Loop through game.turn
		for _, turn := range game.turn {
			//If red, green and blue marbles taken from the bag exceeds the maximum number of marbles that can be taken from the bag per colour, flag the game as false and break the loop
			if turn.red > maxRed || turn.green > maxGreen || turn.blue > maxBlue {
				flag = false
				break
			}
		}

		//If flag is true, add the game ID to sum
		if flag {
			sum += game.id
		}
	}

	//Return gameIDs
	return sum
}

func main() {
	//Load inputs
	inputs := loadInputs()

	//Parse inputs
	games := parseInputs(inputs)

	sum := checkGames(games)

	//Print the sum
	fmt.Println(sum)
}
