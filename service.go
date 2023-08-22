package main

import (
	"context"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal.git/types"
)

type BookManager interface {
	Search(context.Context, string) ([]types.Book, error)
	Filter(context.Context, string) ([]types.Book, error)
	AddtoCart(context.Context, types.Book) error
	DownloadBooks(context.Context, []types.Book) error
	ReviewBook(context.Context, types.Book, types.Review) error
}

type bookManager struct{
}

func NewBookManager() BookManager {
	return &bookManager{}
}

func (b *bookManager) Search(ctx context.Context, query string) ([]types.Book, error){
	return nil, nil
}

func (b *bookManager) Filter(ctx context.Context, contraints string) ([]types.Book, error) {
	return nil, nil
}

func (b *bookManager) AddtoCart(ctx context.Context, book types.Book) error {
	return nil
}

func (b *bookManager) DownloadBooks(ctx context.Context, books []types.Book) error {
	return nil
}

func (b *bookManager) ReviewBook(ctx context.Context, book types.Book, review types.Review) error {
	return nil
}

