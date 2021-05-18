package roomUtils

import entities "ishtaloo.io/Entities"

func UserExistsInSlice(user *entities.User, users []entities.User) int {
	for k, curUser := range users {
		if curUser.UserId == user.UserId {
			return k
		}
	}
	return -1
}
