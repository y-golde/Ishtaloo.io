package entities

type ClientRoom struct {
	RoomId      string `json:"roomId" bson:"_id,omitempty"`
	CurrentWord string `json:"currentWord" bson:"current_word"`
	Guesses     []rune `json:"guesses" bson:"guesses"`
	Users       []User `json:"users" bson:"users"`
}
