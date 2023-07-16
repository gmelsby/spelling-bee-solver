package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	c := make(chan *Node)
	go buildTrieFromDictionary(c)
	letters := getUserLetters()
	root := <-c
	results := findWords(root, letters)
	sort.Strings(results)
	fmt.Println("\nResults:")
	for _, result := range results {
		fmt.Println(result)
	}
}

// builds a trie from the list of words in a file
// passes built tree back through channel
func buildTrieFromDictionary(c chan *Node) {
	trie := &Node{}
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) >= 4 {
			putWordIntoTrie(scanner.Text(), trie)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	c <- trie
}

// adds the passed-in word to the passed-in Trie
func putWordIntoTrie(word string, trie *Node) {
	current := trie
	for _, char := range word {
		child := current.Children[char]
		// check if child exists
		if child == nil {
			// check if map even exists, if not crate it
			if current.Children == nil {
				current.Children = make(map[rune]*Node)
			}
			child = &Node{}
			current.Children[char] = child
		}
		current = child
	}
	current.IsTerminator = true
}

// nodes to make up a Trie
type Node struct {
	Children     map[rune]*Node
	IsTerminator bool
}

// gets user input and cleans it
func getUserLetters() string {
	fmt.Println("Enter in the letters you want to solve the puzzle for in a single line.")
	fmt.Println("Make sure the key letter is the first one.")
	fmt.Print("\nLetters: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// clean up string as much as we can
	line = strings.ToLower(line)
	line = regexp.MustCompile("[^a-z]+").ReplaceAllString(line, "")
	return removeDuplicateRunes(line)
}

func removeDuplicateRunes(line string) string {
	set := make(map[rune]bool)
	result := []rune{}

	for _, char := range line {
		if !set[char] {
			result = append(result, char)
			set[char] = true
		}
	}

	return string(result)
}

func findWords(root *Node, letters string) []string {
	results := []string{}
	findWordsRecursive(root, letters, &results, &[]rune{})
	return results
}

// returns true if the rune is present in the slice, otherwise false
func runeSliceContains(slice *[]rune, target rune) bool {
	for _, char := range *slice {
		if char == target {
			return true
		}
	}
	return false
}

func findWordsRecursive(node *Node, letters string, results *[]string, combination *[]rune) {
	if node.IsTerminator {
		// check that the special letter is in the word
		targetRune, _ := utf8.DecodeRuneInString(letters)
		if runeSliceContains(combination, targetRune) {
			*results = append(*results, string(*combination))
		}
	}

	for _, char := range letters {
		if node.Children != nil && node.Children[char] != nil {
			// add the char to our current combination
			*combination = append(*combination, char)
			findWordsRecursive(node.Children[char], letters, results, combination)
			// remove the char from our current combination
			*combination = (*combination)[:len(*combination)-1]
		}
	}
}
