package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Client struct {
	m *mongo.Client
}

func NewClient(connURI string) *Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connURI))
	if err != nil {
		log.Fatalln(err)
	}
	return &Client{client}
}

func (c *Client) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.m.Connect(ctx)
	return err
}

func (c *Client) Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c.m.Disconnect(ctx)
}
