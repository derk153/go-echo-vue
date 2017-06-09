package handlers

import (
	"database/sql"
	"go-echo-vue/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// H map JSON string
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		c.Bind(&task)
		id, err := models.PutTask(db, task.Name)

		if err == nil {
			return c.JSON(http.StatusCreated, H{"created": id})
		}

		return err
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTasks(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{"deleted": id})
		}

		return err
	}
}
