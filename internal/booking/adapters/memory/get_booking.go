package inmemory

// adapter is depending on domain but didn't violate depenedincy rule
import (
	"context"
	"errors"

	domain "github.com/e346m/bsg/internal/booking/domains"
)

func (i *InMemory) GetBooking(ctx context.Context, id string) (*domain.Booking, error) {

	var booking *domain.Booking
	for key, value := range i.db {
		if key == id {
			booking = value
			break
		}
	}

	if booking == nil {
		return nil, errors.New("no booking")
	}

	return booking, nil
}
