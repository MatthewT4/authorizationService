package HTTP

import (
	"authorizationService/internal/Structs"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type Cookie struct {
	Value string `json:"value"`
}

func (h *Http) userSignUp(ctx *fiber.Ctx) error {
	inputUser := new(Structs.UserSignUpInput)

	if err := ctx.BodyParser(inputUser); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusBadRequest, "bad request, please check params")
	}

	msg, err := h.adminBLogic.SignUp(inputUser)
	if err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusBadRequest, msg)
	}

	return ctx.SendString("successful signup")
}

func (h *Http) userSignIn(ctx *fiber.Ctx) error {
	inputUser := new(Structs.UserSignInInput)

	if err := ctx.BodyParser(inputUser); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusBadRequest, "bad request, please check params")
	}

	msg, err := h.adminBLogic.SignIn(inputUser)
	if err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusBadRequest, msg)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "Authorization",
		Value: msg,
	})
	return ctx.SendString("successful login")
}

func (h *Http) requireAuth(ctx *fiber.Ctx) error {
	return nil
}

func (h *Http) validate(ctx *fiber.Ctx) error {
	return ctx.SendString("successful validate")
}

//func (h *Http) CreateAdminLogin(ctx *fiber.Ctx) error {
//	// check login
//	if len(ctx.Query("login")) == 0 {
//		ctx.Status(http.StatusBadRequest)
//		var err struct {
//			Message string `json:"message"`
//		}
//		err.Message = "incorrect login" // write later...
//		mes, _ := json.Marshal(&err)
//		ctx.Write(mes)
//		return nil
//	}
//	login := ctx.Query("login")
//
//	//check password
//	if len(ctx.Query("password")) == 0 {
//		ctx.Status(http.StatusBadRequest)
//		var err struct {
//			Message string `json:"message"`
//		}
//		err.Message = "incorrect password" // write later...
//		mes, _ := json.Marshal(&err)
//		ctx.Write(mes)
//		return nil
//	}
//	password := ctx.Query("password")
//
//	// Blogic.login()
//	httpStatus, jwtToken := AdminBLogic.CheckValid(login, password)
//	ctx.SendStatus(httpStatus)
//
//	// return JWT
//	ctx.Cookie(&fiber.Cookie{
//		Value: jwtToken,
//	})
//
//	return nil
//}
