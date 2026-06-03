// context.WithValue — передача данных, привязанных к запросу.
//
// Кладёт пару ключ/значение в контекст; вложенные функции достают её через
// ctx.Value(key). Типично для сквозных данных: request ID, пользователь,
// данные трассировки.
//
// ВАЖНЫЕ правила хорошего тона:
//   - ключ должен быть СВОЕГО неэкспортируемого типа, а не string/int —
//     иначе ключи из разных пакетов могут столкнуться;
//   - не передавайте через контекст обязательные параметры функций —
//     только сквозные, опциональные, request-scoped данные.
package main

import (
	"context"
	"fmt"
)

// Свой приватный тип ключа исключает коллизии с чужими ключами.
type ctxKey string

const requestIDKey ctxKey = "requestID"

// handle достаёт значение из контекста, не зная, кто и где его положил.
func handle(ctx context.Context) {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("handle: requestID =", id)
	} else {
		fmt.Println("handle: requestID отсутствует")
	}
}

func main() {
	ctx := context.Background()

	// Кладём значение в производный контекст.
	ctx = context.WithValue(ctx, requestIDKey, "req-12345")

	handle(ctx)

	// А вот по «чужому» ключу того же содержания значения НЕ будет —
	// тип ключа другой (string, а не ctxKey).
	fmt.Println("по строковому ключу:", ctx.Value("requestID"))
}
