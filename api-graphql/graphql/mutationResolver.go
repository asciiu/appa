package graphql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/asciiu/appa/api-graphql/auth"
	repo "github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	constRes "github.com/asciiu/appa/lib/constants/response"
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	token "github.com/asciiu/appa/lib/refreshToken/models"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	protoStory "github.com/asciiu/appa/story-service/proto/story"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email, username, password string) (*user.User, error) {
	newUser := user.NewUser(username, email, password)
	if err := userRepo.InsertUser(r.DB, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *mutationResolver) Login(ctx context.Context, email, password string, remember bool) (*Token, error) {
	log.Info(fmt.Sprintf("login: %s", email))

	loginUser, err := userRepo.FindUserByEmail(r.DB, email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("incorrect password/email")
	case err != nil:
		return nil, err
	case !loginUser.EmailVerified:
		// only verified accounts should be able to login
		return nil, fmt.Errorf("email account not verified")
	default:
		if bcrypt.CompareHashAndPassword([]byte(loginUser.PasswordHash), []byte(password)) == nil {
			jwt, err := auth.CreateJwtToken(loginUser.ID, auth.JwtDuration)
			if err != nil {
				return nil, err
			}

			// issue a refresh token if remember is true
			if remember {
				refreshToken := token.NewRefreshToken(loginUser.ID)
				expiresOn := time.Now().Add(auth.RefreshDuration)
				selectAuth := refreshToken.Renew(expiresOn)

				// this needs to be checked
				if _, err := tokenRepo.InsertRefreshToken(r.DB, refreshToken); err != nil {
					log.Error(fmt.Sprintf("failed to insert refresh token: %s", err.Error()))
				}

				return &Token{
					Jwt:     &jwt,
					Refresh: &selectAuth,
				}, nil
			}

			return &Token{Jwt: &jwt}, nil
		}

		return nil, fmt.Errorf("incorrect password/email")
	}
}

func (r *mutationResolver) CreateStory(ctx context.Context, title, jsonData string) (string, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return "", fmt.Errorf("unauthorized")
	}

	if r.StoryClient == nil {
		return "", fmt.Errorf("story service client was not intantiated")
	}

	story := models.NewStory(loginUser.ID, title, jsonData)
	if err := repo.InsertStory(r.DB, story); err != nil {
		return "", err
	}

	req := protoStory.InitStoryRequest{
		StoryID:   story.ID,
		UserID:    loginUser.ID,
		Username:  loginUser.Username,
		UserEmail: loginUser.Email,
		Title:     title,
		JsonData:  jsonData,
	}
	res, err := r.StoryClient.InitStory(ctx, &req)
	fmt.Println(err)
	fmt.Println(res)

	if res.Status != constRes.Success {
		//res := &ResponseError{
		//	Status:  r.Status,
		//	Message: r.Message,
		//}

		//if r.Status == constRes.Fail {
		//	return c.JSON(http.StatusBadRequest, res)
		//}
		//if r.Status == constRes.Error {
		//	return c.JSON(http.StatusInternalServerError, res)
		//}
		return "", errors.New(res.Message)
	}

	return story.ID, nil
}

func (r *mutationResolver) UpdateStory(ctx context.Context, storyID, title, jsonData, status string) (bool, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return false, fmt.Errorf("unauthorized")
	}

	story := models.Story{
		ID:       storyID,
		AuthorID: loginUser.ID,
		Title:    title,
		Content:  jsonData,
		Status:   status,
	}

	if err := repo.UpdateStory(r.DB, &story); err != nil {
		return false, err
	}

	return true, nil
}
