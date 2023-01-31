package user

import (
	domain "github.com/Doittikorn/school/internal/user/domain"
)

type UserRepository interface {
	FindAll() ([]*domain.UserInfo, error)
	FindByID(id int) (*domain.UserInfo, error)
	Store(u *domain.User) error
	Update(u *domain.User) error
	Delete(u *domain.UserInfo) error
}
