package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jdkato/prose/v2"
)

func classifyWord(word string) string {
	doc, err := prose.NewDocument(word)
	if err != nil {
		return "error"
	}

	for _, tok := range doc.Tokens() {
		switch tok.Tag {
		case "NN", "NNS", "NNP", "NNPS":
			return "noun"
		case "VB", "VBD", "VBG", "VBN", "VBP", "VBZ":
			return "verb"
		case "JJ", "JJR", "JJS":
			return "adjective"
		default:
			return "other"
		}
	}
	return "unknown"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Include a file reference to use")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, "\t")
		var word string
		switch len(parts) {
		case 1:
			word = parts[0]
		case 2:
			word = parts[1]
		default:
			continue
		}
		fmt.Printf("%s -> %s\n", word, classifyWord(word))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
