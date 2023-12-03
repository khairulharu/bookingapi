package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/dto"
	"github.com/khairulharu/bookingapi/internal/util"
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
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
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
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}
