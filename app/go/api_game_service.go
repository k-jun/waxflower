/*
 * WaxFlower
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 * Contact: keijun091221@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/registry"
)

// GameAPIService is a service that implements the logic for the GameAPIServicer
// This service should implement the business logic for every endpoint for the GameAPI API.
// Include any external packages or services that will be required by this service.
type GameAPIService struct {
	db registry.IRegistry
}

// NewGameAPIService creates a default api service
func NewGameAPIService(db registry.IRegistry) GameAPIServicer {
	return &GameAPIService{db}
}

// GamesGameIdGet -
func (s *GameAPIService) GamesGameIdGet(ctx context.Context, gameId string) (ImplResponse, error) {
	mg := &model.Game{Id: gameId}
	mg, err := s.db.SelectGame(mg)
	if err != nil {
		return Response(400, err), nil
	}
	g := &Game{}
	err = copier.Copy(g, mg)
	if err != nil {
		return Response(400, err), nil
	}
	return Response(200, g), nil
}

// GamesPost -
func (s *GameAPIService) GamesPost(ctx context.Context, game Game) (ImplResponse, error) {
	mg := &model.Game{}
	err := copier.Copy(mg, &game)
	if err != nil {
		return Response(400, nil), err
	}
	_, err = s.db.InsertGame(mg)
	if err != nil {
		return Response(400, nil), err
	}
	return Response(200, nil), err
}
