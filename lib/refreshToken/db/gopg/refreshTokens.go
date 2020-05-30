package gopg

import (
	"time"

	refresh "github.com/asciiu/appa/lib/refreshToken/models"
	"github.com/go-pg/pg/v10"
)

// functions
//   DeleteStaleTokens
//   DeleteRefreshToken
//   DeleteRefreshTokenBySelector
//   FindRefreshToken
//   InsertRefreshToken
//   UpdateRefreshToken

func FindRefreshToken(db *pg.DB, selector string) (*refresh.RefreshToken, error) {
	t := new(refresh.RefreshToken)
	err := db.Model(t).Where("selector = ?", selector).Select()
	return t, err
}

func InsertRefreshToken(db *pg.DB, token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	_, err := db.Model(token).Returning("*").Insert()
	return token, err
}

func DeleteRefreshToken(db *pg.DB, token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	err := db.Delete(token)
	return token, err
}

func DeleteRefreshTokenBySelector(db *pg.DB, selector string) error {

	garbage := new(refresh.RefreshToken)
	_, err := db.Model(garbage).Where("selector = ?", selector).Delete()
	return err
}

func DeleteStaleTokens(db *pg.DB, expiresOn time.Time) error {
	garbage := new(refresh.RefreshToken)
	_, err := db.Model(garbage).Where("expires_on < ?", expiresOn).Delete()
	return err
}

func UpdateRefreshToken(db *pg.DB, token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	sqlStatement := `update refresh_tokens 
	set selector = ?, 
	token_hash = ?, 
	expires_on = ? 
	where user_id = ? and id = ?`
	_, err := db.Exec(sqlStatement,
		token.Selector,
		token.TokenHash,
		token.ExpiresOn,
		token.UserID,
		token.ID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
