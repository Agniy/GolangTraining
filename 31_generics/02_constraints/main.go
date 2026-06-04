// Ограничения (constraints) — какие типы разрешено подставлять вместо T.
//
// Ограничение — это интерфейс. Бывают трёх видов:
//
//  1. any         — любой тип (см. пример 01);
//  2. comparable  — встроенное ограничение: типы, поддерживающие == и !=
//     (нужно, например, для ключей map);
//  3. свой интерфейс с МНОЖЕСТВОМ ТИПОВ — перечисление допустимых типов
//     через | . Это новая возможность Go 1.18.
//
// Для упорядочиваемых типов в стандартной библиотеке с Go 1.21 есть готовое
// ограничение cmp.Ordered (раньше его брали из golang.org/x/exp/constraints).
package main

import (
	"cmp"
	"fmt"
)

// comparable: позволяет сравнивать значения через ==.
func Index[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

// cmp.Ordered: типы, поддерживающие < > <= >= (числа и строки).
// Поэтому внутри можно использовать оператор сравнения.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Своё ограничение: только числовые типы. Префикс ~ означает «а также любой
// тип, чьим базовым типом является этот» (например, `type Celsius float64`).
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	fmt.Println("Index:", Index([]string{"a", "b", "c"}, "b")) // 1
	fmt.Println("Max(int):", Max(3, 7))                        // 7
	fmt.Println("Max(string):", Max("apple", "banana"))        // banana
	fmt.Println("Sum(int):", Sum([]int{1, 2, 3, 4}))           // 10
	fmt.Println("Sum(float):", Sum([]float64{1.1, 2.2, 3.3}))  // 6.6
}
