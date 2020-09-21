package account

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/hhyouke/server/models"
	"github.com/hhyouke/server/utils"
	"github.com/jinzhu/gorm"
)

// CheckExistence validate if given provider and uid is unique
func CheckExistence(db *gorm.DB, model interface{}, provider, uid string, dataModel *models.Auth) error {
	if db.Model(model).Where("provider = ? AND uid = ?", provider, uid).Scan(dataModel).RecordNotFound() {
		return nil
	}
	return errors.New(models.ErrOccupiedAccount)
}

// CreateWithUsernamePassword create authInfo with username and password
func CreateWithUsernamePassword(tx *gorm.DB, idInst *utils.IDInstance, provider, cryptedPassword string, userInfo *models.User) (*models.Auth, error) {
	// var signLogs []models.SignLog
	authInfo := &models.Auth{
		ID:         idInst.NextID(),
		UID:        userInfo.Nickname,
		UserID:     userInfo.ID,
		Password:   cryptedPassword,
		Provider:   provider,
		CreateUser: userInfo.ID,
		UpdateUser: userInfo.ID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(authInfo).Error; err != nil {
		return nil, err
	}
	return authInfo, nil
}

// UpdateSignLogs update sogn-in logs
func UpdateSignLogs(tx *gorm.DB, signID, agent, provider string, idInst *utils.IDInstance, req *http.Request) (string, error) {
	var (
		a        models.Auth
		orgUID   string
		signLogs []models.SignLog
	)
	if tx.Where("uid=? and provider=?", signID, provider).First(&a).Scan(&a).RecordNotFound() {
		return "", errors.New(models.ErrInvalidAccount)
	}
	signLogs = a.SignLogs.Logs
	if len(signLogs) == 0 {
		return "", errors.New(models.ErrInternal)
	}
	// latest one
	orgUID = signLogs[0].OrgUID
	// ip resolve
	clientIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		clientIP = "unkown"
	}

	if err := tx.Model(&models.Auth{}).Where("uid=? and provider=?", signID, provider).Update("sign_logs", models.SignLogs{
		Count: a.Count + 1,
		Logs: append([]models.SignLog{{
			Agent:  agent,
			ID:     idInst.NextID(),
			UID:    signID,
			OrgUID: orgUID,
			At:     utils.PtrTime(time.Now()),
			IP:     clientIP,
		},
		}, signLogs...),
	}).Error; err != nil {
		return "", err
	}
	return orgUID, nil
}

// UpdateSignLogsWithSignOrg create sign-in logs used when first create account
func UpdateSignLogsWithSignOrg(tx *gorm.DB, signOrg string, signID, agent, provider string, idInst *utils.IDInstance, req *http.Request) error {
	var (
		a        models.Auth
		orgModel models.Org
		signLogs []models.SignLog
	)
	if signOrg == "" { // org can't be nil
		return errors.New(models.ErrInvalidOrg)
	}

	if tx.Where("uid=?", signOrg).Find(&orgModel).RecordNotFound() {
		return errors.New(models.ErrInvalidOrg)
	}

	if tx.Where("uid=? and provider=?", signID, provider).First(&a).Scan(&a).RecordNotFound() {
		return errors.New(models.ErrInvalidAccount)
	}
	signLogs = a.SignLogs.Logs

	// ip resolve
	clientIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		clientIP = "unkown"
	}

	if err := tx.Model(&models.Auth{}).Where("uid=? and provider=?", signID, provider).Update("sign_logs", models.SignLogs{
		Count: a.Count + 1,
		Logs: append([]models.SignLog{{
			Agent:  agent,
			ID:     idInst.NextID(),
			UID:    signID,
			OrgUID: signOrg,
			At:     utils.PtrTime(time.Now()),
			IP:     clientIP,
		},
		}, signLogs...),
	}).Error; err != nil {
		return err
	}
	return nil
}
