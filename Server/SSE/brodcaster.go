package sse

import (
	"fmt"

	entities "ishtaloo.io/Entities"
)

func Broadcaster(done <-chan interface{}, sseChannel *entities.SSEChannel) {
	fmt.Println("Broadcaster Started.")
	for {
		select {
		case <-done:
			return
		case data := <-sseChannel.Notifier:
			for _, channel := range sseChannel.Clients {
				channel <- data
			}
		}
	}
}
