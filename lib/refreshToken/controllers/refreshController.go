package controllers

import (
	"time"
	"github.com/asciiu/appa/lib/refreshToken/models"
)

type RefreshController struct {
	tokenRepo models.TokenRepo
}

func NewRefreshController(tokenRepo models.TokenRepo) *RefreshController {
	return &RefreshController{tokenRepo: tokenRepo}
}

func (controller *RefreshController) CreateRefreshToken(userID string, expiresOn time.Time) (*models.RefreshToken, string, error) {
	refreshToken := models.NewRefreshToken(userID)
	selectAuth := refreshToken.Renew(expiresOn)

	token, err := controller.tokenRepo.InsertRefreshToken(refreshToken)
	if err != nil {
		return nil, "", err
	}

	return token, selectAuth, nil
}