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

	return Success("Token Added!")
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
