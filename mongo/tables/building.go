package tables

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Building struct {
	mgm.DefaultModel
	CompanyId   primitive.ObjectID
	Name        string
	Coordinates []int
	ClassId     int
	Level       int
}
