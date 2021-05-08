package entities

type Room struct {
	RoomId      string   `json:"roomId" bson:"room_id"`
	CurrentWord []string `json:"currentWord" bson:"current_word"`
	Guesses     []string `json:"Guesses" bson:"guesses"`
	Users       []User   `json:"users" bson:"users"`
}
