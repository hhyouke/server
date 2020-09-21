package password

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/hhyouke/server/app/account"
	"github.com/hhyouke/server/app/org"
	"github.com/hhyouke/server/app/user"
	"github.com/hhyouke/server/auth"
	"github.com/hhyouke/server/auth/tokens"
	"github.com/hhyouke/server/models"
	"github.com/hhyouke/server/utils"
	"go.uber.org/zap"
)

var (
	superAdmin = "sa"
	staff      = "staff"
)

// PasswordAuthenticateHandler default sign-in handler of password pattern
var PasswordAuthenticateHandler = func(context *auth.Context) (*tokens.Token, error) {
	var (
		req         = context.Request
		db          = context.Auth.Config.DB
		provider, _ = context.Provider.(*Provider)
		authInfo    models.Auth
	)

	// id generator
	machineIDStr := context.Auth.MachineID
	machineID, err := strconv.ParseInt(machineIDStr, 10, 64)
	idGenerator := utils.NewCommonIDInstance(machineID)

	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	uid := strings.TrimSpace(req.Form.Get("username"))
	rawPassword := strings.TrimSpace(req.Form.Get("password"))
	signOrg := strings.TrimSpace(req.Form.Get("org"))
	agent := req.Header.Get("User-Agent")

	if err := account.CheckExistence(db, context.Auth.AuthModel, provider.GetName(), uid, &authInfo); err == nil {
		return nil, auth.ErrNilAccount
	}

	if err := provider.Encryptor.Compare(rawPassword, authInfo.Password); err != nil {
		context.Auth.Logger.Error("invalid password", zap.String("uid", uid), zap.String("password", rawPassword))
		return nil, auth.ErrInvalidPassword
	}

	// if doesn't specify the sign-org
	// check sign-log, get the latest entry, take that org as the sign-org and
	// create the new sign-log
	ch := make(chan string)
	go func() {
		tx := db.Begin()
		if signOrg == "" {
			if signOrg, err = account.UpdateSignLogs(db, uid, agent, provider.GetName(), idGenerator, req); err != nil {
				context.Auth.Logger.Error("update sign-logs error", zap.Any("errors", err)) // log errors
				tx.Rollback()
			}
		} else {
			if err := account.UpdateSignLogsWithSignOrg(db, signOrg, uid, agent, provider.GetName(), idGenerator, req); err != nil {
				context.Auth.Logger.Error("update sign-logs error", zap.Any("errors", err)) // log errors
				tx.Rollback()
			}
		}
		tx.Commit()
		ch <- signOrg
	}()

	// block until update sign-logs finished
	<-ch
	// assign the sign-org info in order to put in token
	authInfo.SignOrg = signOrg

	token, err := tokens.GenerateTokenPair(&authInfo)
	if err != nil {
		return nil, errors.New(models.ErrInternal)
	}

	return token, nil
}

// PasswordRefreshTokenHandler is the refresh-token behavior implementation
// of the password pattern
var PasswordRefreshTokenHandler = func(context *auth.Context) (*tokens.Token, error) {
	type reqBody struct {
		RefreshToken string
	}

	var (
		r           = context.Request
		db          = context.Auth.Config.DB
		provider, _ = context.Provider.(*Provider)
		rBody       reqBody
		authInfo    models.Auth
	)

	// id generator
	machineIDStr := context.Auth.MachineID
	machineID, err := strconv.ParseInt(machineIDStr, 10, 64)
	idGenerator := utils.NewCommonIDInstance(machineID)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return nil, err
	}

	if err = json.Unmarshal(body, &rBody); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	agent := r.Header.Get("User-Agent")

	rt, err := jwt.Parse(rBody.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// validate the sign algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// using HMAC for now
		return []byte(tokens.TokenSignSecret), nil
	})
	// verify the token
	if claims, ok := rt.Claims.(jwt.MapClaims); ok && rt.Valid {
		uid := claims["jti"].(string)
		if err := account.CheckExistence(db, context.Auth.AuthModel, provider.GetName(), uid, &authInfo); err == nil {
			return nil, auth.ErrNilAccount
		}

		// update sign-log
		go func() {
			tx := db.Begin()
			if err := account.UpdateSignLogsWithSignOrg(db, claims["org"].(string), uid, agent, claims["provider"].(string), idGenerator, r); err != nil {
				context.Auth.Logger.Error("update sign-logs error", zap.Any("errors", err))
				tx.Rollback()
			}
			tx.Commit()
		}()

		// assign sign-org for token
		authInfo.SignOrg = claims["org"].(string)
		token, err := tokens.GenerateTokenPair(&authInfo)
		if err != nil {
			return nil, errors.New(models.ErrInternal)
		}
		return token, nil
	}
	// anything wrong with token, whether expire or malformat got the unified error strings
	return nil, errors.New(models.ErrInvalidToken)
}

