package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Login struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}
