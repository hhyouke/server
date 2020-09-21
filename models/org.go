package models

import "time"

// Org org represents tenant
type Org struct {
	ID         int64
	UID        string
	JoinCode   string
	Nickname   string
	Verified   bool
	Activated  bool
	RealName   string
	CreateUser int64
	UpdateUser int64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
}

// TableName db table name
func (o *Org) TableName() string {
	return "t_org"
}

// OrgUser org-user relations
type OrgUser struct {
	ID         int64
	OrgID      int64
	UserID     int64
	RoleCode   string
	CreateUser int64
	UpdateUser int64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
}

// TableName db table name
func (ou *OrgUser) TableName() string {
	return "t_org_user"
}

// OrgApp org-app relations
type OrgApp struct {
	ID         int64
	AppUID     string
	PackageUID string `gorm:"column:package_uid"`
	UserID     int64
	OrgID      int64 `gorm:"index:ORG_ID"`
	CreateUser int64
	UpdateUser int64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
}

// TableName db table name
func (oa *OrgApp) TableName() string {
	return "t_org_app"
}
