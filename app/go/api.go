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
	"net/http"
)



// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface { 
	UsersPost(http.ResponseWriter, *http.Request)
	UsersUserIdGet(http.ResponseWriter, *http.Request)
	UsersUserIdPut(http.ResponseWriter, *http.Request)
}


// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface { 
	UsersPost(context.Context, User) (ImplResponse, error)
	UsersUserIdGet(context.Context, string) (ImplResponse, error)
	UsersUserIdPut(context.Context, string, User) (ImplResponse, error)
}