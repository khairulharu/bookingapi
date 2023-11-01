package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/bookingapi/domain"
)

type apiChair struct {
	chairService domain.ChairService
}

func NewChair(app *fiber.App, chairService domain.ChairService) {
	h := apiChair{
		chairService: chairService,
	}

	app.Post("/chairs", h.ChairStore)
	app.Delete("/chairs", h.ChairsDeleting)
}

func (a apiChair) ChairStore(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := a.chairService.StoreChairs(ctx.Context(), id)
	return ctx.Status(200).JSON(res)
}

func (a apiChair) ChairsDeleting(ctx *fiber.Ctx) error {
	res := a.chairService.DeleteChairs(ctx.Context())
	return ctx.Status(200).JSON(res)
}
