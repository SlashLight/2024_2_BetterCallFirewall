package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/2024_2_BetterCallFirewall/internal/models"
	"github.com/2024_2_BetterCallFirewall/pkg/my_err"
)

type UserRepo interface {
	Create(user *models.User, ctx context.Context) (uint32, error)
	GetByEmail(email string, ctx context.Context) (*models.User, error)
}

type AuthServiceImpl struct {
	db UserRepo
}

func NewAuthServiceImpl(db UserRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		db: db,
	}
}

func (a *AuthServiceImpl) Register(user models.User, ctx context.Context) (uint32, error) {
	if !a.validateEmail(user.Email) {
		return 0, fmt.Errorf("auth service: %w", my_err.ErrNonValidEmail)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("registration: %w", err)
	}
	user.Password = string(hashPassword)

	user.ID, err = a.db.Create(&user, ctx)
	if err != nil {
		return 0, fmt.Errorf("registration: %w", err)
	}

	return user.ID, nil
}

func (a *AuthServiceImpl) Auth(user models.User, ctx context.Context) (uint32, error) {
	if !a.validateEmail(user.Email) {
		return 0, fmt.Errorf("auth service: %w", my_err.ErrNonValidEmail)
	}

	dbUser, err := a.db.GetByEmail(user.Email, ctx)
	if errors.Is(err, my_err.ErrUserNotFound) {
		return 0, fmt.Errorf("auth service: %w", my_err.ErrWrongEmailOrPassword)
	}

	if err != nil {
		return 0, fmt.Errorf("auth service: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return 0, fmt.Errorf("auth service: %w", my_err.ErrWrongEmailOrPassword)
	}

	return dbUser.ID, nil
}

func (a *AuthServiceImpl) validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-.]+@([\w-]+\.)\w{2,4}$`)
	return emailRegex.MatchString(email)
}
