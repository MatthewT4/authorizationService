package Structs

import (
	"math/rand"
)

type Service struct {
	ServiceId   int
	OwnerId     int
	Name        string
	Domains     []string
	Discription string
	privateKey  string
}

func (s *Service) PrivateKey() string {
	return s.privateKey
}

func (s *Service) GeneratePrivateKey() {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	s.privateKey = string(b)
}
