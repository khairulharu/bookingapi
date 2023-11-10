package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
)

type apiBooking struct {
	bookingservice domain.BookingService
}

func NewBooking(app *fiber.App, bookingService domain.BookingService) {
	h := apiBooking{
		bookingservice: bookingService,
	}

	app.Get("/booking", h.ShowAllChairs)
	app.Post("/booking", h.StoreBooking)
}

func (b apiBooking) ShowAllChairs(ctx *fiber.Ctx) error {
	res := b.bookingservice.GetBookingChairs(ctx.Context())
	return ctx.Status(200).JSON(res)
}

func (b apiBooking) StoreBooking(ctx *fiber.Ctx) error {
	var reqChair dto.ReqChair
	code := ctx.Query("code")
	err := ctx.BodyParser(&reqChair)
	if err != nil {
		return err
	}
	reqChair.CodeRef = code
	res := b.bookingservice.SaveBookingChair(ctx.Context(), reqChair)
	return ctx.Status(200).JSON(res)
}
