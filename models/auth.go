package models

import (
	"time"

	"github.com/hhyouke/server/auth/claims"
	"github.com/jmoiron/sqlx"
)

const (
	accessTokenExpDuration  = time.Minute * 15
	refreshTokenExpDuration = time.Hour * 24 * 7
)

// Auth is the persistent model of auth
type Auth struct {
	UserID     int64
	SignOrg    string `gorm:"-"`
	Provider   string
	UID        string
	Password   string
	ID         int64
	CreateUser int64
	UpdateUser int64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
	SignLogs
}

// ToClaims convert to access-token Claims
func (a Auth) ToClaims() *claims.Claims {
	claims := claims.Claims{}
	claims.Provider = a.Provider
	claims.Id = a.UID
	claims.SignOrg = a.SignOrg
	// claims.UserID = a.UserID
	claims.ExpiresAt = time.Now().Local().Add(accessTokenExpDuration).Unix()
	claims.IssuedAt = time.Now().Local().Unix()
	return &claims
}

// ToRefreshClaims convert to refresh-token Claims
func (a Auth) ToRefreshClaims() *claims.Claims {
	claims := claims.Claims{}
	claims.Provider = a.Provider
	claims.Id = a.UID
	claims.SignOrg = a.SignOrg
	// claims.UserID = a.UserID
	claims.ExpiresAt = time.Now().Local().Add(refreshTokenExpDuration).Unix()
	claims.IssuedAt = time.Now().Local().Unix()
	return &claims
}

// Create create new auth
func (a *Auth) Create(tx *sqlx.Tx) error {
	iq := "INSERT into t_auth(id, user_id, provider, uid, password, create_user, update_user, create_time, update_time) VALUES (:id, :user_id, :provider, :uid, :password, :create_user, :update_user, :create_time, :update_time)"
	if _, err := tx.NamedExec(iq, a); err != nil {
		return err
	}
	return nil
}

// TableName db table name
func (a *Auth) TableName() string {
	return "t_auth"
}
