package service

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/2024_2_BetterCallFirewall/internal/auth/models"
	"github.com/2024_2_BetterCallFirewall/internal/myErr"
)

type UserRepo interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type AuthServiceImpl struct {
	db UserRepo
}

func NewAuthServiceImpl(db UserRepo) *AuthServiceImpl {
	return &AuthServiceImpl{
		db: db,
	}
}

func (a *AuthServiceImpl) Register(user models.User) error {
	if !a.validateEmail(user.Email) {
		return fmt.Errorf("auth service: %w", myErr.ErrNonValidEmail)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("registration: %w", err)
	}
	user.Password = string(hashPassword)

	err = a.db.Create(&user)
	if err != nil {
		return fmt.Errorf("registration: %w", err)
	}

	return nil
}

func (a *AuthServiceImpl) Auth(user models.User) error {
	if !a.validateEmail(user.Email) {
		return fmt.Errorf("auth service: %w", myErr.ErrNonValidEmail)
	}

	dbUser, err := a.db.GetByEmail(user.Email)
	if errors.Is(err, myErr.ErrUserNotFound) {
		return fmt.Errorf("auth service: %w", myErr.ErrWrongEmailOrPassword)
	}

	if err != nil {
		return fmt.Errorf("auth service: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("auth service: %w", myErr.ErrWrongEmailOrPassword)
	}

	return nil
}

func (a *AuthServiceImpl) validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-.]+@([\w-]+\.)\w{2,4}$`)
	return emailRegex.MatchString(email)
}
