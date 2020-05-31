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

type TokenRepo struct {
	db *pg.DB
}

func NewTokenRepo(db *pg.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (repo *TokenRepo) FindRefreshToken(selector string) (*refresh.RefreshToken, error) {
	t := new(refresh.RefreshToken)
	err := repo.db.Model(t).Where("selector = ?", selector).Select()
	return t, err
}

func (repo *TokenRepo) InsertRefreshToken(token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	_, err := repo.db.Model(token).Returning("*").Insert()
	return token, err
}

func (repo *TokenRepo) DeleteRefreshToken(token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	err := repo.db.Delete(token)
	return token, err
}

func (repo *TokenRepo) DeleteRefreshTokenBySelector(selector string) error {
	garbage := new(refresh.RefreshToken)
	_, err := repo.db.Model(garbage).Where("selector = ?", selector).Delete()
	return err
}

func (repo *TokenRepo) DeleteStaleTokens(expiresOn time.Time) error {
	garbage := new(refresh.RefreshToken)
	_, err := repo.db.Model(garbage).Where("expires_on < ?", expiresOn).Delete()
	return err
}

func (repo *TokenRepo) UpdateRefreshToken(token *refresh.RefreshToken) (*refresh.RefreshToken, error) {
	sqlStatement := `update refresh_tokens 
	set selector = ?, 
	token_hash = ?, 
	expires_on = ? 
	where user_id = ? and id = ?`
	_, err := repo.db.Exec(sqlStatement,
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
