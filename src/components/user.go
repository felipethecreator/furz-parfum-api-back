package components

import (
	import (
		"context"
		"encoding/json"
		"fmt"
		"log"
	
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)
)

type User struct {
	Name		string `json:"name" bson:"name"`
	DisplayName	string `json:"displayname" bson: "displayname"`
	password	string `json:"password" bson: "password"`
}

// TODO: ACRESCENTAR FUNC DE POSTAR USER 