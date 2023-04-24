package HTTP

import (
	AdminBLogic "authorizationService/internal/BLogic/admin"
	"github.com/gofiber/fiber/v2"
)

type Http struct {
	//blogic      *BLogic.B
	adminBLogic AdminBLogic.IAdmBLogic
}

func NewHttp(config string) *Http {
	return &Http{adminBLogic: AdminBLogic.NewAdmBLogic(config)}
}

func (h *Http) Start() {
	app := fiber.New()

	v1 := app.Group("/v1")
	admin := v1.Group("/admin")

	admin.Get("/login", h.CreateAdminLogin)

	app.Listen(":80")
}
