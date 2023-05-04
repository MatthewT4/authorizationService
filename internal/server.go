package internal

import (
	"authorizationService/internal/HTTP"
)

func ServerStarted() {
	config := "user=app dbname=authorizationSystemTEST password=1234 host=localhost sslmode=disable"

	// это мой конфиг, не удаляй пж))
	//config := "host=localhost user=postgres password=1234 dbname=test sslmode=disable"

	app := HTTP.NewHttp(config)

	app.Start()
}
