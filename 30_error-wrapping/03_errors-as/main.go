// errors.As — извлечение из цепочки ошибки КОНКРЕТНОГО ТИПА.
//
// errors.Is отвечает на вопрос «это та самая ошибка?», а errors.As —
// «есть ли в цепочке ошибка такого типа, и если да — дай мне её», чтобы
// прочитать её поля/методы.
//
// Это замена устаревшему type assertion `if e, ok := err.(*MyError); ok`,
// который не умеет заглядывать под обёртки %w.
package main

import (
	"errors"
	"fmt"
)

// Своя структурированная ошибка с дополнительными полями.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %q: %s", e.Field, e.Msg)
}

func validate() error {
	return &ValidationError{Field: "email", Msg: "неверный формат"}
}

func handleRequest() error {
	if err := validate(); err != nil {
		return fmt.Errorf("handleRequest: %w", err) // обернули
	}
	return nil
}

func main() {
	err := handleRequest()
	fmt.Println("ошибка:", err)

	// Достаём *ValidationError из-под обёртки и читаем его поля.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("→ ошибка валидации: поле=%s, причина=%s\n", ve.Field, ve.Msg)
	} else {
		fmt.Println("→ это не ошибка валидации")
	}
}
