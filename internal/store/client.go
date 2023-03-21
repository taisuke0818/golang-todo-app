package store

import (
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	c *mongo.Client
}

func New(c *mongo.Client) *Client {
	return &Client{
		c: c,
	}
}
