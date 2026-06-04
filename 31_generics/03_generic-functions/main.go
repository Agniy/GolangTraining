// Классические обобщённые функции работы с коллекциями: Map, Filter, Reduce.
//
// Это то, ради чего дженерики просили годами: один раз написать Map для среза
// любого типа, а не плодить копии под каждый тип элемента.
//
// Обратите внимание: Map и Reduce используют ДВА параметра типа (T и U) —
// тип входа и тип результата могут различаться.
package main

import (
	"fmt"
	"strings"
)

// Map применяет f к каждому элементу и возвращает новый срез типа []U.
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter оставляет только элементы, для которых keep вернул true.
func Filter[T any](s []T, keep func(T) bool) []T {
	var result []T
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce сворачивает срез в одно значение, накапливая результат в acc.
func Reduce[T, U any](s []T, init U, f func(acc U, cur T) U) U {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Map: int -> string
	labels := Map(nums, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println("Map:", labels)

	// Filter: только чётные
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filter (чётные):", evens)

	// Reduce: сумма
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Reduce (сумма):", sum)

	// Map работает и со строками: применяем функцию из стандартной библиотеки.
	words := Map([]string{"go", "is", "fun"}, strings.ToUpper)
	fmt.Println("Map ToUpper:", words)
}
