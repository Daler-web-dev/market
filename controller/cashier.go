package controller

import (
	"errors"
	"fmt"
	"my-fiber-app/models"
	"strconv"
	"time"

	db "my-fiber-app/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func validateCashierInput(data map[string]string, requiredFields []string) (string, bool) {
	for _, field := range requiredFields {
		if value, exists := data[field]; !exists || value == "" {
			return field + " is required!", false
		}
	}
	return "", true
}

func CahierList(c *fiber.Ctx) error {
	var cashiers []models.Cashier

	// Установка значений по умолчанию
	limit, err := strconv.Atoi(c.Query("limit", "10")) // Значение по умолчанию 10
	if err != nil || limit <= 0 {                      // Проверка на корректность и положительность
		limit = 10
	}

	skip, err := strconv.Atoi(c.Query("skip", "0")) // Значение по умолчанию 0
	if err != nil || skip < 0 {                     // Проверка на корректность и неотрицательность
		skip = 0
	}

	fmt.Printf("Fetching cashiers with limit: %d, skip: %d\n", limit, skip) // Отладочное сообщение

	result := db.DB.Limit(limit).Offset(skip).Find(&cashiers)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error fetching cashiers: " + result.Error.Error(),
		})
	}

	var count int64
	db.DB.Model(&models.Cashier{}).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashiers list",
		"data":    cashiers,
		"count":   count,
	})
}
func GetCashierDetails(c *fiber.Ctx) error {
	// Получаем cashierId из параметров
	cashierId := c.Params("cashierId")
	if cashierId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cashierId is required",
		})
	}

	var cashier models.Cashier

	// Запрос к базе данных с обработкой ошибок
	if err := db.DB.Select("id, name").Where("id = ?", cashierId).First(&cashier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
		}
		// В случае других ошибок базы данных
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Database error",
		})
	}

	// Формируем данные для ответа
	cashierData := map[string]interface{}{
		"cashierId": cashier.Id,
		"name":      cashier.Name,
		"createdAt": cashier.CreatedAt,
		"updatedAt": cashier.UpdatedAt,
	}

	// Возвращаем успешный ответ с данными кассира
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashier details",
		"data":    cashierData,
	})
}
func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false, // Изменено на false
			"message": "Invalid data",
		})
	}

	// Перечисляем обязательные поля для проверки
	requiredFields := []string{"name", "passcode"}

	// Валидация входных данных
	if message, valid := validateCashierInput(data, requiredFields); !valid {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": message,
		})
	}

	// Создание объекта кассира
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Now(), // Установлено текущее время
		UpdatedAt: time.Now(), // Установлено текущее время
	}

	// Сохранение кассира в базе данных
	if err := db.DB.Create(&cashier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create cashier: " + err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{ // Изменено на 201
		"success": true,
		"message": "Cashier added successfully",
		"data":    cashier,
	})
}
func UpdateCashier(c *fiber.Ctx) error {
	// Получаем cashierId из параметров
	cashierId := c.Params("cashierId")
	if cashierId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cashierId is required",
		})
	}

	var cashier models.Cashier

	// Запрос к базе данных с сохранением результата в dbResult
	dbResult := db.DB.Find(&cashier, "id = ?", cashierId)

	// Проверка ошибки, если она есть
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Database error",
		})
	}

	// Проверяем, найдена ли запись
	if dbResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	// Логика обновления кассира (не включена в текущий код)
	var updatedCashier models.Cashier
	err := c.BodyParser(&updatedCashier)
	if err != nil {
		return err
	}

	cashier.Name = updatedCashier.Name // пример изменения данных

	// Сохранение изменений
	if err := db.DB.Save(&cashier).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update cashier",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashier updated successfully",
		"data":    cashier,
	})
}
func EditCashier(c *fiber.Ctx) error {
	return nil
}
func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	// Проверка, что cashierId передан
	if cashierId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "CashierId is required",
		})
	}

	// Проверка на числовое значение cashierId (если это ожидается)
	if _, err := strconv.Atoi(cashierId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid cashierId format",
		})
	}

	var cashier models.Cashier
	dbResult := db.DB.Where("id = ?", cashierId).Delete(&cashier)

	// Проверка на наличие ошибок базы данных
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Database error",
		})
	}

	// Проверка, была ли запись найдена и удалена
	if dbResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	// Успешное удаление кассира
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successfully",
	})
}
