package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username  *string        `gorm:"null" json:"username"`
	Role      string         `gorm:"not null" json:"role"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CreateUser struct {
	Username *string `json:"username"`
	Role     string  `json:"role"`
	Password string  `json:"password"`
}

type TokenClaims struct {
	User *User `json:"user"`
	jwt.StandardClaims
}

type LoginPayload struct {
	Username *string `json:"username"`
	Password string  `json:"password"`
}

type UserRepository interface {
	RetrieveAllUser() ([]User, error)
	RetrieveByUsername(username string) (*User, error)
	RetrieveUserByID(id uint) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id uint) error
}

type UserUseCase interface {
	FetchUsers(ctx context.Context) ([]User, error)
	LoginUser(ctx context.Context, req *LoginPayload) (*User, string, error)
	FetchUserByID(ctx context.Context, id uint) (*User, error)
	CreateUser(ctx context.Context, req *User) (*User, error)
	UpdateUser(ctx context.Context, req *User) (*User, error)
	DeleteUser(ctx context.Context, id uint) error
}
