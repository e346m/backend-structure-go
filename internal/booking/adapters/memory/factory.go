package inmemory

import (
	"context"

	domain "github.com/e346m/bsg/internal/booking/domains"
)

type InMemory struct {
	db map[string]*domain.Booking
}

func NewInMemoryClient() *InMemory {
	db := make(map[string]*domain.Booking, 10)
	return &InMemory{
		db: db,
	}
}

// To abstruct transaction
func (i *InMemory) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {

	ctx = context.TODO()

	v, err := f(ctx)
	if err != nil {
		// similar action like delete request to client
		return nil, err
	}

	return v, nil
}
