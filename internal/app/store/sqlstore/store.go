package sqlstore

import (
	"database/sql"
	"testapi/internal/app/store"
)

type Store struct {
	db 				*sql.DB
	userRepository  *UserRepository
	authRepository  *AuthRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Auth() store.AuthRepository {
	if s.userRepository != nil {
		return s.authRepository
	}

	s.authRepository = &AuthRepository{
		store: s,
	}

	return s.authRepository
}
