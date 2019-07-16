package cpfCnpj

import "time"

type CpfCnpj struct {
	Id string `json:"id" bson:"_id"`
	Number string `json:"number" bson:"number"`
	Type string `json:"type" bson:"type"`
	BlackList bool `json:"blacklist" bson:"blacklist"`
	CreateDate time.Time `json:"createdate" bson:"createdate"`
}
