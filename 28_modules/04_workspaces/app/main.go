// Приложение в воркспейсе. Использует example.com/greeting напрямую из
// исходников соседнего модуля — связь задаёт go.work в родительской папке.
//
// Запуск из корня воркспейса (где лежит go.work):
//
//	go run ./app
package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	fmt.Println(greeting.Hello("мир"))
}
