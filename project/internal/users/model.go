package users

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	IsSuperuser bool      `json:"is_superuser"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
