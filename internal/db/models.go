package db

import (
	"time"
)

type User struct {
	Id         string    `json:"userId" bson:"userId" db:"userId"`
	Account    string    `json:"account" bson:"account" db:"account"`
	Balance    float64   `json:"balance" bson:"balance" db:"balance"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime" db:"updateTime"`
}

type UserId struct {
	Id string `json:"userId" bson:"userId" db:"userId"`
}

type UserAccount struct {
	Account string `json:"account" bson:"account" db:"account"`
}

type UserBalance struct {
	Balance float64 `json:"balance" bson:"balance" db:"balance"`
}
