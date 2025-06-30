package services

import (
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
)

type CartServiceImpl struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) *CartServiceImpl {
	return &CartServiceImpl{
		repo: repo,
	}
}

func (s *CartServiceImpl) Test(ctx context.Context) error {
	fmt.Println("Test")
	return s.repo.Test(ctx)
}
func (s *CartServiceImpl) MethodCheck(allowedMethods ...string) func(http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(allowedMethods))
	for _, m := range allowedMethods {
		allowed[m] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := allowed[r.Method]; !ok {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
