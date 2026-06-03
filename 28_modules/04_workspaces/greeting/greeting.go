// Package greeting — библиотечный модуль внутри воркспейса.
package greeting

import "fmt"

// Hello возвращает приветствие для name.
func Hello(name string) string {
	return fmt.Sprintf("Привет из воркспейса, %s!", name)
}
