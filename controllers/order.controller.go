package controllers

import (
	"errors"

	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/models"
)

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(order, "id = ?", id)

	if order.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}
