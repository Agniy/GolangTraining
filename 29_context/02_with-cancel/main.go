// context.WithCancel — ручная отмена.
//
// Возвращает производный контекст и функцию cancel. Вызов cancel() закрывает
// канал ctx.Done(): все горутины, слушающие этот канал, получают сигнал
// остановиться. Это идиоматичный способ корректно завершать горутины и
// не плодить утечки.
//
// Правило: cancel НУЖНО вызывать всегда (обычно через defer), иначе
// освобождение ресурсов контекста откладывается — go vet это проверяет.
package main

import (
	"context"
	"fmt"
	"time"
)

// worker считает, пока контекст не отменят.
func worker(ctx context.Context) {
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			// ctx.Err() объяснит причину: context.Canceled.
			fmt.Println("worker: останавливаюсь,", ctx.Err())
			return
		default:
			fmt.Println("worker: тик", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // подстраховка, если выйдем раньше

	go worker(ctx)

	// Пусть поработает немного...
	time.Sleep(700 * time.Millisecond)

	// ...и отменяем.
	fmt.Println("main: вызываю cancel()")
	cancel()

	// Дадим worker'у увидеть сигнал и напечатать сообщение.
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: готово")
}
