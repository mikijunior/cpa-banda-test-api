package store

import (
	"testapi/internal/app/response"
)

type AuthRepository interface {
	KeyExists(token string) bool
}

type UserRepository interface {
	GetUsersList(username string) ([]response.Profile, error)
}