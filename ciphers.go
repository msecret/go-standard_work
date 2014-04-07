package work

import (
	"errors"
	"strings"
)

const Alphabet = " ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func buildCoder(shift int) map[string]string {
	coder := make(map[string]string)

	for i, l := range Alphabet {
		var s string
		shifted := i + shift

		if shifted < 0 {
			s = string(Alphabet[len(Alphabet)+shifted])
		} else if shifted > len(Alphabet)-1 {
			s = string(Alphabet[shifted-len(Alphabet)])
		} else {
			s = string(Alphabet[i+shift])
		}

		coder[string(l)] = s
	}

	return coder
}

func checkShift(shift int) error {
	if shift < 0 || shift > 27 {
		err := errors.New("Shift too small or large")
		return err
	}

	return nil
}

// CeasarCipherBuildEncoder employs a ceasar shift cipher by shifting all
// characters in lower/uppercase English alphabet.
// Takes a shift number between 0 and 27.
// Returns a map of original letter to the shifted result:
//   {'A': 'B' ... }
func CeasarCipherBuildEncoder(shift int) (map[string]string, error) {
	var encoder map[string]string
	err := checkShift(shift)

	if err == nil {
		encoder = buildCoder(shift)
	}

	return encoder, err
}

// CeasarCipherBuildDecoder employs a ceasar shift decoder by undoing a shift
// to all characters in lower/upper English alphabet.
// Takes a shift number between 0 and 27.
// Returns a map of shifted letter to the original result:
//   {'A': 'B' ... }
func CeasarCipherBuildDecoder(shift int) (map[string]string, error) {
	var encoder map[string]string
	err := checkShift(shift)

	if err == nil {
		encoder = buildCoder(shift * -1)
	}

	return encoder, err
}

// CeasarCipherApplyShift will apply a Ceasar shift cipher to text.
// Takes a variable lenght string as the first argument, and a shift int as the
// second argument to dertmine how much to shift by.
// Returns a string that has been shifted by amount shift.
func CeasarCipherApplyShift(text string, shift int) (string, error) {
	encoder, err := CeasarCipherBuildEncoder(shift)
	if err != nil {
		return "", err
	}

	var toReturn string
	for _, l := range text {
		r := encoder[string(l)]
		toReturn += r
	}

	return toReturn, nil
}

// CeasarCipherApplyShifts will apply a squence of multiple shifts to a string
// Takes a slice of a slice of ints, each slice will have the first number as
// the position in the string to slice, the second number is the number of
// slices to shift.
func CeasarCipherApplyShifts(text string, shifts [][]int) (string, error) {
	var coded string = text
	for _, shift := range shifts {
		fragment := coded[shift[0]:]
		codedFragment, err := CeasarCipherApplyShift(fragment, shift[1])
		if err != nil {
			return "", err
		}
		coded = strings.Join([]string{coded[:shift[0]], codedFragment}, "")
	}

	return coded, nil
}

// CeasarCipherApplyShift will apply a Ceasar shift cipher to text.
// Takes a variable lenght string as the first argument, and a shift int as the
// second argument to dertmine how much to shift by.
// Returns a string that has been shifted by amount shift.
func CeasarCipherApplyDecoder(text string, shift int) (string, error) {
	encoder, err := CeasarCipherBuildDecoder(shift)
	if err != nil {
		return "", err
	}

	var toReturn string
	for _, l := range text {
		r := encoder[string(l)]
		toReturn += r
	}

	return toReturn, nil
}

// CeasarCipherFindShift will find the ceasar shift which best matches the
// encoded text passed in.
func CeasarCipherFindShift(text string, wordList IWords) (int, error) {
	highestWordCount := 0
	answer := 0
	for i := 0; i < 27; i++ {
		decoded, err := CeasarCipherApplyDecoder(text, i)
		if err != nil {
			return 0, err
		}

		sep := strings.Split(decoded, " ")
		wordCount := 0
		for _, word := range sep {
			if wordList.IsWord(word) {
				wordCount += 1
			}
		}

		if wordCount == len(sep) {
			return i, nil
		} else if wordCount > highestWordCount {
			answer = i
			highestWordCount = wordCount
		}
	}
	return answer, nil
}

func findShiftsR(text string, wordList IWords, start int) [][]int {
	var (
		s      string
		shifts [][]int
	)
	for i := 0; i < 27; i++ {
		shiftedString, _ := CeasarCipherApplyShift(text[start:], i)
		s = text[:start] + shiftedString
		for j := start; j < len(text); j++ {
			if string(s[j]) == " " && wordList.IsWord(s[j+1:]) {
				shift := findShiftsR(s, wordList, j)
				if shift != nil {
					shifts = append(shifts, []int{j, i})
				}
			} else if wordList.IsWord(s[j:]) {
				tmp := make([][]int, 1)
				tmp[0] = []int{j, i}
				return tmp
			}
		}
	}

	return shifts
}

// CeasarCipherFindBestShifts finds a slice of slices of shifts that were
// applied to the past in text.
func CeasarCipherFindBestShifts(text string, wordList IWords) ([][]int, error) {

	shifts := findShiftsR(text, wordList, 0)

	return shifts, nil
}
