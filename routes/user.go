package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/models"
)

type User struct {
	// This is NOT the model user, see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// We could watch for validation libraries here

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted user")
}
