package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bruce-mig/trivia-api/database"
	"github.com/bruce-mig/trivia-api/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	err := database.DB.Db.Find(&facts).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"message": "fact(s) fetched successfully",
		"data":    facts,
	})
	return nil
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := database.DB.Db.Create(&fact).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create fact"})
		return err
	}

	return c.Status(200).JSON(fact)
}

func GetFact(c *fiber.Ctx) error {
	param := c.Params("id")

	if param == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to convert to int"})
		return nil
	}
	fact := &models.Fact{}

	fmt.Println("The ID is", id)

	err = database.DB.Db.Where("id = ?", id).First(fact).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not retrieve fact"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "fact is fetched successfully",
		"data":    fact,
	})
	return nil

}

func DeleteFact(c *fiber.Ctx) error {
	param := c.Params("id")
	if param == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to convert to int"})
		return nil
	}
	fact := models.Fact{}

	err = database.DB.Db.Delete(&fact, id).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not delete fact",
		})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Fact deleted successfully",
	})
	return nil
}

func UpdateFact(c *fiber.Ctx) error {
	type updateFact struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
	// get ID params
	param := c.Params("id")
	if param == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to convert to int"})
		return nil
	}

	var fact models.Fact

	//find single fact in the database by id
	db := database.DB.Db

	err = db.Find(&fact, "id = ?", id).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "ID Not Found",
		})
	}

	var updateFactData updateFact

	err = c.BodyParser(&updateFactData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Something went wrong with your input",
			"data":    err,
		})
	}

	fact.Question = updateFactData.Question
	fact.Answer = updateFactData.Answer
	//Save the changes
	db.Save(&fact)

	//Return the updated fact
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "facts found",
		"data":    fact,
	})

}