// PasswordSignUpHandler the sign-up handler of password pattern
var PasswordSignUpHandler = func(context *auth.Context) (token *tokens.Token, err error) {
	type reqBody struct {
		Username string
		Password string
		JoinCode string `json:"join_code,omitempty"`
	}
	var (
		r           = context.Request
		db          = context.Auth.Config.DB
		logger      = context.Auth.Config.Logger
		provider, _ = context.Provider.(*Provider)
		rBody       reqBody
		authInfo    models.Auth
		orgInfo     *models.Org
	)

	// id generator
	machineIDStr := context.Auth.MachineID
	machineID, err := strconv.ParseInt(machineIDStr, 10, 64)
	idGenerator := utils.NewCommonIDInstance(machineID)
	agent := r.Header.Get("User-Agent")

	tx := db.Begin() // new transaction
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logger.Panic("internal errors", zap.Any("err", p))
		} else if err != nil {
			logger.Error("sign-up errors, rollback tx", zap.Any("errors", err))
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &rBody); err != nil {
		return nil, err
	}

	// check if the account has already been occupied
	if err := account.CheckExistence(db, context.Auth.AuthModel, provider.GetName(), rBody.Username, &authInfo); err != nil {
		return nil, err
	}

	// create user
	u, err := user.CreateWithName(tx, idGenerator, rBody.Username)
	if err != nil {
		return nil, err
	}

	// create auth
	cryptedPassword, err := provider.Encryptor.Digest(rBody.Password)
	if err != nil {
		return nil, err
	}
	if _, err = account.CreateWithUsernamePassword(tx, idGenerator, provider.GetName(), cryptedPassword, u); err != nil {
		return nil, err
	}

	// this will create the account and create an default org at the same time
	// also will make a subscription(package p0) linked to the org
	if rBody.JoinCode == "" {
		orgInfo, err = org.CreateP0(tx, idGenerator, u)
		if err != nil {
			return nil, err
		}

		// create a [p0] subscription for the SA
		if err = org.CreateP0Subscription(tx, idGenerator, orgInfo, u); err != nil {
			return nil, err
		}
	} else { // join an existing activated org and makes no subscription
		if orgInfo, err = org.AddStaffWithJoinCode(tx, rBody.JoinCode, idGenerator, u); err != nil {
			return nil, err
		}
	}

	// update sign-logs
	go func() {
		tx := db.Begin()
		if err := account.UpdateSignLogsWithSignOrg(tx, orgInfo.UID, u.Nickname, agent, provider.GetName(), idGenerator, r); err != nil {
			logger.Error("update sign-logs error", zap.Any("errors", err))
			tx.Rollback()
		}
		tx.Commit()
	}()

	// assign authInfo to generate token
	authInfo.SignOrg = orgInfo.UID
	authInfo.UID = u.Nickname
	authInfo.Provider = provider.GetName()
	// gen token if everything went well
	token, err = tokens.GenerateTokenPair(&authInfo)
	if err != nil {
		return nil, errors.New(models.ErrInternal)
	}

	return token, nil
}
