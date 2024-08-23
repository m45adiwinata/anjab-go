package repositories

import (
	"context"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}
type userRepo struct {
	db orm.Orm
}

func NewUserRepository(db orm.Orm) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	result := []models.User{}
	if err := r.db.WithContext(ctx).Query().Get(&result); err != nil {
		return nil, err
	}

	return result, nil
}
