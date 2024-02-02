package response

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CratedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
