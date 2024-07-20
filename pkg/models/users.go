package models

type User struct {
	ID         int64  `json:"id"`
	RoleId     int64  `json:"role_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	LastAccess string `json:"last_access"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
