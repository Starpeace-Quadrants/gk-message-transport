package mongo

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func init() {
	host := os.Getenv("MONGO_SERVER_HOST")
	port := os.Getenv("MONGO_SERVER_PORT")
	database := os.Getenv("MONGO_SERVER_DATABASE")

	if err := mgm.SetDefaultConfig(nil, database, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))); err != nil {
		panic(err)
	}
}
