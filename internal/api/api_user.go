/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"user-service/internal/users"

	"github.com/gorilla/mux"
)

// A UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiServicer
}

// NewUserApiController creates a default api controller
func NewUserApiController(s UserApiServicer) Router {
	return &UserApiController{service: s}
}

// Routes returns all of the api route for the UserApiController
func (c *UserApiController) Routes() Routes {
	return Routes{
		{
			"CreateUser",
			strings.ToUpper("Post"),
			"/api/v1/user",
			c.CreateUser,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/api/v1/user/{userId}",
			c.DeleteUser,
		},
		{
			"FindUserById",
			strings.ToUpper("Get"),
			"/api/v1/user/{userId}",
			c.FindUserById,
		},
		{
			"UpdateUser",
			strings.ToUpper("Put"),
			"/api/v1/user/{userId}",
			c.UpdateUser,
		},
	}
}

// CreateUser - Create user
func (c *UserApiController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.CreateUser(r.Context(), *user)
	// If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteUser -
func (c *UserApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := parseInt64Parameter(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.DeleteUser(r.Context(), userId)
	// If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FindUserById -
func (c *UserApiController) FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := parseInt64Parameter(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.FindUserById(r.Context(), userId)
	// If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateUser -
func (c *UserApiController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := parseInt64Parameter(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.UpdateUser(r.Context(), userId, *user)
	// If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
