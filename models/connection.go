package models

import (
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/hhyouke/server/conf"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	defaultMaxOpenConn = 50
	defaultMaxIdleConn = 10
	defaultAutoMigrate = false
)

// Connect connect to db
func Connect(conf *conf.DBConfiguration, logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Dialect, conf.URL)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	db.SetLogger(NewDBLogger(logger))

	// switch log mode on
	if conf.LogMode {
		db.LogMode(true)
	}
	// pool settings
	if conf.MaxOpenConn != 0 {
		db.DB().SetMaxOpenConns(conf.MaxOpenConn)
	} else {
		db.DB().SetMaxOpenConns(defaultMaxOpenConn)
	}

	if conf.MaxIdelConn != 0 {
		db.DB().SetMaxIdleConns(conf.MaxIdelConn)
	} else {
		db.DB().SetMaxIdleConns(defaultMaxIdleConn)
	}

	// fmt.Println(db.DB().Stats())

	err = db.DB().Ping()
	if err != nil {
		return nil, errors.Wrap(err, "checking database connection")
	}

	if conf.AutoMigrate {
		migDB := db.New()
		logger, _ := zap.NewDevelopment()
		logger = logger.With(zap.String("task", "migration"))
		migDB.SetLogger(NewDBLogger(logger))
		if err := AutoMigrate(migDB); err != nil {
			return nil, errors.Wrap(err, "migrating tables")
		}
	}

	return db, nil
}

// AutoMigrate automigrate db, only in dev mode
func AutoMigrate(db *gorm.DB) error {
	db = db.AutoMigrate(Auth{},
		User{},
		Org{},
		OrgUser{},
		OrgApp{},
	)
	return db.Error
}
