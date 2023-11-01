package api

import (
	"strconv"

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
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	err := ctx.BodyParser(&reqChair)
	if err != nil {
		return err
	}
	reqChair.Id = int64(id)
	res := b.bookingservice.SaveBookingChair(ctx.Context(), reqChair)
	return ctx.Status(200).JSON(res)
}
