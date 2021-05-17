package entities

type Guess struct {
	Guess  string `json:"guess"`
	RoomId string `json:"roomId"`
}
