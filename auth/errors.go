package auth

import (
	"errors"

	"github.com/hhyouke/server/models"
)

var (
	// ErrInvalidPassword invalid passwoird
	ErrInvalidPassword    = errors.New(models.ErrInvalidPassword)
	ErrInvalidAccount     = errors.New(models.ErrInvalidAccount)
	ErrAccountOccupied    = errors.New(models.ErrOccupiedAccount)
	ErrNilAccount         = errors.New(models.ErrNilAccount)
	ErrInvalidOrgJoinCode = errors.New(models.ErrInvalidOrgJoinCode)
	ErrOrgUnactivated     = errors.New(models.ErrOrgUnActivated)
)
