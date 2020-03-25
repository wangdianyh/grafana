package api

import (
	//"fmt"
	//"log"
	//"encoding/json"
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
