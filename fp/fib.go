package fp

func ImperativeFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		result := 0
		first := 1
		second := 1
		for i := 2; i < n; i++ {
			result = first + second
			first = second
			second = result
		}
		return result
	}
}

func RecursiveFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
	}
}

func FunctionalFibonacci(n int) int {
	return tailRecFibonacci(n, 0, 1, 1, 0)
}

/*
- RecursiveFibonacci(10, 1, 0, 0, 1)
- RecursiveFibonacci(10, 2, 1, 1, 1)
- RecursiveFibonacci(10, 3, 1, 1, 2)
- RecursiveFibonacci(10, 4, 2, 2, 3)
- RecursiveFibonacci(10, 5, 3, 3, 5)
- RecursiveFibonacci(10, 6, 5, 5, 8)
- RecursiveFibonacci(10, 7, 8, 8, 13)
- RecursiveFibonacci(10, 8, 13, 13, 21)
- RecursiveFibonacci(10, 9, 21, 21, 34)
- RecursiveFibonacci(10, 10, 34, 34, 55)
*/
func tailRecFibonacci(n int, counter int, first int, second int, acc int) int {
	if n == 0 {
		return 0
	} else if n <= 1 {
		return 1
	} else if counter == n {
		return acc
	} else {
		first = second
		second = acc
		// fmt.Printf(" - RecursiveFibonacci(%d, %d, %d, %d, %d)\n", n, counter+1, second, acc, first+second)
		return tailRecFibonacci(n, counter+1, first, second, first+second)
	}
}
