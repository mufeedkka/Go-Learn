package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mufeedkka/goecommerce/database"
	"github.com/mufeedkka/goecommerce/models"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(UserModel models.User) UserSerializer {
	return UserSerializer{ID: UserModel.ID, FirstName: UserModel.FirstName, LastName: UserModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	database.Database.Db.Find(&user)
}
