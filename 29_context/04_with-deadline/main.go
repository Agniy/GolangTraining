// context.WithDeadline — отмена в КОНКРЕТНЫЙ момент времени.
//
// Отличие от WithTimeout: задаём не длительность, а абсолютное время
// (time.Time). Удобно, когда дедлайн приходит извне — например, «ответить
// до 12:00:00» или «не позже срока, спущенного вызывающим».
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Дедлайн — через 250 мс от текущего момента.
	deadline := time.Now().Add(250 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Узнать дедлайн можно из самого контекста.
	if d, ok := ctx.Deadline(); ok {
		fmt.Println("дедлайн установлен на:", d.Format("15:04:05.000"))
	}

	// Пытаемся выполнить работу, которая длится дольше дедлайна.
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("работа завершена в срок")
	case <-ctx.Done():
		// Здесь будет context.DeadlineExceeded.
		fmt.Println("дедлайн истёк:", ctx.Err())
	}
}
