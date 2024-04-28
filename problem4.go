package main

import (
	"fmt"
	"strconv"
)

type Utils struct {
}

func (t *Utils) printMultTable(size int) error {
	padding := len(strconv.Itoa(size*size)) + 1

	firstLine := fmt.Sprintf("%*d", padding*2-1, 1)
	for i := 2; i <= size; i++ {
		firstLine += fmt.Sprintf("%*d", padding, i)
	}
	fmt.Println(firstLine)

	for col := 1; col <= size; col++ {
		nextLine := fmt.Sprintf("%*d", padding-1, col)

		for row := 1; row <= size; row++ {
			nextLine += fmt.Sprintf("%*d", padding, row*col)
		}

		fmt.Println(nextLine)
	}

	return nil
}
