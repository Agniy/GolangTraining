// Вывод типов (type inference) — почему обычно НЕ нужно писать [T] вручную.
//
// Компилятор сам определяет параметры типа из типов аргументов. Это делает
// вызовы обобщённых функций такими же лаконичными, как обычных.
//
// Здесь показано:
//   - когда тип выводится автоматически;
//   - когда его ПРИХОДИТСЯ указывать явно (вывести не из чего);
//   - как работает частичный вывод.
package main

import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Zero не принимает аргументов типа T — выводить тип не из чего.
func Zero[T any]() T {
	var z T
	return z
}

func main() {
	nums := []int{1, 2, 3}

	// 1) Полный автоматический вывод: T=int, U=int.
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("вывод из аргументов:", doubled)

	// 2) U выводится из типа возврата функции f (здесь string).
	strs := Map(nums, func(n int) string { return fmt.Sprint(n) })
	fmt.Println("U выведен из f:", strs)

	// 3) Явное указание обязательно: у Zero нет аргументов, из которых
	//    можно было бы вывести T.
	fmt.Printf("Zero[int]()    = %v\n", Zero[int]())
	fmt.Printf("Zero[string]() = %q\n", Zero[string]())
	fmt.Printf("Zero[bool]()   = %v\n", Zero[bool]())

	// 4) Явное указание можно использовать и там, где вывод сработал бы —
	//    иногда для читаемости.
	tripled := Map[int, int](nums, func(n int) int { return n * 3 })
	fmt.Println("явные параметры типа:", tripled)
}
