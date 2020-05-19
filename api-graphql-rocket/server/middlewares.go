package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// refresh window
const refreshDuration = 9 * time.Hour

// token valid window
const jwtDuration = 2 * time.Hour

func createJwtToken(userID string, duration time.Duration) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	token, err := rawToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// Middleware decodes the share session cookie and packs the session into context
func secure(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), userContextKey, nil)
			auth := r.Header.Get("Authorization")
			str := strings.Replace(auth, "Bearer ", "", 1)

			if str != "" {
				tok, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(os.Getenv("JWT_SECRET")), nil
				})

				// if valid token and no error
				if tok != nil && err == nil {
					userID := tok.Claims.(jwt.MapClaims)["jti"].(string)
					loginUser, err := userRepo.FindUserByID(db, userID)
					if loginUser != nil {
						ctx = context.WithValue(r.Context(), userContextKey, loginUser)
					}
					if err != nil {
						log.Error(err.Error())
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
							refreshToken, err := tokenRepo.FindRefreshToken(db, selector)
							if err == nil {
								if refreshToken.Valid(authenticator) {
									accessToken, err := createJwtToken(refreshToken.UserID, jwtDuration)
									if err == nil {
										expiresOn := time.Now().Add(refreshDuration)
										selectAuth := refreshToken.Renew(expiresOn)

										w.Header().Set("set-authorization", accessToken)
										w.Header().Set("set-refresh", selectAuth)
										if _, err := tokenRepo.UpdateRefreshToken(db, refreshToken); err != nil {
											log.Println(err)
										}
									}
								}
								if refreshToken.ExpiresOn.Before(time.Now()) {
									tokenRepo.DeleteRefreshToken(db, refreshToken)
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
