package domain

import (
	"context"

	"github.com/khairulharu/bookingapi/dto"
)

type BookingService interface {
	GetBookingChairs(ctx context.Context) dto.Response
	SaveBookingChair(ctx context.Context, chair dto.ReqChair) dto.Response
	GetBookingChair(ctx context.Context, id int64) dto.Response
}
