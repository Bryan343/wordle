package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"math/rand"
	"os"
	"unicode"
	"unicode/utf8"
)

var answers []string = getWordsFromTxt()

func main() {
	word := answers[rand.Intn(len(answers) - 1)]
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	wordWithoutAccents, _, _ := transform.String(t, word)
	fmt.Println("ANSWER: ", word)
	fmt.Println("ANSWER WITHOUT ACCENTS: ", wordWithoutAccents)
	score := 0

	if playRound(wordWithoutAccents) == true {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("ANSWER: ", word)
		score = 0
	}
}

func getWordsFromTxt() []string {
	var wordsTxt []string
	file, err := os.Open("./res/spanish_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := scanner.Text()
		if utf8.RuneCountInString(test) == 5 {
			wordsTxt = append(wordsTxt, test)
			fmt.Println(test)	
		}
	}
	file.Close()

	return wordsTxt
}

func playRound(roundWord string) bool {
	var answer string 
	var answers []string

	fmt.Println("Escribe una palabra de cinco letras")
	for len(answers) <= 5  {
		fmt.Scan(&answer)
		if utf8.RuneCountInString(answer) != 5 {
			fmt.Println("Escribe una palabra de cinco letras")
			continue
		}

		if answer == roundWord {
			return true
		}

		answers = append(answers, answer)
	}
	return false
}
