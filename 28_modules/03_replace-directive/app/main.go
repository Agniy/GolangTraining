// Приложение использует модуль example.com/greeting, который физически лежит
// рядом на диске и подключён через `replace` в go.mod.
//
// Запуск:
//
//	cd app
//	go run .
package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	fmt.Println(greeting.Hello("мир"))
}
