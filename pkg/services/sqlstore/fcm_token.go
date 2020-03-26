package sqlstore

import (
	//"fmt"
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"log"
	//"strconv"
	"strings"
	"time"
)

func init() {
	bus.AddHandler("sql", SaveToken)
	bus.AddHandler("sql", LoadTokenByChannel)
	bus.AddHandler("sql", LoadTokenByUser)
}
func SaveToken(cmd *models.AddTokenCommand) error {
	if cmd.Token == "" || cmd.ChannelId == "" {
		return models.FCMFieldMissing
	}

	return inTransaction(func(sess *DBSession) error {
		fcmToken := models.FcmToken{
			Token:     cmd.Token,
			ChannelId: cmd.ChannelId,
			UserId:    cmd.UserId,
			Created:   time.Now(),
			Updated:   time.Now(),
		}

		_, err := sess.Insert(&fcmToken)
		if err != nil {
			log.Fatal(err)
		}

		cmd.Result = fcmToken

		return err
	})
}

func LoadTokenByChannel(query *models.GetTokenByChannelQuery) error {
	var tokens []*models.FcmToken

	err := x.SQL("select * from fcm_token where channel_id=" + query.ChannelId).Find(&tokens)
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
