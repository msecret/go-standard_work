package work

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type IWords interface {
	IsWord(word string) bool
}

// WordList provides a dictinary of words and functions related to the
// dictionary.
type WordList struct {
	words []string
}

// Creates and initializes a new WordList.
// Takes a list of words as an argument, or an empty list. If an empty list, it
// will use load words from a default text file.
// Returns a pointer to a new WordList.
func NewWordList(l []string) *WordList {
	if len(l) == 0 {
		l, _ = ReadWordsFromFile("words.txt")
	}
	return &WordList{l}
}

// IsWord will take a word passed in and check if its a word in the dictionary
// of words stored in the WordList.
// Returns true if its in the WordList, false if its not.
func (l WordList) IsWord(word string) bool {
	isWord := false
	for _, w := range l.words {
		if strings.ToLower(word) == w {
			isWord = true
		}
	}

	return isWord
}

// ReadWordsFromFile will read the passed in file path and return a list of
// words that were separated by new line. It's important that each word is
// separated by newline or it will no work properly.
// Will return a list of words and an error.
func ReadWordsFromFile(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0660)
	if err != nil {
		fmt.Println("File error: " + err.Error())
		return make([]string, 0), errors.New("File error: " + err.Error())
	}

	defer file.Close()

	words, err := ReadLinesString(file)

	return words, err
}

// ReadLinesString takes an os.File and returns a list of string separated by
// newlines. It will make no effort to parse the data into other types
// Returns a list and an error.
// TODO Move to something like file.go
func ReadLinesString(file *os.File) ([]string, error) {
	var err error
	words := make([]string, 0)
	reader := bufio.NewReader(file)
	err = nil
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			break
		}
		words = append(words, str)
	}

	return words, err
}
