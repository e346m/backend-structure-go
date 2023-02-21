package psql

// adapter is depending on domain but didn't violate depenedincy rule
import (
	"context"

	domain "github.com/e346m/bsg/internal/booking/domains"
)

func (p *PSQL) GetBooking(ctx context.Context, id string) (*domain.Booking, error) {

	// select * from guest where id = id by p.db

	return &domain.Booking{
		ID: id,
		// Guest
		// createdAt
		// updatedAt
	}, nil
}
