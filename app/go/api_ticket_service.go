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

// TicketAPIService is a service that implements the logic for the TicketAPIServicer
// This service should implement the business logic for every endpoint for the TicketAPI API.
// Include any external packages or services that will be required by this service.
type TicketAPIService struct {
	db registry.IRegistry
}

// NewTicketAPIService creates a default api service
func NewTicketAPIService(db registry.IRegistry) TicketAPIServicer {
	return &TicketAPIService{db}
}

// TicketsPost -
func (s *TicketAPIService) TicketsPost(ctx context.Context, ticket Ticket) (ImplResponse, error) {
	mt := model.Ticket{}
	err := copier.Copy(&mt, &ticket)
	if err != nil {
		return Response(400, nil), err
	}
	_, err = s.db.InsertTicket(&mt)
	if err != nil {
		return Response(400, nil), err
	}
	return Response(200, nil), err
}

// TicketsTicketIdGet -
func (s *TicketAPIService) TicketsTicketIdGet(ctx context.Context, ticketId string) (ImplResponse, error) {
	mt := &model.Ticket{Id: ticketId}
	mt, err := s.db.SelectTicket(mt)
	if err != nil {
		return Response(400, nil), err
	}
	t := &Ticket{}
	err = copier.Copy(t, mt)
	if err != nil {
		return Response(400, nil), err
	}
	return Response(200, nil), err
}
