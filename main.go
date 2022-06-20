package main

import (
	"github.com/gofiber/fiber"
	"github.com/abisekhsubedi/golang-01/database"
)

func setupRoute(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLeads)
	app.post(NewLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase(){

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoute(app)

	// listen to port 3000
	app.Listen(3000)

	defer database.DBConn.Close()

}
