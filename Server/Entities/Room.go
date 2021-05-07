package entities

type Room struct {
	RoomId       string `json:"roomId" bson:"room_id"`
	CurrentWord  []rune `json:"currentWord" bson:"current_word"`
	WrongGuesses []rune `json:"wrongGuesses" bson:"wrong_guesses"`
	Users        []User `json:"users" bson:"users"`
}
