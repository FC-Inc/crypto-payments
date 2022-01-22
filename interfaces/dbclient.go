package interfaces

import (
	"context"
	"web3pay/internal/db"
)

type DBClient interface {
	CreateUser(ctx context.Context, userId, account string) error
	GetUser(ctx context.Context, userId string) (*db.User, error)
	GetUserAccount(ctx context.Context, userId string) (*db.UserAccount, error)
	GetUserBalance(ctx context.Context, userId string) (*db.UserBalance, error)
	IncUserBalance(ctx context.Context, userId string, balanceChange float64) (*db.UserBalance, error)
	DecUserBalance(ctx context.Context, userId string, balanceChange float64) (*db.UserBalance, error)
	DeleteUser(ctx context.Context, userId string) (*db.User, error)
}
