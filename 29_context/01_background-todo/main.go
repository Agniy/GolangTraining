// Пакет context (Go 1.7, 2016) переносит между горутинами и вызовами функций:
//   - сигнал отмены / дедлайн;
//   - значения, привязанные к запросу (request-scoped values).
//
// Любое дерево контекстов начинается с КОРНЕВОГО контекста. Их два:
//
//	context.Background() — пустой корень. Используется в main, в init,
//	                       в самых верхних обработчиках входящих запросов.
//	context.TODO()       — заглушка "ещё не решил, какой контекст сюда нужен".
//	                       Сигнал читателю кода и линтерам, что место временное.
//
// Оба пустые: их нельзя отменить, у них нет дедлайна и значений.
// Реальную функциональность добавляют производные контексты:
// context.WithCancel / WithTimeout / WithDeadline / WithValue (см. примеры 02–05).
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// У пустого корня нет дедлайна...
	deadline, ok := ctx.Deadline()
	fmt.Printf("Background: deadline=%v, есть дедлайн? %v\n", deadline, ok)

	// ...нет ошибки (он никогда не отменяется)...
	fmt.Printf("Background: Err()=%v\n", ctx.Err())

	// ...и нет значений.
	fmt.Printf("Background: Value(\"key\")=%v\n", ctx.Value("key"))

	// TODO ведёт себя так же — это маркер «контекст сюда добавим позже».
	todo := context.TODO()
	fmt.Printf("TODO: Err()=%v\n", todo.Err())
}
