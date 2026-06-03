// context.WithTimeout — автоматическая отмена через заданный интервал.
//
// Это сокращение для WithDeadline(parent, time.Now().Add(d)).
// По истечении таймаута ctx.Done() закрывается, а ctx.Err() возвращает
// context.DeadlineExceeded.
//
// Типичный сценарий: «операция должна уложиться в N секунд, иначе бросаем».
package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation имитирует долгую работу, уважающую отмену контекста.
func slowOperation(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d): // работа завершилась сама
		return nil
	case <-ctx.Done(): // контекст истёк/отменён раньше
		return ctx.Err()
	}
}

func main() {
	// Даём операции максимум 300 мс.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Случай 1: операция (100 мс) укладывается в таймаут.
	if err := slowOperation(ctx, 100*time.Millisecond); err != nil {
		fmt.Println("операция 1: ошибка:", err)
	} else {
		fmt.Println("операция 1: успех")
	}

	// Случай 2: операция (1 с) не укладывается — сработает DeadlineExceeded.
	if err := slowOperation(ctx, 1*time.Second); err != nil {
		fmt.Println("операция 2: ошибка:", err)
	} else {
		fmt.Println("операция 2: успех")
	}
}
