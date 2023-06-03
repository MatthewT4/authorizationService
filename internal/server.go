package internal

import (
	"authorizationService/internal/HTTP"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ServerStarted() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file found")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, sslmode)

	// твой конфиг
	//config := "user=app dbname=authorizationSystemTEST password=1234 host=localhost sslmode=disable"

	// это мой конфиг)
	//config := "host=localhost user=postgres password=1234 dbname=test sslmode=disable"

	app := HTTP.NewHttp(config)

	app.Start()
}
