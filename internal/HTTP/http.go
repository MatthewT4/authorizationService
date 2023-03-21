package HTTP

import "authorizationService/internal/BLogic"

type Http struct {
	blogic *BLogic.BLogic
}

func NewHttp(config string) *Http {
	return &Http{blogic: BLogic.NewBLogic(config)}
}
