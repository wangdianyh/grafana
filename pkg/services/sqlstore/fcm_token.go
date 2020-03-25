package sqlstore

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"log"
	"time"
)

func init() {
	bus.AddHandler("sql", SaveToken)
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
