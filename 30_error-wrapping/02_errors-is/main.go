// errors.Is — проверка, есть ли В ЦЕПОЧКЕ ошибок конкретная «сигнальная»
// ошибка (sentinel error).
//
// Раньше писали `if err == ErrNotFound`, но это ломается, как только ошибку
// обернули. errors.Is разворачивает цепочку (по %w) и сравнивает на каждом
// уровне — поэтому работает сквозь любую глубину обёрток.
package main

import (
	"errors"
	"fmt"
)

var ErrPermission = errors.New("доступ запрещён")

// Несколько уровней обёртки, как в реальном стеке вызовов.
func readFile() error {
	return ErrPermission
}

func loadConfig() error {
	if err := readFile(); err != nil {
		return fmt.Errorf("loadConfig: %w", err)
	}
	return nil
}

func startApp() error {
	if err := loadConfig(); err != nil {
		return fmt.Errorf("startApp: %w", err)
	}
	return nil
}

func main() {
	err := startApp()
	fmt.Println("полная ошибка:", err)

	// Несмотря на две обёртки, Is находит ErrPermission в глубине цепочки.
	if errors.Is(err, ErrPermission) {
		fmt.Println("→ обнаружена проблема с правами доступа")
	}

	// Чужой ошибки в цепочке нет.
	fmt.Println("это io.EOF?", errors.Is(err, errors.New("eof")))
}
