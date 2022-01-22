package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"web3pay/internal/db"
)

func (c *Client) GetUserBalance(ctx context.Context, userId string) (*db.UserBalance, error) {
	coll := c.m.Database("payment").Collection("userInfo")
	var balance *db.UserBalance
	err := coll.FindOne(
		ctx,
		bson.D{
			{"userId", userId},
		},
	).Decode(&balance)
	return balance, err
}
