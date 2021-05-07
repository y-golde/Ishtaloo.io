package entities

type LoginRequest struct {
	UserName string `json:"userName" bson:"user_name"`
	UserId   string `json:"userId" bson:"user_id"`
}
