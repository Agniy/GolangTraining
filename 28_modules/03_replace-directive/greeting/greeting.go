// Package greeting — крошечная библиотека, которую мы подключим из app
// через директиву replace (без публикации в интернет).
package greeting

import "fmt"

// Hello возвращает приветствие для name.
func Hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}
