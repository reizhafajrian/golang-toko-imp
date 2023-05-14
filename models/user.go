package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JsonToko struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Date     string             `bson:"date"`
	Title    string             `bson:"title"`
	Link     string             `bson:"link"`
	Variants []string           `bson:"variants"`
	Toko     string             `bson:"toko"`
	Category []string           `bson:"category"`
	Stock    int                `bson:"stock"`
}



type BsonToko struct {
	Date     string   `json:"date"`
	Title    string   `json:"title"`
	Link     string   `json:"link"`
	Variants []string `json:"variants"`
	Toko     string   `json:"toko"`
	Category []string `json:"category"`
	Stock    int      `json:"stock"`
}