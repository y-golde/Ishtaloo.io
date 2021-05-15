package entities

type Room struct {
	RoomId      string   `json:"roomId" bson:"_id,omitempty"`
	CurrentWord []string `json:"currentWord" bson:"current_word"`
	Guesses     []string `json:"guesses" bson:"guesses"`
	Users       []User   `json:"users" bson:"users"`
}
