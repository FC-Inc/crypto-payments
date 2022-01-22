package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"web3pay/internal/db"
)

func (c *Client) GetUserAccount(ctx context.Context, userId string) (*db.UserAccount, error) {
	coll := c.m.Database("payment").Collection("userInfo")
	var userAccount *db.UserAccount
	err := coll.FindOne(
		ctx,
		bson.D{{"userId", userId}},
	).Decode(&userAccount)
	return userAccount, err
}
