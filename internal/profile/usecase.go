package profile

import (
	"context"

	"github.com/2024_2_BetterCallFirewall/internal/models"
)

type ProfileUsecase interface {
	GetProfileById(context.Context, uint32) (*models.FullProfile, error)
	GetAll(ctx context.Context, self uint32) ([]*models.ShortProfile, error)
	UpdateProfile(uint32, *models.FullProfile) error
	DeleteProfile(uint32) error

	SendFriendReq(receiver uint32, sender uint32) error
	AcceptFriendReq(who uint32, whose uint32) error
	RemoveFromFriends(who uint32, whose uint32) error
	GetAllFriends(ctx context.Context, self uint32) ([]*models.ShortProfile, error)
}
