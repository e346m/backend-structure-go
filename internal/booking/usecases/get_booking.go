package usecases

import (
	"context"

	domain "github.com/e346m/bsg/internal/booking/domains"
)

func (u *Usecase) GetBooking(ctx context.Context, id string) (dom *domain.Booking, err error) {

	dom, err = u.repo.GetBooking(ctx, id)

	if err != nil {
		return
	}

	return
}
