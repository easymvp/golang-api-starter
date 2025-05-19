package users

import (
	"context"
	"easymvp_api/internal/users/entries"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/xerror"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	fmt.Println("db", db)
	return &UserService{db: db}
}

func (s *UserService) Get(ctx context.Context, id string) (*entries.User, error) {
	var user *entries.User
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.New("missing user")
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) Save(ctx context.Context, user *entries.User) error {
	if user == nil {
		return xerror.New("user is nil")
	}
	fmt.Println("==========+>", s.db)
	return s.db.WithContext(ctx).Save(user).Error
}
