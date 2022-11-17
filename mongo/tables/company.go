package tables

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	mgm.DefaultModel
	UserId primitive.ObjectID
	Name   string
}

func NewCompany(userId primitive.ObjectID, name string) *Company {
	return &Company{
		UserId: userId,
		Name:   name,
	}
}
