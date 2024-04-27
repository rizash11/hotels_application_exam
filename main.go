package main

import (
	"fmt"
	"log"
)

func main() {
	var nProblem int

	fmt.Println("Функцию какой задачи запустить?")

	_, err := fmt.Scanln(&nProblem)
	checkErr(err)
	if nProblem < 1 || nProblem > 4 {
		log.Fatalln("нет такой задачи")
	}

	switch nProblem {
	case 1:
		var nLaptops int
		fmt.Println("Введите количество компьютеров:")
		_, err = fmt.Scanln(&nLaptops)
		checkErr(err)

		res, err := Laptops(nLaptops)
		checkErr(err)

		fmt.Println(res)
	case 2:
		var nNumbers []int
		var nNumber int
		fmt.Println("Введите числа, делители которых необходимо найти. Чтобы остановить ввод чисел, введите любой символ не являющийся цифрой.")
		for {
			_, err := fmt.Scan(&nNumber)
			if err != nil {
				break
			}
			nNumbers = append(nNumbers, nNumber)
		}

		res, err := CommonDivisors(nNumbers)
		checkErr(err)
		fmt.Println("Общие делители для введенных чисел:")
		fmt.Println(res)
	case 3:
		fmt.Println("Введите минимум и максимум, между которыми нужно найти простые числа.")
		var min, max int
		_, err := fmt.Scanln(&min, &max)
		checkErr(err)

		res, err := PrimeNumbers(min, max)
		checkErr(err)

		fmt.Println("Простые числа между минимумом и максимумом:")
		fmt.Println(res)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
