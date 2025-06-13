package services

import "simple-go/internal/models"

func GetAllUsers() []models.User {
	return []models.User{
		{
			Name: "John Doe",
			ID:   30,
		},
		{
			Name: "Jane Doe",
			ID:   28,
		},
	}
}
