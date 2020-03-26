package api

import (
	"fmt"
	"log"
	//"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
)

// POST /add-token/:tId/channel/:cId/user/:uId
func addToken(c *models.ReqContext) Response {
	token := c.Params(":tId")
	channel := c.Params(":cId")
	user := c.Params(":uId")

	cmd := models.AddTokenCommand{Token: token, ChannelId: channel, UserId: user}

	if cmd.Token == "" {
		return Error(400, "Missing Token", nil)
	}

	if cmd.ChannelId == "" {
		return Error(400, "Missing Channel ID", nil)
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to add Token to db", err)
	}

	return Success("Token Added!")
}

// GET  /get-token-by-channel/:cId
func getTokenByChannel(c *models.ReqContext) Response {
	channel := c.Params(":cId")
	query := models.GetTokenByChannelQuery{ChannelId: channel}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get Token", err)
	}

	return JSON(200, query.Result)
}

// get tokens by channel in FCM notifier
func GetTokenByChannelForNotification(cId string) ([]*models.FcmToken, error) {
	query := models.GetTokenByChannelQuery{ChannelId: cId}
	if err := bus.Dispatch(&query); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return query.Result, nil
}

// get tokens by user list
func GetTokenByUser(list []string) ([]string, error) {
	query := models.GetTokenByUserQuery{UserId: list}

	if err := bus.Dispatch(&query); err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("tokenlist: %v\n", query.Result)

	return query.Result, nil
}
