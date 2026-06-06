// Пакет slices (Go 1.21) — обобщённые функции для работы со срезами.
//
// Раньше такие операции писали руками или брали из golang.org/x/exp/slices.
// Теперь они в стандартной библиотеке, реализованы на дженериках (см. раздел
// 31), поэтому работают с любым типом элементов и проверяются на этапе
// компиляции.
//
// Здесь показаны самые употребительные функции: Sort, Contains, Index, а также
// несколько дополнительных, которые часто нужны на практике.
package main

import (
	"fmt"
	"slices"
)

func main() {
	// --- Sort: сортировка на месте (in-place) ---
	// Работает для любых упорядочиваемых типов (cmp.Ordered): чисел и строк.
	nums := []int{5, 2, 8, 1, 9, 3}
	slices.Sort(nums)
	fmt.Println("Sort:", nums) // [1 2 3 5 8 9]

	words := []string{"banana", "apple", "cherry"}
	slices.Sort(words)
	fmt.Println("Sort strings:", words) // [apple banana cherry]

	// --- Contains: есть ли элемент в срезе ---
	fmt.Println("Contains 8:", slices.Contains(nums, 8)) // true
	fmt.Println("Contains 7:", slices.Contains(nums, 7)) // false

	// --- Index: позиция первого вхождения (или -1) ---
	fmt.Println("Index of 8:", slices.Index(nums, 8)) // 4
	fmt.Println("Index of 7:", slices.Index(nums, 7)) // -1

	// --- BinarySearch: бинарный поиск в ОТСОРТИРОВАННОМ срезе ---
	// Возвращает позицию и флаг «найдено». Если не найдено — позиция, куда
	// элемент следовало бы вставить, сохранив порядок.
	pos, found := slices.BinarySearch(nums, 8)
	fmt.Printf("BinarySearch 8: pos=%d found=%t\n", pos, found) // pos=4 found=true

	// --- Min / Max: минимум и максимум среза ---
	// (паникуют на пустом срезе — не путать со встроенными min/max из раздела 33)
	fmt.Println("Min:", slices.Min(nums)) // 1
	fmt.Println("Max:", slices.Max(nums)) // 9

	// --- Reverse: разворот на месте ---
	slices.Reverse(nums)
	fmt.Println("Reverse:", nums) // [9 8 5 3 2 1]

	// --- Equal: поэлементное сравнение двух срезов ---
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println("Equal:", slices.Equal(a, b)) // true

	// --- SortFunc: сортировка по своему компаратору ---
	// Компаратор возвращает отрицательное/0/положительное (как cmp.Compare).
	people := []string{"Tom", "Al", "Bob", "Jo"}
	slices.SortFunc(people, func(x, y string) int {
		return len(x) - len(y) // по длине строки
	})
	fmt.Println("SortFunc by len:", people) // [Al Jo Tom Bob]
}
