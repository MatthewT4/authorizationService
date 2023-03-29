package HTTP

import (
	AdminBLogic "authorizationService/internal/BLogic/admin"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Cookie struct {
	Value string `json:"value"`
}

func (h *Http) CreateAdminLogin(ctx *fiber.Ctx) error {
	// check login
	if len(ctx.Query("login")) == 0 {
		ctx.Status(http.StatusBadRequest)
		var err struct {
			Message string `json:"message"`
		}
		err.Message = "incorrect login" // write later...
		mes, _ := json.Marshal(&err)
		ctx.Write(mes)
		return nil
	}

	login := ctx.Query("login")

	//check password
	if len(ctx.Query("password")) == 0 {
		ctx.Status(http.StatusBadRequest)
		var err struct {
			Message string `json:"message"`
		}
		err.Message = "incorrect password" // write later...
		mes, _ := json.Marshal(&err)
		ctx.Write(mes)
		return nil
	}

	password := ctx.Query("password")

	// Blogic.login()
	httpStatus, jwtToken := AdminBLogic.CheckValid(login, password)
	ctx.SendStatus(httpStatus)

	// return JWT
	ctx.Write([]byte("OK"))
	ctx.Cookie(&fiber.Cookie{
		Value: jwtToken,
	})

	return nil
}
