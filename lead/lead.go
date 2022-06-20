package lead

import (
	"github.com/abisekhsubedi/golang-01/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// struct: A datatype I am creating on my own
type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func Getlead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)

}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendString("Error parsing body")
		return
	}

	db.Create(&lead)
	c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead deleted")
}
