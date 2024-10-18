package profile

import (
	"context"

	"github.com/2024_2_BetterCallFirewall/internal/models"
)

type Repository interface {
	GetProfileById(uint32, context.Context) (*models.FullProfile, error)
	GetAll(self uint32, ctx context.Context) ([]*models.ShortProfile, error)
	UpdateProfile(*models.FullProfile) error
	DeleteProfile(uint32) error

	AddFriendsReq(reciever uint32, sender uint32) error
	AcceptFriendsReq(who uint32, whose uint32) error
	MoveToSubs(who uint32, whom uint32) error
	RemoveSub(who uint32, whom uint32) error
	CheckStatus(uint32, uint32, context.Context) (int, error)
	GetAllFriends(uint32, context.Context) ([]*models.ShortProfile, error)

	GetFriendsID(uint32, context.Context) ([]uint32, error)
	GetHeader(uint32) (*models.Header, error)
}

type PostGetter interface {
	GetAuthorsPosts(models.Header) ([]*models.Post, error)
}
