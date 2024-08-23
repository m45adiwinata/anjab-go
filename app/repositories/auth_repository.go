package repositories

import (
	"context"
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
)

type AuthRepository interface {
	Login(ctx context.Context, email string) (*models.User, error)
	Register(ctx context.Context, data *models.User) error
}

type authRepo struct {
	db orm.Orm
}

func NewAuthRepository(db orm.Orm) AuthRepository {
	return &authRepo{db: db}
}

func (r *authRepo) Login(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}
	err := r.db.WithContext(ctx).Query().Select("email", "password", "fullname", "nickname", "skpd_id").Where("email", email).First(&user)
	if user.Email == "" {
		return nil, fmt.Errorf("user not found")
	}
	return &user, err
}

func (r *authRepo) Register(ctx context.Context, data *models.User) error {
	return r.db.WithContext(ctx).Query().Create(data)
}
