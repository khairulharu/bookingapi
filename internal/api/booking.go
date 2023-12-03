package api

import (
	"strconv"

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
	app.Get("/booking/chair/", h.GetBookingChair)
}

func (b apiBooking) ShowAllChairs(ctx *fiber.Ctx) error {
	res := b.bookingservice.GetBookingChairs(ctx.Context())
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
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
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}

func (b apiBooking) GetBookingChair(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := b.bookingservice.GetBookingChair(ctx.Context(), int64(id))
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}
