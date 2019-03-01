package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	repoUser "github.com/asciiu/appa/apiql/db/sql"
	"github.com/asciiu/appa/apiql/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// refresh window
const RefreshDuration = 9 * time.Hour

// token valid window
const JwtDuration = 2 * time.Hour

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
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

				// if valid token and no error
				if tok != nil && err == nil {
					userID := tok.Claims.(jwt.MapClaims)["jti"].(string)
					user, _ := repoUser.FindUserByID(db, userID)
					if user != nil {
						ctx = context.WithValue(r.Context(), userCtxKey, user)
					}
				}

				// if expired token examine refresh
				if err != nil && r.Method != http.MethodOptions {
					selectAuth := r.Header.Get("Refresh")
					if selectAuth != "" {
						sa := strings.Split(selectAuth, ":")
						if len(sa) == 2 {
							selector := sa[0]
							authenticator := sa[1]
							refreshToken, err := repoUser.FindRefreshToken(db, selector)
							if err == nil {
								if refreshToken.Valid(authenticator) {
									accessToken, err := CreateJwtToken(refreshToken.UserID, JwtDuration)
									if err == nil {
										expiresOn := time.Now().Add(RefreshDuration)
										selectAuth := refreshToken.Renew(expiresOn)

										w.Header().Set("set-authorization", accessToken)
										w.Header().Set("set-refresh", selectAuth)
										if _, err := repoUser.UpdateRefreshToken(db, refreshToken); err != nil {
											log.Println(err)
										}
									}
								}
								if refreshToken.ExpiresOn.Before(time.Now()) {
									repoUser.DeleteRefreshToken(db, refreshToken)
								}
							}
						}
					}
				}
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
