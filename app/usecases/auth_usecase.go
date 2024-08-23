package usecases

import (
	"context"
	"fmt"
	"goravel/app/models"
	"goravel/app/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/facades"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (*models.User, string, error)
	Register(ctx context.Context, user *models.User) (*models.User, string, error)
}

type authUc struct {
	config config.Config
	repo   repositories.AuthRepository
}

func NewAuthUsecase(repo repositories.AuthRepository) AuthUsecase {
	config := facades.Config()
	return &authUc{repo: repo, config: config}
}

func (uc *authUc) generateToken(user *models.User) (string, error) {
	jwtSecret := []byte(uc.config.Env("JWT_SECRET").(string))
	claims := models.CustomClaim{
		Email:    user.Email,
		FullName: user.Fullname,
		NickName: user.Nickname,
		SkpdId:   user.SkpdId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uc *authUc) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	user, err := uc.repo.Login(ctx, email)
	if err != nil {
		return nil, "", err
	}
	pass, err := facades.Crypt().DecryptString(*user.Password)
	if err != nil {
		return nil, "", err
	}
	if pass != password {
		return nil, "", fmt.Errorf("invalid username or password")
	}

	token, err := uc.generateToken(user)
	if err != nil {
		return nil, "", err
	}
	user.Password = nil

	return user, token, nil
}

func (uc *authUc) Register(ctx context.Context, user *models.User) (*models.User, string, error) {
	user.CreatedAt = func() *time.Time { t := time.Now(); return &t }()
	user.UpdatedAt = user.CreatedAt
	password, err := facades.Crypt().EncryptString(*user.Password)
	if err != nil {
		return nil, "", err
	}
	user.Password = &password

	err = uc.repo.Register(ctx, user)
	if err != nil {
		return nil, "", err
	}

	token, err := uc.generateToken(user)
	if err != nil {
		return nil, "", err
	}
	user.Password = nil

	return user, token, nil
}
