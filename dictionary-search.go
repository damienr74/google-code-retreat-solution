package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/**
 * Dawson Wilson Conceptualised the use of bits as the presence of a letter in
 * a word. This makes for a simple, yet efficient implementation.
 */
type word struct {
	str string
	bit uint
}

func createMask(str string) uint {
	mask := 0

	for i := 0; i < len(str); i++ {
		offset := asciiVal(str[i])
		mask |= 1 << offset
	}
	return uint(mask)
}

func asciiVal(char byte) uint {
	if char >= 65 && char < 90 {
		return uint(char - 64)
	}

	if char > 96 || char < 123 {
		return uint(char - 96)
	}

	return 0
}

func initDict() []word {
	file, err := os.Open("dict")

	if err != nil {
		log.Fatal(err)
	}

	wordList := []word{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		wordList = append(wordList, word{line, createMask(line)})
	}

	return wordList
}

func getLicences() []string {
	reg, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	licenceList := []string{}

	for scanner.Scan() {
		licence := strings.ToLower(scanner.Text())
		filter := reg.ReplaceAllString(licence, "")
		licenceList = append(licenceList, filter)
	}

	return licenceList
}

func maskMatch(wordList []word, searchMask uint) []string {
	matchList := []string{}

	for _, element := range wordList {
		if element.bit&searchMask == searchMask {
			matchList = append(matchList, element.str)
		}
	}

	return matchList
}

func shortestWord(words []string) string {
	if len(words) < 1 {
		return ""
	}

	shortestWord := words[0]

	for _, element := range words {
		if len(shortestWord) > len(element) {
			shortestWord = element
		}
	}

	return shortestWord
}

func main() {
	wordList := initDict()
	licenceList := getLicences()

	for _, element := range licenceList {
		matches := maskMatch(wordList, createMask(element))
		fmt.Printf("%s matches: %s\n", element, shortestWord(matches))
	}
}
