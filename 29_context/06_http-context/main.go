// context в HTTP — самый частый практический сценарий.
//
// На СТОРОНЕ СЕРВЕРА: у каждого *http.Request есть r.Context(), который
// отменяется, когда клиент закрыл соединение/ушёл. Долгий обработчик должен
// слушать r.Context().Done() и прекращать работу — не жечь ресурсы впустую.
//
// На СТОРОНЕ КЛИЕНТА: контекст с таймаутом (http.NewRequestWithContext)
// ограничивает время запроса. По истечении таймаута запрос отменяется,
// и сервер видит это в своём r.Context().
//
// Пример самодостаточный: поднимает сервер на свободном порту, шлёт один
// запрос с коротким таймаутом и показывает отмену с обеих сторон.
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func slowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("сервер: начал обработку, работаю ~1с")

	select {
	case <-time.After(1 * time.Second):
		// Успели — отвечаем нормально.
		fmt.Println("сервер: работа завершена, отвечаю 200")
		fmt.Fprintln(w, "done")
	case <-ctx.Done():
		// Клиент отвалился / истёк таймаут клиента.
		fmt.Println("сервер: клиент отменил запрос:", ctx.Err())
	}
}

func main() {
	// Поднимаем сервер на случайном свободном порту 127.0.0.1.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/slow", slowHandler)
	srv := &http.Server{Handler: mux}
	go srv.Serve(ln)

	url := "http://" + ln.Addr().String() + "/slow"

	// Клиент даёт запросу всего 300 мс — меньше, чем длится обработчик.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	fmt.Println("клиент: отправляю запрос с таймаутом 300мс")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Ожидаемо: context deadline exceeded.
		fmt.Println("клиент: запрос прерван:", err)
	} else {
		resp.Body.Close()
		fmt.Println("клиент: получен ответ", resp.Status)
	}

	// Небольшая пауза, чтобы увидеть сообщение сервера об отмене.
	time.Sleep(200 * time.Millisecond)

	// Корректно гасим сервер.
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second)
	defer shutdownCancel()
	_ = srv.Shutdown(shutdownCtx)
}
