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

	if len(query.Result) <= 0 {
		return JSON(200, "no expired token need to delete")
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

// Delete /delete-expired-token
func (hs *HTTPServer) DeleteExpiredToken(c *models.ReqContext) Response {
	query := models.GetExpiredTokenQuery{}
	// check if we have expired token in db before delete
	if err := bus.Dispatch(&query); err != nil {
		log.Fatal(err)
		return Error(500, "Fail get expired token query", err)
	}

	if len(query.Result) <= 0 {
		return Success("no expired token need to delete")
	}
	// send delete command ....
	cmd := models.DeleteExpiredTokenCommand{}
	if err := bus.Dispatch(&cmd); err != nil {
		log.Fatal(err)
		return Error(500, "fail to delete the expired tokens", err)
	}

	return Success("expired token deleted")
}
