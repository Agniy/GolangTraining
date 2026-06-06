// Пакет maps (Go 1.21) — обобщённые функции для работы с map.
//
// Как и slices, реализован на дженериках и работает с любыми типами ключей и
// значений. Самое частое — Keys/Values (получить ключи или значения) и Clone
// (поверхностная копия).
//
// ВАЖНО про Keys/Values: в Go 1.21 они возвращали срез, но в Go 1.23 их
// заменили на ИТЕРАТОРЫ (iter.Seq, см. раздел 36). Чтобы пример работал на
// go 1.21, ключи/значения здесь собираем циклом range вручную — это надёжно на
// любой версии. Ниже в комментарии показан современный способ (Go 1.23+).
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}

	// --- Ключи и значения ---
	// map не упорядочен, поэтому сразу сортируем, чтобы вывод был стабильным.
	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	fmt.Println("Keys:", keys) // [Alice Bob Carol]

	// Современный способ (Go 1.23+), когда maps.Keys возвращает итератор:
	//   keys := slices.Sorted(maps.Keys(ages))

	// --- Clone: поверхностная копия map ---
	// Без неё две переменные ссылались бы на ОДИН и тот же map.
	clone := maps.Clone(ages)
	clone["Alice"] = 99
	fmt.Println("original Alice:", ages["Alice"]) // 30 — оригинал не изменился
	fmt.Println("clone Alice:", clone["Alice"])   // 99

	// --- Equal: поэлементное сравнение двух map ---
	fmt.Println("Equal(ages, clone):", maps.Equal(ages, clone)) // false

	another := map[string]int{"Alice": 30, "Bob": 25, "Carol": 35}
	fmt.Println("Equal(ages, another):", maps.Equal(ages, another)) // true

	// --- Copy: скопировать все пары из одного map в другой ---
	// Существующие ключи перезаписываются, новые добавляются.
	dst := map[string]int{"Bob": 0, "Dave": 40}
	maps.Copy(dst, ages)
	fmt.Println("Copy result Bob:", dst["Bob"])   // 25 (перезаписан)
	fmt.Println("Copy result Dave:", dst["Dave"]) // 40 (сохранён)

	// --- DeleteFunc: удалить пары по условию ---
	maps.DeleteFunc(dst, func(k string, v int) bool {
		return v < 30 // удалить всех, кому меньше 30
	})
	remaining := make([]string, 0, len(dst))
	for k := range dst {
		remaining = append(remaining, k)
	}
	slices.Sort(remaining)
	fmt.Println("After DeleteFunc(<30):", remaining) // [Alice Carol Dave]
}
