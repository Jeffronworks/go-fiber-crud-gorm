package routes

import (
		"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jeffronworks/fiber-api/database"
	"github.com/jeffronworks/fiber-api/models"

)

type User struct {
	// This is not the model user, this is the user serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:last_name`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
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

func GetUsers(c *fiber.Ctx) error {

	usersSlice := []models.User{} //declare an empty userSlice

	database.Database.Db.Find(&usersSlice) // get users from the db and pass into the user slice
	responseUsers := []User{}

	for _, user := range usersSlice {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id) // find the user in the database

	if user.ID == 0 {
		return errors.New("user does not exist")
	}


	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that the :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}



