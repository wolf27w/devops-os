package model

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	SystemRole string `json:"system_role"` // super_admin, user
	Status     string `json:"status"`      // active, inactive
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email"`
	SystemRole string `json:"system_role"`
}

type UpdateUserRequest struct {
	Email      string `json:"email"`
	SystemRole string `json:"system_role"`
	Status     string `json:"status"`
}
