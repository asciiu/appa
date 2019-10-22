package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	asql "github.com/asciiu/appa/api/db/sql"
	constRes "github.com/asciiu/appa/common/constants/response"
	repoUser "github.com/asciiu/appa/user-service/db/sql"
	protoUser "github.com/asciiu/appa/user-service/proto/user"
	micro "github.com/micro/go-micro"

	apiModels "github.com/asciiu/appa/api/models"
	models "github.com/asciiu/appa/user-service/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// refresh is set to 15 days
const refreshDuration = 360 * time.Hour

//const jwtDuration = 20 * time.Minute
const jwtDuration = 12 * time.Hour

type AuthController struct {
	DB         *sql.DB
	UserClient protoUser.UserService
}

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// swagger:parameters signup
type SignupRequest struct {
	// Optional.
	// in: body
	First string `json:"first"`
	// Optional.
	// in: body
	Last string `json:"last"`
	// Required. Must be unique! We need to validate these coming in.
	// in: body
	Email string `json:"email"`
	// Required. We need password requirements.
	// in: body
	Password string `json:"password"`
}

// A ResponseSuccess will always contain a status of "successful".
// swagger:model responseSuccess
type ResponseSuccess struct {
	Status string    `json:"status"`
	Data   *UserData `json:"data"`
}

type Device struct {
	DeviceID         string `json:"deviceID"`
	DeviceType       string `json:"deviceType"`
	ExternalDeviceID string `json:"externalDeviceID"`
	DeviceToken      string `json:"deviceToken"`
}

type UserData struct {
	User    *models.UserInfo `json:"user"`
	Devices []*Device        `json:"devices"`
}

func NewAuthController(db *sql.DB, service micro.Service) *AuthController {
	controller := AuthController{
		DB:         db,
		UserClient: protoUser.NewUserService("users", service.Client()),
	}

	return &controller
}

func createJwtToken(userID string, duration time.Duration) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	token, err := rawToken.SignedString([]byte(os.Getenv("appa_JWT")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// Renews the refresh token and the access token in the reponse headers.
func renewTokens(c echo.Context, refreshToken *apiModels.RefreshToken) {
	// renew access
	accessToken, err := createJwtToken(refreshToken.UserID, jwtDuration)
	if err != nil {
		log.Fatal(err)
	}

	// renew the refresh token
	expiresOn := time.Now().Add(refreshDuration)
	selectAuth := refreshToken.Renew(expiresOn)

	c.Response().Header().Set("set-authorization", accessToken)
	c.Response().Header().Set("set-refresh", selectAuth)
}

// A custom middleware function to check the refresh token on each request.
func (controller *AuthController) RefreshAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return next(c)
		}

		tokenString := strings.Split(auth, " ")[1]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("appa_JWT")), nil
		})

		if err != nil && c.Request().Method != http.MethodOptions {

			selectAuth := c.Request().Header.Get("Refresh")
			if selectAuth != "" {
				sa := strings.Split(selectAuth, ":")

				if len(sa) != 2 {
					return next(c)
				}

				selector := sa[0]
				authenticator := sa[1]

				refreshToken, err := asql.FindRefreshToken(controller.DB, selector)
				if err != nil {
					return next(c)
				}

				if refreshToken.Valid(authenticator) {
					// renew access
					renewTokens(c, refreshToken)
					_, err3 := asql.UpdateRefreshToken(controller.DB, refreshToken)

					if err3 != nil {
						log.Fatal(err3)
					}
				}

				if refreshToken.ExpiresOn.Before(time.Now()) {
					asql.DeleteRefreshToken(controller.DB, refreshToken)
				}
			}
		}

		return next(c)
	}
}

// swagger:parameters login
type LoginRequest struct {
	// Required. Backend code does not check email atm.
	// in: body
	Email string `json:"email" validate:"required,email"`
	// Required. Backend code does not have any password requirements atm.
	// in: body
	Password string `json:"password" validate:"required"`
	// Optional. Return refresh token in response
	// in: body
	Remember bool `json:"remember"`
}

