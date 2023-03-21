package HTTP

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Http) CreateAdminLogin(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	log.Println(user)
	w.WriteHeader(http.StatusOK)
}
