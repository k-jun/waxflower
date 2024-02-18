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
	"net/http"
	"strings"
)

// MainAPIController binds http requests to an api service and writes the service results to the http response
type MainAPIController struct {
	service      MainAPIServicer
	errorHandler ErrorHandler
}

// MainAPIOption for how the controller is set up.
type MainAPIOption func(*MainAPIController)

// WithMainAPIErrorHandler inject ErrorHandler into controller
func WithMainAPIErrorHandler(h ErrorHandler) MainAPIOption {
	return func(c *MainAPIController) {
		c.errorHandler = h
	}
}

// NewMainAPIController creates a default api controller
func NewMainAPIController(s MainAPIServicer, opts ...MainAPIOption) Router {
	controller := &MainAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MainAPIController
func (c *MainAPIController) Routes() Routes {
	return Routes{
		"ResetGet": Route{
			strings.ToUpper("Get"),
			"/reset",
			c.ResetGet,
		},
	}
}

// ResetGet -
func (c *MainAPIController) ResetGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ResetGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
