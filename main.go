package main

import (
	"fmt"
	"log"
)

func main() {
	var nProblem int

	fmt.Println("Функцию какой задачи запустить?")

	_, err := fmt.Scanln(&nProblem)
	fatalErr(err)
	if nProblem < 1 || nProblem > 4 {
		log.Fatalln("нет такой задачи")
	}

	switch nProblem {
	case 1:
		var nLaptops int
		fmt.Println("Введите количество компьютеров:")
		_, err = fmt.Scanln(&nLaptops)
		fatalErr(err)

		res, err := Laptops(nLaptops)
		fatalErr(err)

		fmt.Println(res)
	}
}

func fatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
