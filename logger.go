// TODO
/*
- integrate zap/zerolog library into this for log rotation 
- Database ke liye shayad alag se banega
*/

package main

import (
	"context"
	"log"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)

/*
This basically demonstrates how the logging service can be seperated from the bussiness
logic which can therefore increase maintainability.
*/
type loggingService struct {
	next BookManager
}

func NewLoggingService(svc BookManager) BookManager {
	return &loggingService{
		next: svc,
	}
}

func (l *loggingService) Search(ctx context.Context, query string) (books []types.Book, err error) {
	defer func(begin time.Time) {
		log.Printf("books: %v\terr: %v", books, err)
	}(time.Now())

	return l.next.Search(ctx, query)
}

func (l *loggingService) Filter(ctx context.Context, contraints string) (books []types.Book, err error) {
	defer func(begin time.Time) {
		log.Printf("books: %v\terr: %v\n", books, err)
	}(time.Now())

	return l.next.Filter(ctx, contraints)
}

func (l *loggingService) AddtoCart(ctx context.Context, book types.Book) (err error) {
	defer func(begin time.Time) {
		log.Printf("err: %v\n", err)
	}(time.Now())

	return l.next.AddtoCart(ctx, book)
}

func (l *loggingService) DownloadBooks(ctx context.Context, books []types.Book) (err error) {
	defer func(begin time.Time) {
		log.Printf("err: %v\n", err)
	}(time.Now())

	return l.next.DownloadBooks(ctx, books)
}

func (l *loggingService) ReviewBook(ctx context.Context, book types.Book, review types.Review) (err error) {
	defer func(begin time.Time) {
		log.Printf("err: %v\n", err)
	}(time.Now())

	return l.next.ReviewBook(ctx, book, review)
}
