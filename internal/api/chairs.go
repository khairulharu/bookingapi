package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/bookingapi/domain"
	"github.com/khairulharu/bookingapi/internal/util"
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
	app.Get("/chair", h.SearchChair)
}

func (a apiChair) ChairStore(ctx *fiber.Ctx) error {
	idString := ctx.Query("id")
	id, _ := strconv.Atoi(idString)
	res := a.chairService.StoreChairs(ctx.Context(), id)
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}

func (a apiChair) ChairsDeleting(ctx *fiber.Ctx) error {
	res := a.chairService.DeleteChairs(ctx.Context())
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}

func (a apiChair) SearchChair(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	res := a.chairService.SearchChairs(ctx.Context(), code)
	return ctx.Status(util.GetHttpCode(res.Code)).JSON(res)
}
