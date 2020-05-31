package models

import (
	"log"
	"time"

	balance "github.com/asciiu/appa/lib/balance/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RoleType string
type PermissionType string

const (
	// role types
	AdminRole   RoleType = "admin"
	InvalidRole RoleType = "invalid"

	// permissions
	UnrestrictedPermission PermissionType = "unrestricted"
	UnknownPermission      PermissionType = "unknown"
)

func RoleFromString(role string) RoleType {
	switch role {
	case "admin":
		return AdminRole
	default:
		return InvalidRole
	}
}

func PermissionFromString(permission string) PermissionType {
	switch permission {
	case "unrestricted":
		return UnrestrictedPermission
	default:
		return UnknownPermission
	}
}

func NewUser(username, email, password string) *User {
	// assign new uuid for user
	newID := uuid.New()
	user := User{
		ID:            newID.String(),
		Username:      username,
		Email:         email,
		EmailVerified: false,
		PasswordHash:  GenHash([]byte(password)),
	}
	return &user
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	// hash this using a server secret key
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func GenHash(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	// hash this using a server secret key
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

type Order struct {
	ID  string
	Txt string
}

type User struct {
	ID            string
	Username      string
	AvatarURL     string
	Email         string
	EmailVerified bool
	PasswordHash  string
	CreatedOn     time.Time
	UpdatedOn     time.Time
	DeletedOn     time.Time
}

type UserRepo interface {
	DeleteUserHard(userID string) error
	DeleteUserSoft(userID string) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(userID string) (*User, error)
	InsertUser(user *User) error
	UpdatePassword(userID, hash string) (*User, error)
	UpdateUsername(userID, username string) (*User, error)
	UpdateEmailVerified(userID string, verified bool) (*User, error)
}

type UserSummary struct {
	User    *User            `json:"user"`
	Balance *balance.Balance `json:"balance"`
}

//func (user *User) Summary() *UserSummary {
//	return &UserSummary{
//		User:   user,
//		Balance:
//	}
//}

type UserInfo struct {
	UserID        string `json:"userID"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
}

func (user *User) Info() *UserInfo {
	return &UserInfo{
		UserID:        user.ID,
		Username:      user.Username,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
	}
}

type AdminUser struct {
	ID           uint
	First        string
	Last         string
	Email        string
	PasswordHash string
	Role         RoleType
	Permissions  []PermissionType
	InsertedAt   time.Time
	UpdatedAt    time.Time
}

type AdminInfo struct {
	UserID      uint             `json:"userID"`
	First       string           `json:"first"`
	Last        string           `json:"last"`
	Email       string           `json:"email"`
	Permissions []PermissionType `json:"permissions"`
}

func (admin *AdminUser) Info() *AdminInfo {
	return &AdminInfo{
		UserID:      admin.ID,
		First:       admin.First,
		Last:        admin.Last,
		Email:       admin.Email,
		Permissions: admin.Permissions,
	}
}

func (admin *AdminUser) HasPermission(permission PermissionType) bool {
	for _, p := range admin.Permissions {
		if p == UnrestrictedPermission || p == permission {
			return true
		}
	}
	return false
}
