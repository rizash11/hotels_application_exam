package main

func CommonDivisors(Numbers []int) ([]int, error) {
	var res []int
	min := Numbers[0]
	for _, num := range Numbers {
		if num < min {
			min = num
		}
	}

outer:
	for i := 2; i < min; i++ {
		for _, num := range Numbers {
			if num%i != 0 {
				continue outer
			}
		}

		res = append(res, i)
	}

	return res, nil
}
