package db

import (
	"time"
)

type User struct {
	Id               string    `json:"userId" bson:"userId" db:"userId"`
	Account          string    `json:"account" bson:"account" db:"account"`
	AvailableBalance float64   `json:"availableBalance" bson:"availableBalance" db:"availableBalance"`
	SpentBalance     float64   `json:"spentBalance" bson:"spentBalance" db:"spentBalance"`
	UpdateTime       time.Time `json:"updateTime" bson:"updateTime" db:"updateTime"`
}

type UserId struct {
	Id string `json:"userId" bson:"userId" db:"userId"`
}

type UserAccount struct {
	Account string `json:"account" bson:"account" db:"account"`
}

type UserBalance struct {
	AvailableBalance float64 `json:"availableBalance" bson:"availableBalance" db:"availableBalance"`
	SpentBalance     float64 `json:"spentBalance" bson:"spentBalance" db:"spentBalance"`
}

type UserAccountAndBalance struct {
	Account          string  `json:"account" bson:"account" db:"account"`
	AvailableBalance float64 `json:"availableBalance" bson:"availableBalance" db:"availableBalance"`
	SpentBalance     float64 `json:"spentBalance" bson:"spentBalance" db:"spentBalance"`
}
