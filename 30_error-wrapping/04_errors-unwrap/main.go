// errors.Unwrap — снять ОДИН слой обёртки.
//
// Возвращает ошибку, вложенную через %w, или nil, если её нет.
// На практике напрямую Unwrap нужен редко (обычно хватает Is/As), но он
// показывает, КАК устроена цепочка: Is/As внутри просто вызывают Unwrap
// в цикле.
//
// Также здесь показан кастомный тип ошибки со своим методом Unwrap() error —
// именно его ищут стандартные функции, разбирая цепочку.
package main

import (
	"errors"
	"fmt"
)

var ErrDB = errors.New("сбой базы данных")

// Своя ошибка, которая умеет «разворачиваться» в обёрнутую.
type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("запрос %q: %v", e.Query, e.Err)
}

// Метод Unwrap делает QueryError частью цепочки для errors.Is/As/Unwrap.
func (e *QueryError) Unwrap() error { return e.Err }

func main() {
	// Строим цепочку из трёх уровней.
	err := fmt.Errorf("сервис недоступен: %w",
		&QueryError{Query: "SELECT 1", Err: ErrDB})

	fmt.Println("исходная ошибка:", err)
	fmt.Println("--- разворачиваем по одному слою ---")

	for err != nil {
		fmt.Printf("  %v\n", err)
		err = errors.Unwrap(err) // снимаем один слой за итерацию
	}
	fmt.Println("дно цепочки достигнуто (nil)")
}
