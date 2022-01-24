package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"web3pay/internal/db"
)
func (c *Client) DecUserBalance(ctx context.Context, userId string, balanceChange float64) (*db.UserBalance, error) {
	coll := c.m.Database("payment").Collection("userInfo")
	result := coll.FindOneAndUpdate(
		ctx,
		bson.M{"userId": userId},
		bson.M{
			"$dec": bson.M{"balance": balanceChange},
		},
	)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var updatedBalance *db.UserBalance
	err := result.Decode(&updatedBalance)
	return updatedBalance, err
}
