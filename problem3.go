package main

import "errors"

func PrimeNumbers(min, max int) ([]int, error) {
	if min > max {
		return nil, errors.New("минимум не может быть больше максимума")
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
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
