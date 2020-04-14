package sqlstore

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"log"
	"strings"
	"time"
)

func init() {
	bus.AddHandler("sql", SaveToken)
	bus.AddHandler("sql", LoadTokenByUser)
	bus.AddHandler("sql", LoadToken)
}
func SaveToken(cmd *models.AddTokenCommand) error {
	if cmd.Token == "" || cmd.UserId == "" {
		return models.FCMFieldMissing
	}

	return inTransaction(func(sess *DBSession) error {

		isRegistered, errR := isTokenRegistered(cmd.Token, sess)
		if errR != nil {
			return errR
		} else if isRegistered {
			errU := updateTokenLoginTime(cmd.Token, sess)
			if errU != nil {
				return errU
			}

			return models.ErrTokenRegistered
		}

		fcmToken := models.FcmToken{
			Token:   cmd.Token,
			UserId:  cmd.UserId,
			Created: time.Now(),
			Updated: time.Now(),
		}

		_, err := sess.Insert(&fcmToken)
		if err != nil {
			log.Fatal(err)
			return err
		}

		cmd.Result = fcmToken

		return err
	})
}

func updateTokenLoginTime(token string, sess *DBSession) error {
	update := models.FcmToken{
		Token:   token,
		Updated: time.Now(),
	}

	_, err := sess.Where("token=?", token).Update(&update)
	if err != nil {
		return err
	}

	return nil
}

// check if token is registered already
func isTokenRegistered(token string, sess *DBSession) (bool, error) {
	var fcmToken models.FcmToken
	registered, err := sess.Where("token=?", token).Get(&fcmToken)

	if err != nil {
		log.Fatal(err)
		return false, nil
	}

	if registered {
		return true, nil
	}

	return false, nil
}

// get all tokens registered in db
func LoadToken(query *models.GetTokeQuery) error {
	sqlStr := "select * from fcm_token"
	var tokens []*models.FcmToken

	err := x.SQL(sqlStr).Find(&tokens)
	if err != nil {
		return err
	}
	query.Result = tokens

	return nil
}

func LoadTokenByUser(query *models.GetTokenByUserQuery) error {
	var tokens []*models.FcmToken
	// generate user id list into string format for sql
	uidList := query.UserId
	uidStr := ""
	for _, uid := range uidList {
		uidStr += "'" + uid + "',"
	}
	uidStr = strings.TrimSuffix(uidStr, ",")

	sqlStr := "select token from fcm_token where user_id in (" + uidStr + ")"

	err := x.SQL(sqlStr).Find(&tokens)
	if err != nil {
		return err
	}

	tokenList := make([]string, len(tokens))
	for i, t := range tokens {
		tokenList[i] = t.Token
	}
	query.Result = tokenList

	return nil
}