// swagger:route POST /login authentication login
//
// user login (open)
//
// The login endpoint returns an authorization token in the "set-authorization" response header.
// You may also receive an optional refresh token if you include 'remember': true in your post request.
// Both tokens will be returned in the reponse headers as "set-refresh" and "set-authorization".
//
// responses:
//  200: responseSuccess "data" will be non null with "status": "success"
//  400: responseError email and password are not found in request with "status": "fail"
//  401: responseError unauthorized user because of incorrect password with "status": "fail"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *AuthController) HandleLogin(c echo.Context) error {
	defer c.Request().Body.Close()

	loginRequest := new(LoginRequest)
	err := c.Bind(loginRequest)
	if err != nil {
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{"malformed json request for 'email' and 'password'"},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err = c.Validate(loginRequest); err != nil {
		msgs := strings.Split(err.Error(), "\n")
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{msgs},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	// lookup user by email
	user, err := repoUser.FindUser(controller.DB, loginRequest.Email)
	switch {
	case err == sql.ErrNoRows:
		response := &ResponseError{
			Status:   constRes.Fail,
			Messages: []string{"password/login incorrect"},
		}
		// no user by this email send unauthorized response
		return c.JSON(http.StatusUnauthorized, response)

	case err != nil:
		response := &ResponseError{
			Status:   constRes.Error,
			Messages: []string{err.Error()},
		}
		return c.JSON(http.StatusInternalServerError, response)

	default:
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password)) == nil {

			accessToken, err := createJwtToken(user.ID, jwtDuration)
			if err != nil {
				response := &ResponseError{
					Status:   constRes.Error,
					Messages: []string{err.Error()},
				}
				return c.JSON(http.StatusInternalServerError, response)
			}

			// issue a refresh token if remember is true
			if loginRequest.Remember {
				refreshToken := apiModels.NewRefreshToken(user.ID)
				renewTokens(c, refreshToken)

				_, err3 := asql.InsertRefreshToken(controller.DB, refreshToken)

				if err3 != nil {
					response := &ResponseError{
						Status:   constRes.Error,
						Messages: []string{err.Error()},
					}
					return c.JSON(http.StatusInternalServerError, response)
				}
			} else {
				c.Response().Header().Set("set-authorization", accessToken)
			}

			response := &ResponseSuccess{
				Status: constRes.Success,
				Data: &UserData{
					User: user.Info(),
				},
			}

			return c.JSON(http.StatusOK, response)
		}
	}

	response := &ResponseError{
		Status:  constRes.Fail,
		Messages: []string{"password/login incorrect"},
	}
	return c.JSON(http.StatusUnauthorized, response)
}

// swagger:route GET /logout authentication logout
//
// logout user (protected)
//
// If a valid refresh token is found in the request headers, the server
// will attempt to remove the refresh token from the database.
//
//	Responses:
//	  200: responseSuccess data will be null with status "success"
//	  400: responseError you either sent in no refresh token or the refresh token in the header is not valid with status: "fail"
func (controller *AuthController) HandleLogout(c echo.Context) error {
	selectAuth := c.Request().Header.Get("Refresh")
	if selectAuth != "" {
		sa := strings.Split(selectAuth, ":")

		if len(sa) != 2 {
			response := &ResponseError{
				Status:  constRes.Fail,
				Messages: []string{"refresh token invalid"},
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		asql.DeleteRefreshTokenBySelector(controller.DB, sa[0])
	}

	response := &ResponseSuccess{
		Status: constRes.Success,
	}
	return c.JSON(http.StatusOK, response)
}

// swagger:route POST /signup authentication signup
//
// user registration (open)
//
// Registers a new user. Expects email to be unique. Duplicate email will result
// in a bad request.
//
// responses:
//  200: responseSuccess "data" will be non null with "status": "success"
//  400: responseError message should relay information with regard to bad request with "status": "fail"
//  410: responseError the user-service is not reachable. The user-service is a microservice that runs independantly from the api. When we take it offline you will receive this error.
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *AuthController) HandleSignup(c echo.Context) error {
	signupRequest := SignupRequest{}

	err := json.NewDecoder(c.Request().Body).Decode(&signupRequest)
	if err != nil {
		response := &ResponseError{
			Status:  constRes.Fail,
			Messages: []string{err.Error()},
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	if signupRequest.Email == "" || signupRequest.Password == "" {
		response := &ResponseError{
			Status:  constRes.Fail,
			Message: []string{"email and password are required"},
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	createRequest := protoUser.CreateUserRequest{
		First:    signupRequest.First,
		Last:     signupRequest.Last,
		Email:    signupRequest.Email,
		Password: signupRequest.Password,
	}

	r, _ := controller.UserClient.CreateUser(context.Background(), &createRequest)
	if r.Status != constRes.Success {
		response := &ResponseError{
			Status:  r.Status,
			Messages: []string{r.Message},
		}

		if r.Status == constRes.Fail {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == constRes.Error {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseSuccess{
		Status: constRes.Success,
		Data: &UserData{
			User: &models.UserInfo{
				UserID: r.Data.User.UserID,
				First:  r.Data.User.First,
				Last:   r.Data.User.Last,
				Email:  r.Data.User.Email,
			},
		},
	}

	return c.JSON(http.StatusOK, response)
}
