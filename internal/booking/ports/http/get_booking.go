package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestBooking struct {
	ID string `param:"id"`
}

type ResponseBooking struct {
	GuestName string `json:"guest_name"`
	BookedAt  string `json:"booked_at"`
}

// Any server is okay
// In this example, echo is used.
func (h *Handler) GetBooking(c echo.Context) error {
	var request RequestBooking
	err := c.Bind(&request)
	if err != nil {
		return err // you can convert err to http error
	}

	dom, err := h.usecase.GetBooking(
		c.Request().Context(),
		request.ID,
	)

	if err != nil {
		return err
	}

	res := ResponseBooking{
		GuestName: dom.Guest.FullName(),
		BookedAt:  dom.LocalBookedAt(),
	}

	return c.JSON(http.StatusOK, res)
}
