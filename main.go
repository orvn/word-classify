package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jdkato/prose/v2"
)

var posTags = map[string][]string{
	"noun":      {"NN", "NNS", "NNP", "NNPS"},
	"verb":      {"VB", "VBD", "VBG", "VBN", "VBP", "VBZ"},
	"adjective": {"JJ", "JJR", "JJS"},
	"adverb":    {"RB", "RBR", "RBS"},
}

func classifyWord(word string) string {
	doc, err := prose.NewDocument(word)
	if err != nil {
		return "error"
	}

	for _, tok := range doc.Tokens() {
		for category, tags := range posTags {
			for _, tag := range tags {
				if tok.Tag == tag {
					return category
				}
			}
		}
		return "other"
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

	os.MkdirAll("output", os.ModePerm)

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
		pos := classifyWord(word)
		var outPath string
		switch pos {
		case "noun":
			outPath = "output/nouns.txt"
		case "verb":
			outPath = "output/verbs.txt"
		case "adjective":
			outPath = "output/adjectives.txt"
		case "adverb":
			outPath = "output/adverbs.txt"
		default:
			outPath = "output/others.txt"
		}
		f, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			continue
		}
		defer f.Close()
		if _, err := f.WriteString(word + "\n"); err != nil {
			fmt.Printf("Error writing word: %s\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
