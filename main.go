package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	root := Node{}
	buildTrieFromDictionary(&root)
}

func buildTrieFromDictionary(trie *Node) {
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) >= 4 {
			putWordIntoTrie(scanner.Text(), trie)
			log.Print(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
