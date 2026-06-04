// errors.Join — объединение НЕСКОЛЬКИХ ошибок в одну (Go 1.20, 2023).
//
// До этого, чтобы вернуть сразу несколько ошибок (например, все ошибки
// валидации формы), приходилось писать свои типы или склеивать строки.
// errors.Join делает это штатно:
//   - объединяет список ошибок (nil-ы игнорируются);
//   - в тексте печатает каждую с новой строки;
//   - errors.Is / errors.As находят ЛЮБУЮ из объединённых ошибок.
package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName = errors.New("имя не заполнено")
	ErrBadEmail  = errors.New("неверный email")
	ErrTooYoung  = errors.New("возраст меньше 18")
)

type Form struct {
	Name  string
	Email string
	Age   int
}

func (f Form) Validate() error {
	var errs []error // копим все нарушения

	if f.Name == "" {
		errs = append(errs, ErrEmptyName)
	}
	if f.Email == "" {
		errs = append(errs, ErrBadEmail)
	}
	if f.Age < 18 {
		errs = append(errs, ErrTooYoung)
	}

	// Join вернёт nil, если errs пуст — удобно возвращать напрямую.
	return errors.Join(errs...)
}

func main() {
	err := Form{Name: "", Email: "", Age: 15}.Validate()

	fmt.Println("ошибки валидации:")
	fmt.Println(err)

	// Is видит каждую из объединённых ошибок.
	fmt.Println("--- проверки ---")
	fmt.Println("пустое имя?  ", errors.Is(err, ErrEmptyName))
	fmt.Println("плохой email?", errors.Is(err, ErrBadEmail))
	fmt.Println("молод?       ", errors.Is(err, ErrTooYoung))

	// Корректная форма → ошибок нет.
	fmt.Println("--- валидная форма ---")
	fmt.Println("err:", Form{Name: "Иван", Email: "i@e.ru", Age: 30}.Validate())
}
