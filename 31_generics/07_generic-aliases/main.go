// Обобщённые алиасы типов (generic type aliases), Go 1.24 (2025).
//
// Алиас типа (`type A = B`) появился ещё в Go 1.9, но до 1.24 он НЕ мог иметь
// параметров типа. С Go 1.24 так можно:
//
//	type Set[T comparable] = map[T]bool
//	         ^^^^^^^^^^^^^^   ^ знак = делает это именно алиасом, а не новым типом
//
// Алиас — это ДРУГОЕ ИМЯ для того же типа (не новый тип). Поэтому Set[string]
// и map[string]bool полностью взаимозаменяемы: можно передавать одно туда,
// где ждут другое, без конвертации.
//
// Зачем: давать короткие, говорящие имена громоздким обобщённым типам и
// постепенно мигрировать API, не ломая совместимость.
package main

import "fmt"

// Обобщённый алиас: короткое имя для map[T]bool.
type Set[T comparable] = map[T]bool

// Ещё один: пара ключ-значение поверх обычной map.
type Dict[K comparable, V any] = map[K]V

// Функция объявлена в терминах сырого map[T]bool...
func add[T comparable](m map[T]bool, v T) {
	m[v] = true
}

func main() {
	// ...но мы свободно передаём ей Set[string] — это один и тот же тип.
	s := Set[string]{"go": true}
	add(s, "rust") // алиас взаимозаменяем с map[string]bool
	fmt.Println("Set:", s)

	// Алиас можно создавать и через make.
	seen := make(Set[int])
	for _, n := range []int{1, 2, 2, 3, 3, 3} {
		seen[n] = true
	}
	fmt.Println("уникальных чисел:", len(seen))

	// Многопараметрический алиас.
	ages := Dict[string, int]{"Иван": 30, "Мария": 25}
	fmt.Println("Dict:", ages)

	// Доказательство «это тот же тип»: присваиваем Set[string] переменной
	// типа map[string]bool без всякой конвертации.
	var raw map[string]bool = s
	fmt.Println("как обычная map:", raw["go"])
}
