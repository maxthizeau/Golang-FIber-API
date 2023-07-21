package controllers

import (
	"errors"

	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/models"
)

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}
