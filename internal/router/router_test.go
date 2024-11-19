package router

import (
	"context"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/2024_2_BetterCallFirewall/internal/models"
)

type mockProfileController struct{}

func (m mockProfileController) GetProfile(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetProfileById(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetAll(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) UpdateProfile(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) DeleteProfile(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetHeader(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) SendFriendReq(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) AcceptFriendReq(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) Unsubscribe(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) RemoveFromFriends(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetAllFriends(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetAllSubs(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {}

func (m mockProfileController) GetCommunitySubs(w http.ResponseWriter, r *http.Request) {}

type mockPostController struct{}

func (m mockPostController) Create(w http.ResponseWriter, r *http.Request) {}

func (m mockPostController) GetOne(w http.ResponseWriter, r *http.Request) {}

func (m mockPostController) Update(w http.ResponseWriter, r *http.Request) {}

func (m mockPostController) Delete(w http.ResponseWriter, r *http.Request) {}

func (m mockPostController) GetBatchPosts(w http.ResponseWriter, r *http.Request) {}

type mockMiddleware struct{}

func (m mockMiddleware) Check(str string) (*models.Session, error) { return nil, nil }

func (m mockMiddleware) Create(userID uint32) (*models.Session, error) { return nil, nil }

func (m mockMiddleware) Destroy(sess *models.Session) error { return nil }

type mockChatController struct{}

func (m mockChatController) SetConnection(w http.ResponseWriter, r *http.Request) {}

func (m mockChatController) GetAllChats(w http.ResponseWriter, r *http.Request) {}

func (m mockChatController) GetChat(w http.ResponseWriter, r *http.Request) {}

func (m mockChatController) SendChatMsg(ctx context.Context, reqID string) {}

type mockCommunityController struct{}

func (m mockCommunityController) GetOne(w http.ResponseWriter, r *http.Request) {}

func (m mockCommunityController) GetAll(w http.ResponseWriter, r *http.Request) {}

func (m mockCommunityController) Update(w http.ResponseWriter, r *http.Request) {}

func (m mockCommunityController) Delete(w http.ResponseWriter, r *http.Request) {}

func (m mockCommunityController) Create(w http.ResponseWriter, r *http.Request) {}

func TestNewRouter(t *testing.T) {
	router := NewRouter(mockProfileController{},
		mockPostController{},
		mockMiddleware{},
		mockChatController{},
		mockCommunityController{},
		logrus.New(),
	)
	assert.NotNil(t, router)
}
