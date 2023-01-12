package models

type SignUpBody struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	Username string `form:"username" binding:"required"`
	Phone    string `form:"phone"`
	Role     string `form:"role"`
}

type LoginBody struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type User struct {
	Id        string `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Username  string `db:"username"`
	Phone     string `db:"phone"`
	Role      string `db:"role"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	IsActive  bool   `db:"is_active"`
}

type LoginResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
	IsActive    bool   `json:"isActive"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ReadBook struct {
	Id         string `json:"id" db:"id"`
	UserId     string `json:"userId" db:"user_id" binding:"required"`
	Title      string `json:"title" db:"title" binding:"required"`
	Author     string `json:"author" db:"author" binding:"required"`
	Notes      string `json:"notes" db:"notes"`
	IsFinished bool   `json:"isFinished" db:"is_finished"`
}

type ReadBookInput struct {
	Title      string `form:"title" binding:"required"`
	Author     string `form:"author" binding:"required"`
	Notes      string `form:"notes"`
	IsFinished bool   `form:"isFinished"`
}
