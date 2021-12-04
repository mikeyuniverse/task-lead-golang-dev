package models

import (
	"grpc-practice/pkg/proto/transport"
)

type Item struct {
	Name  string `bson:"name"`
	Price int    `bson:"price"`
}

func (i *Item) ToPB() *transport.Item {
	return &transport.Item{
		Name:  i.Name,
		Price: int32(i.Price),
	}
}
