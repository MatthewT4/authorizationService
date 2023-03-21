package HTTP

import (
	"authorizationService/internal/BLogic"
	"github.com/gorilla/mux"
)

type Http struct {
	blogic *BLogic.BLogic
}

func NewHttp(config string) *Http {
	return &Http{blogic: BLogic.NewBLogic(config)}
}

func (h *Http) Start() {
	router := mux.NewRouter()
	routerAdmin := router.PathPrefix("/admin").Subrouter()
	routerAdmin.HandleFunc("/create_user", h.CreateAdminLogin)
}
