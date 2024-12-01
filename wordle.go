package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

var answers []string = getWordsFromTxt()
var gameWords []string

// Get all the words that contain 5 letters
func getWordsFromTxt() []string {
	wordsTxt := []string{}
	file, err := os.Open("./res/spanish_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if utf8.RuneCountInString(word) == 5 {
			wordsTxt = append(wordsTxt, word)
		}
	}
	file.Close()

	return wordsTxt
}

func main() {
	playGame()
}

// Play a full game (with rounds)
func playGame() {
	answer := answers[rand.Intn(len(answers)-1)]
	score := 0

	if playRound(answer) == true {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("¡Has perdido! La palabra era:\033[38;5;41m", answer, "\033[0m")
		score = 0
	}
}

// Play a round
func playRound(answer string) bool {
	guess := ""
	guessedWords := []string{}
	//guessedLetters := []string{}
	fmt.Println("Escribe una palabra de cinco letras")
	for len(guessedWords) <= 5 {
		//fmt.Println("Escribe !info para mostrar la información de la ronda acutal")
		fmt.Scan(&guess)
		/*if guess == "!info" {
			showRoundInfo(guessedLetters)
			continue
		}*/
		if utf8.RuneCountInString(guess) != 5 {
			fmt.Println("Escribe una palabra de cinco letras")
			continue
		}

		if guess == answer {
			return true
		}

		guessChars := []rune(guess)
		coloredWord := strings.Builder{}
		for i, r := range guessChars {
			containsLetter, samePos := compareChars(answer, r, i)
			coloredLetter := strings.Builder{}
			if containsLetter {
				if samePos {
					coloredLetter.WriteString("\033[38;5;46m")
				} else {
					coloredLetter.WriteString("\033[38;5;226m")
				}
			} else {
				coloredLetter.WriteString("\033[38;5;238m")
			}
			coloredLetter.WriteRune(r)
			coloredLetter.WriteString("\033[0m")
			//guessedLetters = append(guessedLetters, coloredLetter.String())
			coloredWord.WriteString(coloredLetter.String())
		}
		fmt.Println(coloredWord.String())

		guessedWords = append(guessedWords, coloredWord.String())
	}
	return false
}

// Check if the string contains the letter and if it is in the same positiion
func compareChars(answer string, letter rune, i int) (bool, bool) {
	var (
		contains    bool
		samePos     bool
		answerChars = []rune(answer)
	)
	for ai, ac := range answerChars {
		if ac == letter {
			contains = true
			if ai == i {
				samePos = true
			}
		}
	}

	return contains, samePos
}

// Show round information (letter found in the word,
// letters that are not in the word and letters that haven-t been entered
/*func showRoundInfo(guessedLetters []string) {
	validLetters := []string{}
	invalidLetters := []string{}
	for i, guess := range guessedLetters {

	}
}*/
