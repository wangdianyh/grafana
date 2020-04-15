package api

import (
	"fmt"
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
	"log"
)

// POST /add-token/
func (hs *HTTPServer) AddToken(c *models.ReqContext, cmd models.AddTokenCommand) Response {
	if err := hs.Bus.Dispatch(&cmd); err != nil {

		return Error(500, "Failed to create Token", err)
	}

	if cmd.Result.Id == 0 {
		return Success("Token last seen time updated...")
	}

	return Success("Token Added!")
}

// GET /api/fcm/get-token-easy
func GetToken(c *models.ReqContext) Response {
	query := models.GetTokeQuery{}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Fail get token", err)
	}

	return JSON(200, query.Result)
}

// GET /api/fcm//token-expired
func GetExpiredToken(c *models.ReqContext) Response {
	query := models.GetExpiredTokenQuery{}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Fail get expired token", err)
	}

	return JSON(200, query.Result)
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
