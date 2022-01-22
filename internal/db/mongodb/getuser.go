package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"web3pay/internal/db"
)

func (c *Client) GetUser(ctx context.Context, userId string) (*db.User, error) {
	coll := c.m.Database("payment").Collection("userInfo")
	var user *db.User
	err := coll.FindOne(
		ctx,
		bson.D{{"userId", userId}},
	).Decode(&user)
	return user, err
}
