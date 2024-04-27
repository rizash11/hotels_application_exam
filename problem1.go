package main

import (
	"errors"
	"strconv"
)

func Laptops(nLaptops int) (string, error) {
	lastNum := nLaptops % 10

	switch {
	case lastNum == 1:
		return strconv.Itoa(nLaptops) + " компьютер", nil
	case lastNum >= 2 && lastNum <= 4:
		return strconv.Itoa(nLaptops) + " компьютера", nil
	case lastNum == 0 || (lastNum >= 5 && lastNum <= 9):
		return strconv.Itoa(nLaptops) + " компьютеров", nil
	}

	return "", errors.New("unexpected error")
}
