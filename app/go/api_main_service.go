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
	"errors"
	"net/http"

	"github.com/k-jun/waxflower/registry"
)

// MainAPIService is a service that implements the logic for the MainAPIServicer
// This service should implement the business logic for every endpoint for the MainAPI API.
// Include any external packages or services that will be required by this service.
type MainAPIService struct {
	db registry.IRegistry
}

// NewMainAPIService creates a default api service
func NewMainAPIService(db registry.IRegistry) MainAPIServicer {
	return &MainAPIService{db}
}

// ReservePut -
func (s *MainAPIService) ReservePut(ctx context.Context, ticketReserve TicketReserve) (ImplResponse, error) {
	// TODO - update ReservePut with the required logic for this service method.
	// Add api_main_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReservePut method not implemented")
}

// ResetGet -
func (s *MainAPIService) ResetGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update ResetGet with the required logic for this service method.
	// Add api_main_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	// return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ResetGet method not implemented")
}

// SearchGet -
func (s *MainAPIService) SearchGet(ctx context.Context, date string, seq int32) (ImplResponse, error) {
	// TODO - update SearchGet with the required logic for this service method.
	// Add api_main_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Ticket{}) or use other options such as http.Ok ...
	// return Response(200, []Ticket{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("SearchGet method not implemented")
}
