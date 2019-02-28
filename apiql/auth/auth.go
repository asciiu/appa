package auth

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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
const JwtDuration = 12 * time.Hour

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
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error reading body: %v", err)
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}

			// forward orginal body to next chain
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			var req Req
			json.Unmarshal(body, &req)
			switch {
			case req.OperationName == "" || req.OperationName == "IntrospectionQuery":
				next.ServeHTTP(w, r)
				return
			case req.OperationName == "Login" || req.OperationName == "SignUp":
				next.ServeHTTP(w, r)
				return
			default:
				auth := r.Header.Get("Authorization")
				refresh := r.Header.Get("Refresh")
				fmt.Println(auth)
				fmt.Println(refresh)

				response := `{"message": "auth required"}`

				http.Error(w, response, http.StatusForbidden)
				return
			}

			//jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//	// Don't forget to validate the alg is what you expect:
			//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			//	}

			//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			//	return []byte(os.Getenv("appa_JWT")), nil
			//})

			// get the user from the database
			//user := getUserByID(db, userId)

			//// put it in context
			//ctx := context.WithValue(r.Context(), userCtxKey, user)

			//// and call the next with our new context
			//r = r.WithContext(ctx)
			//next.ServeHTTP(w, r)
			//http.Error(w, "niope", http.StatusForbidden)
			//next.ServeHTTP(w, r)
			//return
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
