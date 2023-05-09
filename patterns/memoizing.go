package patterns

import (
	"fmt"
	"playground/fp"
)

type Number interface {
	int | int64 | float64 | int32 | int16 | int8 | uint | uint8 | uint16 | uint32
}

type Memoized[T Number] func(T) T

func Memoize[T Number](mf func(T) T) func(T) T {
	cache := make(map[T]T)

	return func(i T) T {
		if val, found := cache[i]; found {
			fmt.Println("found value -> ", i, val)
			return val
		}
		result := mf(i)
		cache[i] = result
		return result
	}
}

func tryMemoize() {
	var memoizedFib = Memoize[int](fp.RecursiveFibonacci)
	value1 := memoizedFib(10)
	println(value1)
	value2 := memoizedFib(10)
	println(value2)
}
