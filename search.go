package work

import (
	"errors"
)

func SearchBrute(list []float64, answer float64) (int, error) {
	for i, j := range list {
		if j == answer {
			return i, nil
		}
	}

	return 0.0, errors.New("Could not find answer")
}

func SearchBiSection(list []float64, answer float64) (int, error) {
	var h, l, c int
	var a float64

	l = 0
	h = len(list) - 1
	c = (h + l) / 2
	a = list[c]

	for h-l > 1 {
		if list[h] == answer {
			return h, nil
		}
		if list[l] == answer {
			return l, nil
		}
		if a > answer {
			h = c
		} else if a < answer {
			l = c
		} else {
			return c, nil
		}
		c = (h + l) / 2
		a = list[c]
	}

	return 0, errors.New("Could not find element")
}
