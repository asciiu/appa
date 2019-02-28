package auth

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	repoUser "github.com/asciiu/appa/apiql/db/sql"
	"github.com/asciiu/appa/apiql/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// refresh is set to 15 days
const RefreshDuration = 360 * time.Hour

//const jwtDuration = 20 * time.Minute
const JwtDuration = 1 * time.Hour

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

func validateAndGetUserID(c *http.Cookie) (string, error) {
	return "", nil
}

func getUserByID(db *sql.DB, userId string) User {
	return User{}
}

func CreateJwtToken(userID string, duration time.Duration) (string, error) {
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
// func renewTokens(c echo.Context, refreshToken *apiModels.RefreshToken) {
// 	// renew access
// 	accessToken, err := createJwtToken(refreshToken.UserID, jwtDuration)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// renew the refresh token
// 	expiresOn := time.Now().Add(refreshDuration)
// 	selectAuth := refreshToken.Renew(expiresOn)

// 	c.Response().Header().Set("set-authorization", accessToken)
// 	c.Response().Header().Set("set-refresh", selectAuth)
// }

type Req struct {
	OperationName string `json:"operationName"`
}

// Middleware decodes the share session cookie and packs the session into context
func Secure(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), userCtxKey, nil)
			auth := r.Header.Get("Authorization")
			str := strings.Replace(auth, "Bearer ", "", 1)

			if str != "" {
				tok, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(os.Getenv("appa_JWT")), nil
				})

				if tok != nil && err == nil {
					userID := tok.Claims.(jwt.MapClaims)["jti"].(string)
					user, _ := repoUser.FindUserByID(db, userID)
					if user != nil {
						ctx = context.WithValue(r.Context(), userCtxKey, user)
					}
				}
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}
