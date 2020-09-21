package org

import (
	"errors"
	"time"

	"github.com/hhyouke/server/models"
	"github.com/hhyouke/server/utils"
	"github.com/jinzhu/gorm"
)

const (
	baseUID = "F6BGQ2DZX9C7P5IK3MJUAR4WYLTNVE8S"
	padUID  = "H"

	defaultApp = "50tin"
	package0   = "p0"
	// package1   = "p1"
	superAdmin = "sa"
	staff      = "staff"
)

// CreateP0 create p0-subscription backstage org info
func CreateP0(tx *gorm.DB, idInst *utils.IDInstance, userInfo *models.User) (*models.Org, error) {
	orgID := idInst.NextID()
	uid := utils.NewCustomInvCode(baseUID, padUID).IDToCode(orgID)
	joinCode := utils.NewInvCode().IDToCode(orgID)
	orgInfo := &models.Org{
		ID:         orgID,
		UID:        uid,
		JoinCode:   joinCode,
		Verified:   false,
		Activated:  false,
		CreateUser: userInfo.ID,
		UpdateUser: userInfo.ID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(orgInfo).Error; err != nil {
		return nil, err
	}
	// create org-user relational info
	ouid := idInst.NextID()
	ou := &models.OrgUser{
		ID:         ouid,
		OrgID:      orgID,
		UserID:     userInfo.ID,
		RoleCode:   superAdmin,
		CreateUser: userInfo.ID,
		UpdateUser: userInfo.ID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(ou).Error; err != nil {
		return nil, err
	}
	return orgInfo, nil
}

// AddStaffWithJoinCode add staff to org with the join code
// 1. validate the org(with the given join code) exists and activated
// 2. create org-user relations
func AddStaffWithJoinCode(tx *gorm.DB, joinCode string, idInst *utils.IDInstance, userInfo *models.User) (*models.Org, error) {
	var curOrg models.Org
	if tx.Model(&curOrg).Select([]string{"id", "activated"}).Where("join_code = ?", joinCode).Scan(&curOrg).RecordNotFound() {
		return nil, errors.New(models.ErrInvalidOrgJoinCode)
	}
	if !curOrg.Activated {
		return nil, errors.New(models.ErrOrgUnActivated)
	}
	// only when the org is activated, then it can add memebers
	ou := &models.OrgUser{
		ID:         idInst.NextID(),
		OrgID:      curOrg.ID,
		UserID:     userInfo.ID,
		RoleCode:   staff,
		CreateUser: userInfo.ID,
		UpdateUser: userInfo.ID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(ou).Error; err != nil {
		return nil, err
	}
	return &curOrg, nil
}

// CreateP0Subscription create a [p0] subscription
func CreateP0Subscription(tx *gorm.DB, idInst *utils.IDInstance, orgInfo *models.Org, userInfo *models.User) error {
	id := idInst.NextID()
	ua := &models.OrgApp{
		ID:         id,
		AppUID:     defaultApp,
		PackageUID: package0,
		UserID:     userInfo.ID,
		OrgID:      orgInfo.ID,
		CreateUser: userInfo.ID,
		UpdateUser: userInfo.ID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.Create(ua).Error; err != nil {
		return err
	}
	return nil
}
