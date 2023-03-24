package internal

import (
	"authorizationService/internal/HTTP"
)

func ServerStarted() {
	config := "user=app dbname=authorizationSystemTEST password=1234 host=localhost sslmode=disable"
	app := HTTP.NewHttp(config)

	app.Start()
}
