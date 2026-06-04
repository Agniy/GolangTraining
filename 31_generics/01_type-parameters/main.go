// Дженерики (обобщённое программирование), Go 1.18 (2022) — крупнейшее
// дополнение языка со времён его выхода.
//
// Идея: писать функции и типы, работающие с РАЗНЫМИ типами, но при этом
// сохраняя проверку типов на этапе компиляции (в отличие от interface{}/any,
// где приходится делать приведение типов в рантайме).
//
// Синтаксис ПАРАМЕТРОВ ТИПА — в квадратных скобках после имени функции:
//
//	func Name[T constraint](arg T) T { ... }
//	        ^^^^^^^^^^^^^^^
//	        T — параметр типа, constraint — ограничение (какие типы допустимы).
//
// `any` — ограничение «любой тип» (это псевдоним для interface{}).
package main

import "fmt"

// PrintSlice печатает срез ЛЮБОГО типа элементов.
// До дженериков пришлось бы писать PrintIntSlice, PrintStringSlice и т.д.
func PrintSlice[T any](s []T) {
	for i, v := range s {
		fmt.Printf("  [%d] %v\n", i, v)
	}
}

// First возвращает первый элемент среза и флаг «срез не пуст».
// Тип возвращаемого значения T выводится из типа аргумента.
func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T // нулевое значение для типа T
		return zero, false
	}
	return s[0], true
}

func main() {
	ints := []int{1, 2, 3}
	strs := []string{"go", "rust", "zig"}

	fmt.Println("ints:")
	PrintSlice(ints) // T выведен как int
	fmt.Println("strings:")
	PrintSlice(strs) // T выведен как string

	if v, ok := First(ints); ok {
		fmt.Println("первый int:", v)
	}
	if v, ok := First(strs); ok {
		fmt.Println("первая строка:", v)
	}

	// Можно указать тип явно, хотя обычно он выводится автоматически.
	PrintSlice[float64]([]float64{1.5, 2.5})
}
