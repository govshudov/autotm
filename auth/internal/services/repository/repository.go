package repository

import (
	"context"
)

type CartServiceInterface interface {
	Test(ctx context.Context, userID int64, sku uint32, count uint16) error
}
