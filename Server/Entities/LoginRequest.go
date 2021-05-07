package entities

type LoginRequest struct {
	UserName string `json:"userName" bson:"userName"`
}
