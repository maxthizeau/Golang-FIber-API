package helpers

import (
	"github.com/maxthizeau/api-fiber/controllers"
	"github.com/maxthizeau/api-fiber/models"
)

func PopulateOrder(order *models.Order, userId int, productId int) error {
	var user models.User
	if err := controllers.FindUser(int(order.UserRefer), &user); err != nil {
		return err
	}

	var product models.Product
	if err := controllers.FindProduct(int(order.ProductRefer), &product); err != nil {
		return err
	}

	order.Product = product
	order.User = user

	return nil
}
