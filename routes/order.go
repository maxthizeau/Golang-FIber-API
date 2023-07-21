package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/api-fiber/controllers"
	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/helpers"
	"github.com/maxthizeau/api-fiber/models"
)

type Order struct {
	ID        uint      `json:"id"`
	Product   Product   `json:"product"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		Product:   product,
		User:      user,
		CreatedAt: order.CreatedAt,
	}
}

func CreateFullOrderResponse(order models.Order) Order {
	responseUser := CreateResponseUser(order.User)
	responseProduct := CreateProductResponse(order.Product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return responseOrder
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(int(order.UserRefer), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(int(order.ProductRefer), &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	createFullOrderResponse := CreateFullOrderResponse(order)

	return c.Status(fiber.StatusOK).JSON(createFullOrderResponse)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	queries := new(struct {
		UserID int `query:"user_id"`
	})

	if err := c.QueryParser(queries); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// database.Database.Db.Find(&orders, "user_id = ?", queries.userID)

	database.Database.Db.Where(&models.User{ID: uint(queries.UserID)}).Find(&orders)
	responseOrders := []Order{}

	for _, o := range orders {

		if err := helpers.PopulateOrder(&o, int(o.UserRefer), int(o.ProductRefer)); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responseOrders = append(responseOrders, CreateFullOrderResponse(o))
	}

	return c.Status(fiber.StatusOK).JSON(responseOrders)
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := controllers.FindOrder(id, &order); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if err := helpers.PopulateOrder(&order, int(order.UserRefer), int(order.ProductRefer)); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseOrder := CreateFullOrderResponse(order)
	return c.Status(fiber.StatusOK).JSON(responseOrder)
}

// func GetOrdersByUser(c *fiber.Ctx) error {

// 	orders := []models.Order{}
// 	database.Database.Db.Find(&orders, "user_id = ?", queries.userID)
// }
