package models

import (
	"cloud.google.com/go/firestore"
	"log"
	"rent-house/consts"
)

type Admin struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}


func (g *Admin) GetCollectionKey() string {
	return consts.ADMIN
}

func (g *Admin) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (g *Admin) GetFromKey(id string) (error) {
	log.Println(id)
	doc, err := g.GetCollection().Doc(id).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(g)
}