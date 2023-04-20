package handlers

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/mBuergi86/GO-Microservices/src/models"
	"strconv"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ReadUsers(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, _ := models.InitData()

	return c.JSON(data.GetUsers())
}

func ReadUserById(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(c.Params("id"))
	data, _ := models.InitData()

	return c.JSON(data.GetUserById(id))
}

func CreateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, _ := models.InitData()
	var user models.Users

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.CreateUser(user))
}

func UpdateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(c.Params("id"))
	var user models.Users
	data, _ := models.InitData()

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.UpdateUser(id, user))
}

func DeleteUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(c.Params("id"))
	data, _ := models.InitData()

	return c.JSON(data.DeleteUser(id))
}