package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/2024_2_BetterCallFirewall/internal/auth/controller"
	"github.com/2024_2_BetterCallFirewall/internal/models"
	"github.com/2024_2_BetterCallFirewall/internal/myErr"
	"github.com/2024_2_BetterCallFirewall/internal/profile"
)

type ProfileHandler struct {
	ProfileManager profile.ProfileUsecase
	Responder      controller.Responder
}

func NewProfileController(repo profile.ProfileUsecase, responder controller.Responder) *ProfileHandler {
	return &ProfileHandler{
		ProfileManager: repo,
		Responder:      responder,
	}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	sess, err := models.SessionFromContext(r.Context())
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}
	userId := sess.UserID
	userProfile, err := h.ProfileManager.GetProfileById(userId)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}
	h.Responder.OutputJSON(w, userProfile)
}

func (h *ProfileHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	sess, err := models.SessionFromContext(r.Context())
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}
	userId := sess.UserID
	profiles, err := h.ProfileManager.GetAll(userId)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
		return
	}
	h.Responder.OutputJSON(w, profiles)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	newProfile := models.FullProfile{}
	err := json.NewDecoder(r.Body).Decode(&newProfile)
	if err != nil {
		h.Responder.ErrorBadRequest(w, fmt.Errorf("update error:%w", err))
		return
	}

	sess, err := models.SessionFromContext(r.Context())
	if err != nil {
		h.Responder.ErrorBadRequest(w, myErr.ErrSessionNotFound)
	}
	userId := sess.UserID

	err = h.ProfileManager.UpdateProfile(userId, &newProfile)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
	}
	h.Responder.OutputJSON(w, newProfile)
}

// Фронт сам дергает еще и логаут следом за этим. Person удалим через каскад
func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	sess, err := models.SessionFromContext(r.Context())
	if err != nil {
		h.Responder.ErrorBadRequest(w, myErr.ErrSessionNotFound)
	}
	userId := sess.UserID
	err = h.ProfileManager.DeleteProfile(userId)
	if err != nil {
		h.Responder.ErrorInternal(w, err)
	}
	return
}
