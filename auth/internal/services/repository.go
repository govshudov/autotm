package services

import (
	"context"
	"net/http"
)

type CartService interface {
	Test(ctx context.Context) error
	MethodCheck(allowedMethods ...string) func(http.Handler) http.Handler
}
