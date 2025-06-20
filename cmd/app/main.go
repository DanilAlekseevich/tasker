package main

import (
	"database/sql"
	"fmt"
	"log"

	"tasker/internal/controllers"
	"tasker/internal/repositories"
	"tasker/internal/routes"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "yourpassword"
	dbname   = "yourdb"
)

func main() {
	app := fiber.New()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(err)
		}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	repositories := repositories.NewRepositories(db)
	controllers := controllers.NewControllers(repositories)
	routes.Setup(app, controllers)

	app.Listen(":3000")
}
