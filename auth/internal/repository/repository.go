package repository

import (
	"context"
	"errors"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type CartRepository interface {
	Test(context.Context) error
}
