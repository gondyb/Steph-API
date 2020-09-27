package repositories

import (
	"../config"
	"../models"
)

// GetAllUsers returns the list of every user in the database
func GetAllUsers(u *[]models.User) (err error) {
	if err = config.DB.Find(u).Error; err != nil {
		return err
	}

	return nil
}
