package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"web3pay/internal/db"
)
func (c *Client) DeleteUser(ctx context.Context, userId string) (*db.User, error) {
	coll := c.m.Database("payment").Collection("userInfo")
	result := coll.FindOneAndDelete(
		ctx,
		bson.M{"userId": userId},
	)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var deletedUser *db.User
	err := result.Decode(&deletedUser)
	return deletedUser, err
}
