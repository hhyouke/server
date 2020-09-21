package models

// common used error code
const (
	ErrNil                = "0"
	ErrInternal           = "1000"
	ErrInvalidAccount     = "2000"
	ErrInvalidPassword    = "2001"
	ErrInvalidToken       = "2002"
	ErrOccupiedAccount    = "2003"
	ErrNilAccount         = "2004"
	ErrInvalidOrgJoinCode = "3001"
	ErrOrgUnActivated     = "3002"
	ErrInvalidOrg         = "3003"
)

// ErrText error text
var ErrText = map[string]string{
	ErrNil:                "api call succeeded",
	ErrInvalidAccount:     "invalid account",
	ErrInvalidPassword:    "invalid password",
	ErrInvalidToken:       "invalid token",
	ErrOccupiedAccount:    "account already occupied",
	ErrNilAccount:         "account doesn't exist",
	ErrInvalidOrgJoinCode: "invalid org join code",
	ErrOrgUnActivated:     "org has not been activated yet",
	ErrInvalidOrg:         "org invalid",
}
