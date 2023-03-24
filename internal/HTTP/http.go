package HTTP

import (
	"authorizationService/internal/BLogic"
	"github.com/gofiber/fiber/v2"
)

type Http struct {
	blogic *BLogic.BLogic
}

func NewHttp(config string) *Http {
	return &Http{blogic: BLogic.NewBLogic(config)}
}

func (h *Http) Start() {
	app := fiber.New()

	v1 := app.Group("/v1")
	admin := v1.Group("/admin")

	admin.Get("/login", h.CreateAdminLogin)

	app.Listen(":80")
}
