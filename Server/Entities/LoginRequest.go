package entities

type LoginRequest struct {
	UserName string `json:"userName" bson:"userName"`
	UserId   string `json:"userId" bson:"userId"`
}
