package services

import (
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
)

type CartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) Test(ctx context.Context) error {
	fmt.Println("Test")
	return s.repo.Test(ctx)
}
