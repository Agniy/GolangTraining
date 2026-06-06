// Встроенные функции min и max (Go 1.21).
//
// До Go 1.21 их не было: писали `if a < b { ... }` либо math.Min/Max (только
// для float64) или свою дженерик-функцию (см. раздел 31). Теперь min и max —
// ВСТРОЕННЫЕ: импорт не нужен, работают с любым числом аргументов и любым
// упорядочиваемым типом (числа и строки).
//
// Не путать со slices.Min/slices.Max (раздел 32): те принимают ОДИН срез,
// а встроенные min/max — два и более ОТДЕЛЬНЫХ аргумента.
package main

import "fmt"

func main() {
	// --- Два аргумента ---
	fmt.Println("min(3, 7):", min(3, 7)) // 3
	fmt.Println("max(3, 7):", max(3, 7)) // 7

	// --- Сколько угодно аргументов ---
	fmt.Println("min(5, 2, 8, 1):", min(5, 2, 8, 1)) // 1
	fmt.Println("max(5, 2, 8, 1):", max(5, 2, 8, 1)) // 8

	// --- Работают с любым упорядочиваемым типом, не только int ---
	fmt.Println("max(3.14, 2.71):", max(3.14, 2.71))            // 3.14
	fmt.Println("min строк:", min("banana", "apple", "cherry")) // apple

	// --- Тип результата совпадает с типом аргументов ---
	var a, b uint8 = 200, 100
	fmt.Printf("min(uint8): %d (%T)\n", min(a, b), min(a, b)) // 100 (uint8)

	// --- Практический приём: ограничить значение диапазоном [lo, hi] ---
	// clamp(x) = max(lo, min(x, hi))
	lo, hi := 0, 100
	clamp := func(x int) int { return max(lo, min(x, hi)) }
	fmt.Println("clamp(-20):", clamp(-20)) // 0
	fmt.Println("clamp(50):", clamp(50))   // 50
	fmt.Println("clamp(150):", clamp(150)) // 100
}
