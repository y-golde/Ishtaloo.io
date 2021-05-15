package roomUtils

import (
	entities "ishtaloo.io/Entities"
)

func RoomToClientRoom(room *entities.Room) entities.ClientRoom {
	EncryptWord(room.CurrentWord, room.Guesses)
	clientRoom := entities.ClientRoom{
		RoomId:      room.RoomId,
		CurrentWord: string(room.CurrentWord),
		Guesses:     room.Guesses,
		Users:       room.Users,
	}
	return clientRoom
}
