// Обобщённые ТИПЫ: структуры и их методы тоже могут иметь параметры типа.
//
// Здесь — типобезопасный стек Stack[T]. До дженериков пришлось бы либо
// хранить []any (и приводить типы в рантайме, рискуя паникой), либо
// копировать код стека под каждый тип элемента.
//
// Синтаксис:
//
//	type Stack[T any] struct { ... }          // параметр типа у структуры
//	func (s *Stack[T]) Push(v T) { ... }      // методы повторяют [T]
//
// Важно: у методов НЕЛЬЗЯ вводить собственные новые параметры типа —
// они работают только с параметрами своего типа-получателя.
package main

import "fmt"

// Stack — стек элементов типа T (LIFO).
type Stack[T any] struct {
	items []T
}

// Push кладёт элемент на вершину.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop снимает элемент с вершины. Второе значение — «стек был не пуст».
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func (s *Stack[T]) Len() int { return len(s.items) }

func main() {
	// Стек целых чисел.
	var ints Stack[int]
	ints.Push(10)
	ints.Push(20)
	ints.Push(30)
	fmt.Println("размер:", ints.Len())
	for {
		v, ok := ints.Pop()
		if !ok {
			break
		}
		fmt.Println("pop:", v)
	}

	// Тот же тип, но для строк — без единой строчки нового кода.
	var strs Stack[string]
	strs.Push("a")
	strs.Push("b")
	v, _ := strs.Pop()
	fmt.Println("строковый стек, pop:", v)

	// Параметром типа может быть и структура.
	type Point struct{ X, Y int }
	var pts Stack[Point]
	pts.Push(Point{1, 2})
	p, _ := pts.Pop()
	fmt.Println("стек точек, pop:", p)
}
