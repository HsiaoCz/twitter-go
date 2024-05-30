package db

import "context"

type UserStorer interface {
	CreateUser(context.Context) error
}
