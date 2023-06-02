package HTTP

import (
	AdminBLogic "authorizationService/internal/BLogic/admin"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Http struct {
	adminBLogic AdminBLogic.IAdmBLogic
}

func NewHttp(config string) *Http {
	return &Http{adminBLogic: AdminBLogic.NewAdmBLogic(config)}
}

func (h *Http) Start() {
	app := fiber.New()

	v1 := app.Group("/v1")
	admin := v1.Group("/admin")

	admin.Post("/signup", h.userSignUp)
	admin.Post("/login", h.userSignIn)
	admin.Get("/validate", h.requireAuth, h.validate)

	log.Fatal(app.Listen(":80"))
}
