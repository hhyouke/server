package user

import (
	"time"

	"github.com/hhyouke/server/models"
	"github.com/hhyouke/server/utils"
	"github.com/jinzhu/gorm"
)

const (
	defaultApp = "50tin"
	package0   = "p0"
	package1   = "p1"
)

// CreateWithName create user by username
func CreateWithName(tx *gorm.DB, idInst *utils.IDInstance, username string) (*models.User, error) {
	id := idInst.NextID()
	u := &models.User{
		ID:         id,
		Nickname:   username,
		CreateUser: id,
		UpdateUser: id,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
