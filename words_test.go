package work

import (
	"reflect"
	"testing"
)

var testWords = []string{"leagues", "under", "the", "sea"}

func newWordListFactory() *WordList {
	return NewWordList(testWords)
}

func TestWordListIsWord_WithWord(t *testing.T) {
	testWordList := newWordListFactory()
	actual := testWordList.IsWord("under")

	if !actual {
		t.Errorf("Will return true when given an actual word")
	}
}

func TestWordListIsWord_WithNonWord(t *testing.T) {
	testWordList := newWordListFactory()
	actual := testWordList.IsWord("alsdfkj")

	if actual {
		t.Errorf("Will return false when given an non word")
	}
}

func TestWordListIsWord_WithEmptyString(t *testing.T) {
	testWordList := newWordListFactory()
	actual := testWordList.IsWord("")

	if actual {
		t.Errorf("Will return false when given an empty string")
	}

}

func TestNewWordList_WhenPassedStrings(t *testing.T) {
	testWordList := NewWordList(testWords)

	expected := testWords
	actual := testWordList.words

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`Will return a new WordList with the words attribute set to the 
      slice of words passed in\n
      expected: %v\n
      actual: %v\n`, expected, actual)
	}
}

// TODO how to test NewWordList without string? How to mock file open call, etc.
