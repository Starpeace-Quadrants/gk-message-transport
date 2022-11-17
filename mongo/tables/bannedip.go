package tables

import (
	"github.com/kamva/mgm/v3"
)

type BannedIp struct {
	mgm.DefaultModel
	Host string `bson:"host" json:"host"`
}

func NewBannedIp(host string) *BannedIp {
	return &BannedIp{
		Host: host,
	}
}
