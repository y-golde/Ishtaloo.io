package entities

type SSEChannel struct {
	Clients  []chan string
	Notifier chan string
}
