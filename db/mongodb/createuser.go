package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (c *Client) CreateUser(ctx context.Context, userId, account string) error {
	coll := c.m.Database("payment").Collection("userInfo")
	_, err := coll.UpdateOne(
		ctx,
		bson.D{{"userId", userId}},
		bson.D{
			{"$set", bson.D{
				{"userId", userId},
				{"account", account},
				{"balance", 0.0},
				{"updateTime", time.Now().UTC()},
			}},
		},
		options.Update().SetUpsert(true),
	)
	return err
}
