package bookService

import (
	"context"
	"os"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/rs/zerolog"
)

type loggingService struct {
	next   Manager
	logger zerolog.Logger
}

// DeleteTag implements Manager.
func NewLoggingService(svc Manager) Manager {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &loggingService{
		next:   svc,
		logger: logger,
	}
}

func (l *loggingService) Search(ctx context.Context, req types.SearchBooksV1Request) (res []database.SearchBooksV2Row, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Str("query", req.WebsearchToTsquery).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Str("query", req.WebsearchToTsquery).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.Search(ctx, req)
}

func (l *loggingService) GetHotSelling(ctx context.Context, req types.GetHotSellingRequest) (res []database.GetHotSellingBooksRow, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetHotSelling(ctx, req)
}

func (l *loggingService) GetPersonalRecommendations(ctx context.Context, req types.GetPersonalRecommendationsRequest) (res []database.SearchBooksV2Row, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetPersonalRecommendations(ctx, req)
}

func (l *loggingService) GetBook(ctx context.Context, req types.GetBookRequest) (res database.GetBookByIdRow, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetBook(ctx, req)
}

func (l *loggingService) AddBook(ctx context.Context, req types.AddBookRequest) (res database.Book, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.AddBook(ctx, req)
}

func (l *loggingService) UpdateBook(ctx context.Context, req types.UpdateBookRequest) (res database.Book, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.UpdateBook(ctx, req)
}

func (l *loggingService) GetAllTags(ctx context.Context) (res []database.Tag, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetAllTags(ctx)
}

func (l *loggingService) CreateTag(ctx context.Context, req types.CreateTagRequest) (res database.Tag, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.CreateTag(ctx, req)
}

func (l *loggingService) UpdateTag(ctx context.Context, req types.UpdateTagRequest) (res database.Tag, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.UpdateTag(ctx, req)
}


func (l *loggingService) DeleteTag(ctx context.Context, req types.DeleteTagRequest) (err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id (admin)", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.DeleteTag(ctx, req)
}


func (l *loggingService) GetCart(ctx context.Context, req types.GetCartRequest) (res []database.GetCartItemsByUserIdRow, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetCart(ctx, req)
}

func (l *loggingService) AddToCart(ctx context.Context, req types.AddToCartRequest) (res database.Cart, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.AddToCart(ctx, req)
}

func (l *loggingService) DeleteCartItem(ctx context.Context, req types.DeleteCartItemRequest) (err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.DeleteCartItem(ctx, req)
}

func (l *loggingService) PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (res database.Order, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.PlaceOrder(ctx, req)
}

func (l *loggingService) VerifyOrder(ctx context.Context, req types.VerifyOrderRequest) (res database.UpdateOrderTxResult, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.VerifyOrder(ctx, req)
}

func (l *loggingService) GetPurchases(ctx context.Context, req types.GetPurchasesRequest) (res []database.GetPurchasedBooksRow, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetPurchases(ctx, req)
}

func (l *loggingService) GetReviews(ctx context.Context, req types.GetReviewsRequest) (res []database.Review, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.GetReviews(ctx, req)
}

func (l *loggingService) AddReview(ctx context.Context, req types.AddReviewRequest) (res database.Review, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.AddReview(ctx, req)
}

func (l *loggingService) UpdateReview(ctx context.Context, req types.UpdateReviewRequest) (resp database.Review, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.UpdateReview(ctx, req)
}

func (l *loggingService) DeleteReview(ctx context.Context, req types.DeleteReviewRequest) (err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Error().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Str("err", err.Error()).
				Dur("took", time.Since(begin)).
				Send()
		} else {
			l.logger.Info().
				Str("addr", ctx.Value(types.RemoteAddress).(string)).
				Int64("id", req.UserID).
				Dur("took", time.Since(begin)).
				Send()
		}
	}(time.Now())

	return l.next.DeleteReview(ctx, req)
}
