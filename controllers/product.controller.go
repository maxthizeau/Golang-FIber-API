package controllers

import (
	"errors"

	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/models"
)

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("product does not exist")
	}

	return nil
}
