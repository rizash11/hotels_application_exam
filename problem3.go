package main

import "errors"

func PrimeNumbers(min, max int) ([]int, error) {
	switch {
	case min > max:
		return nil, errors.New("минимум не может быть больше максимума")
	case max < 0:
		return nil, errors.New("простое число не может быть отрицательным")
	}

	var res []int

	for i := min; i <= max; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}

	return res, nil
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
