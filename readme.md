# Word classify

A quick fuzzy logic script that classifies words to their part of speech, given a list. NER (named entity recognition), is accomplished using Prose v2.

## Usage

Accepts a wordlist file where each word is a newline (an awk-style plain text table also works).

1. Clone this repo
2. Run with `go run . wordlist.txt` (or any suitable plain text source)

 Each word's part of speech is saved to a file in the `/output` directory. (Go must be installed).

