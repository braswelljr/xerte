package user

import "time"

type User struct {
	ID          string    `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Othernames  string    `json:"othernames"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Phone       string    `json:"phone"`
	Roles       []string  `json:"roles"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserResponse struct {
	ID          string    `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Othernames  string    `json:"othernames"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Phone       string    `json:"phone"`
	Roles       []string  `json:"roles"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
