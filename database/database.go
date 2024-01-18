package database

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	db *mongo.Client
}
