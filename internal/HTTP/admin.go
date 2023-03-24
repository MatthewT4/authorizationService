package HTTP

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Http) CreateAdminLogin(ctx *fiber.Ctx) error {
	if len(ctx.Params("login")) == 0 {
		ctx.Status(http.StatusBadRequest)
		var err struct {
			Message string `json:"message"`
		}
		err.Message = "" // write later...
		mes, _ := json.Marshal(&err)
		ctx.Write(mes)
		return nil
	}

	if len(ctx.Params("password")) == 0 {
		ctx.Status(http.StatusBadRequest)
		var err struct {
			Message string `json:"message"`
		}
		err.Message = "password" // write later...
		mes, _ := json.Marshal(&err)
		ctx.Write(mes)
		return nil
	}

	login := ctx.Get("login")

	password := ctx.Get("password")

	// Blogic.login()

	// return JWT
	fmt.Println(login, password)
	ctx.Write([]byte("OK"))

	return nil
}
