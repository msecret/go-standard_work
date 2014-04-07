package work

import (
	"reflect"
	"testing"
)

var words []string = []string{"word", "words", "a", "tree", "submarine",
	"lion", "very", "long", "sentence", "about", "and", "mouth"}

func TestCeasarCipherBuildEncoder_ShiftOfOne(t *testing.T) {
	in := 1
	expected := map[string]string{" ": "A", "A": "B", "B": "C", "C": "D", "D": "E", "E": "F", "F": "G", "G": "H", "H": "I", "I": "J", "J": "K", "K": "L",
		"L": "M", "M": "N", "N": "O", "O": "P", "P": "Q", "Q": "R", "R": "S", "S": "T", "T": "U", "U": "V", "V": "W", "W": "X", "X": "Y", "Y": "Z", "Z": "a",
		"a": "b", "b": "c", "c": "d", "d": "e", "e": "f", "f": "g", "g": "h", "h": "i", "i": "j", "j": "k", "k": "l", "l": "m", "m": "n", "n": "o", "o": "p",
		"p": "q", "q": "r", "r": "s", "s": "t", "t": "u", "u": "v", "v": "w", "w": "x", "x": "y", "y": "z", "z": " "}
	actual, _ := CeasarCipherBuildEncoder(in)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`Encoder will return a map of all characters unshifted by %v, 
      expected: %v\n
      actual: %v\n 
    `, in, expected, actual)
	}
}

func TestCeasarCipherBuildDecoder_DeShiftOne(t *testing.T) {
	in := 1
	expected := map[string]string{" ": "z", "A": " ", "B": "A", "C": "B", "D": "C", "E": "D", "F": "E", "G": "F", "H": "G", "I": "H", "J": "I", "K": "J", "L": "K", "M": "L", "N": "M", "O": "N", "P": "O", "Q": "P", "R": "Q", "S": "R", "T": "S", "U": "T", "V": "U", "W": "V", "X": "W", "Y": "X", "Z": "Y", "a": "Z", "b": "a", "c": "b", "d": "c", "e": "d", "f": "e", "g": "f", "h": "g", "i": "h", "j": "i", "k": "j", "l": "k", "m": "l", "n": "m", "o": "n", "p": "o", "q": "p", "r": "q", "s": "r", "t": "s", "u": "t", "v": "u", "w": "v", "x": "w", "y": "x", "z": "y"}
	actual, _ := CeasarCipherBuildDecoder(in)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`Encoder will return a map of all characters unshifted by %v, 
      expected: %v\n
      actual: %v\n 
    `, in, expected, actual)
	}
}

// TODO test that encoder calls shfitCheck

func TestCeasarCipherApplyShift_WithLowerStringAndShiftOne(t *testing.T) {
	testShift := 1
	testString := "atest"

	expected := "buftu"
	actual, _ := CeasarCipherApplyShift(testString, testShift)

	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShift_WithMixedStringAndShiftOne(t *testing.T) {
	testShift := 1
	testString := "ATestOne"

	expected := "BUftuPof"

	actual, _ := CeasarCipherApplyShift(testString, testShift)
	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShift_WithSpacesStringAndShiftOne(t *testing.T) {
	testShift := 1
	testString := "A Test One"

	expected := "BAUftuAPof"

	actual, _ := CeasarCipherApplyShift(testString, testShift)
	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShift_WithEndingStringAndShiftFive(t *testing.T) {
	testShift := 5
	testString := "A TestY Onez"

	expected := "FEYjxydETsjD"

	actual, _ := CeasarCipherApplyShift(testString, testShift)
	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShifts_withSingleShift(t *testing.T) {
	testString := "word"
	testShift := make([][]int, 1)
	testShift[0] = []int{0, 1}

	expected := "xpse"

	actual, _ := CeasarCipherApplyShifts(testString, testShift)

	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShifts_withTwoShifts(t *testing.T) {
	testString := "word"
	testShift := make([][]int, 2)
	testShift[0] = []int{0, 1}
	testShift[1] = []int{1, 1}

	expected := "xqtf"

	actual, _ := CeasarCipherApplyShifts(testString, testShift)

	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShifts_withTwoShiftsVaryLengths(t *testing.T) {
	testString := "word"
	testShift := make([][]int, 2)
	testShift[0] = []int{0, 6}
	testShift[1] = []int{1, 2}

	expected := "Bwzl"

	actual, _ := CeasarCipherApplyShifts(testString, testShift)

	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShifts_withMultipleShifts(t *testing.T) {
	testString := "longer word"
	testShift := make([][]int, 3)
	testShift[0] = []int{0, 2}
	testShift[1] = []int{1, 1}
	testShift[2] = []int{3, 3}

	expected := "nrqmkxFBuxj"

	actual, _ := CeasarCipherApplyShifts(testString, testShift)

	if actual != expected {
		t.Errorf(`The apply shift did not apply the shift of %v correctly,\n
      expected: %v\n
      actual: %v\n
      `, testShift, expected, actual)
	}
}

func TestCeasarCipherApplyShifts_withEmptyShifts(t *testing.T) {

}
func TestcheckShift_TooLargeShift(t *testing.T) {
	in := len(Alphabet) + 1

	actualErr := checkShift(in)

	if actualErr == nil {
		t.Errorf(`Accepted a shift value over the alphabet length and did  not return
      an error`)
	}
}

func TestcheckShift_NegativeShift(t *testing.T) {
	in := -1

	actualErr := checkShift(in)

	if actualErr == nil {
		t.Errorf("Accepted a negative shift value and did  not return an error")
	}
}

func TestCeasarCipherFindShift_SmallWordShiftOne(t *testing.T) {
	testWordList := NewWordList(words)
	plainText := "word"
	expected := 1

	codedText, _ := CeasarCipherApplyShift(plainText, expected)

	actual, _ := CeasarCipherFindShift(codedText, testWordList)

	if actual != expected {
		t.Errorf(`Returns the correct shift one
      expected: %v\n
      actual: %v\n`, expected, actual)
	}
}

func TestCeasarCipherFindShift_SmallWordNoShift(t *testing.T) {
	testWordList := NewWordList(words)
	plainText := "word"
	expected := 0

	codedText, _ := CeasarCipherApplyShift(plainText, expected)

	actual, _ := CeasarCipherFindShift(codedText, testWordList)

	if actual != expected {
		t.Errorf(`Returns the correct zero shift
      expected: %v\n
      actual: %v\n`, expected, actual)
	}
}

func TestCeasarCipherFindShift_SmallWordLargeShift(t *testing.T) {
	testWordList := NewWordList(words)
	plainText := "word"
	expected := 4

	codedText, _ := CeasarCipherApplyShift(plainText, expected)

	actual, _ := CeasarCipherFindShift(codedText, testWordList)

	if actual != expected {
		t.Errorf(`Returns the correct shift of 3
      expected: %v\n
      actual: %v\n`, expected, actual)
	}
}

func TestCeasarCipherFindShift_SentenceNormalShift(t *testing.T) {
	testWordList := NewWordList(words)
	plainText := "a very long sentence about words and mouth"
	expected := 6

	codedText, _ := CeasarCipherApplyShift(plainText, expected)

	actual, _ := CeasarCipherFindShift(codedText, testWordList)

	if actual != expected {
		t.Errorf(`Returns the correct shift of 6
      expected: %v\n
      actual: %v\n`, expected, actual)
	}
}
