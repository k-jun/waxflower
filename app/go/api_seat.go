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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// SeatAPIController binds http requests to an api service and writes the service results to the http response
type SeatAPIController struct {
	service SeatAPIServicer
	errorHandler ErrorHandler
}

// SeatAPIOption for how the controller is set up.
type SeatAPIOption func(*SeatAPIController)

// WithSeatAPIErrorHandler inject ErrorHandler into controller
func WithSeatAPIErrorHandler(h ErrorHandler) SeatAPIOption {
	return func(c *SeatAPIController) {
		c.errorHandler = h
	}
}

// NewSeatAPIController creates a default api controller
func NewSeatAPIController(s SeatAPIServicer, opts ...SeatAPIOption) Router {
	controller := &SeatAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SeatAPIController
func (c *SeatAPIController) Routes() Routes {
	return Routes{
		"SeatsPost": Route{
			strings.ToUpper("Post"),
			"/seats",
			c.SeatsPost,
		},
		"SeatsSeatIdGet": Route{
			strings.ToUpper("Get"),
			"/seats/{seatId}",
			c.SeatsSeatIdGet,
		},
	}
}

// SeatsPost - 
func (c *SeatAPIController) SeatsPost(w http.ResponseWriter, r *http.Request) {
	seatParam := Seat{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&seatParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSeatRequired(seatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSeatConstraints(seatParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SeatsPost(r.Context(), seatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SeatsSeatIdGet - 
func (c *SeatAPIController) SeatsSeatIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seatIdParam := params["seatId"]
	if seatIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"seatId"}, nil)
		return
	}
	result, err := c.service.SeatsSeatIdGet(r.Context(), seatIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
