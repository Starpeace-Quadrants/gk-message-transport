package tables

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	mgm.DefaultModel
	UserId    primitive.ObjectID
	Alias     string
	ImagePath string
}

func NewUserProfile(userId primitive.ObjectID, alias string) *UserProfile {
	return &UserProfile{
		UserId: userId,
		Alias:  alias,
	}
}
