package main

import (
	"context"
	"fmt"
	"time"
)

// Существует 2 канала,
// в бесконечном цикле каждые 0.5 сек в первый канал заносится значение счетчика (Считает количество итераций).
// Если число четное его записать во второй канал.
// При получения значения во втором канале значение выводится в консоль.
// Через 4 секунды корректо завершить работу всех горутин.
func main() {
	count := 1
	allCh := make(chan int)
	evenCh := make(chan int)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(time.Millisecond * 500):
				allCh <- count
				if count%2 == 0 {
					evenCh <- count
				}
				count++
			case <-ctx.Done():
				close(allCh)
				close(evenCh)
				return
			}
		}
	}(ctxWithTimeout)

	go func(ctx context.Context) {
		for {
			select {
			case <-allCh:
				continue
			case <-ctx.Done():
				return
			}
		}
	}(ctxWithTimeout)

	for even := range evenCh {
		fmt.Println(even)
	}
}
