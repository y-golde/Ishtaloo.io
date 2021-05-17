package sse

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func getSSE(c echo.Context, sseChannel *entities.SSEChannel) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	sseChan := make(chan string)
	sseChannel.Clients = append(sseChannel.Clients, sseChan)

	d := make(chan interface{})
	defer close(d)
	defer fmt.Println("Closing channel.")

	for {
		select {
		case <-d:
			close(sseChan)
			return c.NoContent(http.StatusOK)
		case data := <-sseChan:
			fmt.Fprintf(c.Response(), "new event: %v \n\n", data)
			c.Response().Flush()
		}
	}
}
