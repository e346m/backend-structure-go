package usecases

import (
	"context"

	domain "github.com/e346m/bsg/internal/booking/domains"
)

type RepositoryKeeper interface {
	Reader
	Writer
	Transactioner
}

type Reader interface {
	GetBooking(ctx context.Context, b *domain.Booking) (err error)
}

type Writer interface {
	SaveBooking(ctx context.Context, b *domain.Booking) error
}

type Transactioner interface {
	DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
