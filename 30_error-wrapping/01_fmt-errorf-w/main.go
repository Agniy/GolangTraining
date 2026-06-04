// Обёртка ошибок (error wrapping), Go 1.13 (2019).
//
// До 1.13 при добавлении контекста к ошибке исходная ошибка терялась:
//
//	return fmt.Errorf("loadConfig: %v", err)   // err превращается в строку
//
// Глагол %w в fmt.Errorf СОХРАНЯЕТ исходную ошибку внутри новой —
// получается «цепочка» ошибок. Позже её можно разобрать через
// errors.Is / errors.As / errors.Unwrap (примеры 02–04).
//
// Разница на практике:
//
//	%v — только текст, связь с исходной ошибкой потеряна;
//	%w — текст + сохранённая ссылка на исходную ошибку.
package main

import (
	"errors"
	"fmt"
)

// Базовая «сигнальная» ошибка пакета.
var ErrNotFound = errors.New("не найдено")

func findUser(id int) error {
	// Оборачиваем ErrNotFound, добавляя контекст, но НЕ теряя её.
	return fmt.Errorf("findUser(%d): %w", id, ErrNotFound)
}

func main() {
	err := findUser(42)
	fmt.Println("ошибка:", err)

	// %w сохранил связь — errors.Is находит ErrNotFound в цепочке.
	fmt.Println("обёртка через w-глагол, найдено ErrNotFound?", errors.Is(err, ErrNotFound))

	// Для сравнения — обёртка через v-глагол связь теряет.
	errV := fmt.Errorf("findUser: %v", ErrNotFound)
	fmt.Println("обёртка через v-глагол, найдено ErrNotFound?", errors.Is(errV, ErrNotFound))
}
