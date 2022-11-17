package tables

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel
	Email     string `bson:"email" json:"email"`
	SessionId string `bson:"session_id" json:"sessionId"`
	Provider  string `bson:"provider" json:"provider"`
	CreatedIp string `bson:"created_ip" json:"createdIp"`
}

func NewUser(email string, sessionId string, provider string, createdIp string) *User {
	return &User{
		Email:     email,
		SessionId: sessionId,
		Provider:  provider,
		CreatedIp: createdIp,
	}
}
