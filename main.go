package main

import (
	"github.com/gofiber/fiber"
)

func setupRoute(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLeads)
	app.post(NewLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoute(app)

	// listen to port 3000
	app.Listen(3000)

}
