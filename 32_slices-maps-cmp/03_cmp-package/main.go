// Пакет cmp (Go 1.21) — сравнение упорядочиваемых значений.
//
// Маленький, но полезный пакет. Содержит:
//
//   - cmp.Ordered  — ограничение (constraint) для дженериков: все типы,
//     поддерживающие < > <= >= (числа и строки). Уже встречалось в разделе 31.
//   - cmp.Compare(a, b) — возвращает -1 / 0 / +1. Удобно как готовый
//     компаратор для slices.SortFunc.
//   - cmp.Less(a, b)    — true, если a < b.
//   - cmp.Or(vals...)   — первое НЕнулевое значение (Go 1.22).
package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// --- Compare: -1 / 0 / +1 ---
	fmt.Println("Compare(1, 2):", cmp.Compare(1, 2)) // -1
	fmt.Println("Compare(2, 2):", cmp.Compare(2, 2)) // 0
	fmt.Println("Compare(3, 2):", cmp.Compare(3, 2)) // 1

	// --- Less: булево сравнение ---
	fmt.Println("Less(\"a\", \"b\"):", cmp.Less("a", "b")) // true

	// --- Compare как компаратор для сортировки структур ---
	// Сортируем людей по возрасту, а при равном возрасте — по имени.
	people := []Person{
		{"Bob", 30},
		{"Alice", 30},
		{"Carol", 25},
	}
	slices.SortFunc(people, func(a, b Person) int {
		// cmp.Or берёт первое ненулевое сравнение: сначала по возрасту,
		// и только если возраст равен (Compare вернул 0) — по имени.
		return cmp.Or(
			cmp.Compare(a.Age, b.Age),
			cmp.Compare(a.Name, b.Name),
		)
	})
	fmt.Println("Sorted people:", people)
	// [{Carol 25} {Alice 30} {Bob 30}]

	// --- Or: первое ненулевое значение (Go 1.22) ---
	// Частый приём для значений по умолчанию.
	fmt.Println("Or(\"\", \"def\"):", cmp.Or("", "def"))       // def
	fmt.Println("Or(\"set\", \"def\"):", cmp.Or("set", "def")) // set
	fmt.Println("Or(0, 0, 7, 9):", cmp.Or(0, 0, 7, 9))         // 7
}
