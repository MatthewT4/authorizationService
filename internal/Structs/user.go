package Structs

type User struct {
	UserId       int    `json:"user_id"`
	Email        string `json:"email"`
	HashPassword string `json:"password"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Phone        string `json:"phone"`
}

type UserSignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
}
