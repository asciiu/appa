package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	keyProto "github.com/asciiu/gomo/apikey-service/proto/apikey"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type ApiKeyController struct {
	DB     *sql.DB
	Client keyProto.ApiKeyServiceClient
}

// swagger:parameters postKey
type ApiKeyRequest struct {
	// Required.
	// in: body
	Exchange string `json:"exchange"`
	// Required.
	// in: body
	Key string `json:"key"`
	// Required.
	// in: body
	Secret string `json:"secret"`
	// Optional.
	// in: body
	Description string `json:"description"`
}

// swagger:parameters updateKey
type UpdateKeyRequest struct {
	// Required.
	// in: body
	Description string `json:"description"`
}

// A ResponseKeySuccess will always contain a status of "successful".
// swagger:model responseKeySuccess
type ResponseKeySuccess struct {
	Status string                   `json:"status"`
	Data   *keyProto.UserApiKeyData `json:"data"`
}

// A ResponseKeysSuccess will always contain a status of "successful".
// swagger:model responseKeysSuccess
type ResponseKeysSuccess struct {
	Status string                    `json:"status"`
	Data   *keyProto.UserApiKeysData `json:"data"`
}

func NewApiKeyController(db *sql.DB) *ApiKeyController {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("apikey.client"))
	service.Init()

	controller := ApiKeyController{
		DB:     db,
		Client: keyProto.NewApiKeyServiceClient("go.srv.apikey-service", service.Client()),
	}
	return &controller
}

// swagger:route GET /keys/:keyId keys getKey
//
// get a key (protected)
//
// Gets a user's key by the key ID. The secret will not be returned in the response data.
//
// responses:
//  200: responseKeySuccess "data" will contain key stuffs with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *ApiKeyController) HandleGetKey(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["jti"].(string)
	keyId := c.Param("keyId")

	getRequest := keyProto.GetUserApiKeyRequest{
		ApiKeyId: keyId,
		UserId:   userId,
	}

	r, err := controller.Client.GetUserApiKey(context.Background(), &getRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "error",
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, response)
	}

	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseKeySuccess{
		Status: "success",
		Data: &keyProto.UserApiKeyData{
			ApiKey: &keyProto.ApiKey{
				ApiKeyId:    r.Data.ApiKey.ApiKeyId,
				UserId:      r.Data.ApiKey.UserId,
				Exchange:    r.Data.ApiKey.Exchange,
				Key:         r.Data.ApiKey.Key,
				Description: r.Data.ApiKey.Description,
				Status:      r.Data.ApiKey.Status,
			},
		},
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route GET /keys keys getAllKey
//
// get all user keys (protected)
//
// Get all the user keys for this user. The api secrets will not be returned in the response data.
//
// responses:
//  200: responseKeysSuccess "data" will contain a list of key info with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *ApiKeyController) HandleListKeys(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["jti"].(string)

	getRequest := keyProto.GetUserApiKeysRequest{
		UserId: userId,
	}

	r, err := controller.Client.GetUserApiKeys(context.Background(), &getRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "error",
			Message: err.Error(),
		}

		return c.JSON(http.StatusGone, response)
	}

	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	data := make([]*keyProto.ApiKey, len(r.Data.ApiKey))
	for i := range data {
		// api removes the secret
		data[i] = &keyProto.ApiKey{
			ApiKeyId:    r.Data.ApiKey[i].ApiKeyId,
			UserId:      r.Data.ApiKey[i].UserId,
			Exchange:    r.Data.ApiKey[i].Exchange,
			Key:         r.Data.ApiKey[i].Key,
			Description: r.Data.ApiKey[i].Description,
			Status:      r.Data.ApiKey[i].Status,
		}
	}

	response := &ResponseKeysSuccess{
		Status: "success",
		Data: &keyProto.UserApiKeysData{
			ApiKey: data,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route POST /keys keys postKey
//
// add an api key (protected)
//
// Associate a new exchange api key to a user's account. Secrets will not be returned in response data.
//
// responses:
//  200: responseKeySuccess "data" will contain key info with "status": "success"
//  400: responseError missing params with "status": "fail"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *ApiKeyController) HandlePostKey(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["jti"].(string)
	addKeyRequest := ApiKeyRequest{}

	err := json.NewDecoder(c.Request().Body).Decode(&addKeyRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "fail",
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	// verify that all params are present
	if addKeyRequest.Exchange == "" || addKeyRequest.Key == "" || addKeyRequest.Secret == "" {
		response := &ResponseError{
			Status:  "fail",
			Message: "exchange, key, and secret are required!",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	createRequest := keyProto.ApiKeyRequest{
		UserId:      userId,
		Exchange:    addKeyRequest.Exchange,
		Key:         addKeyRequest.Key,
		Secret:      addKeyRequest.Secret,
		Description: addKeyRequest.Description,
	}

	r, err := controller.Client.AddApiKey(context.Background(), &createRequest)
	if err != nil {
		fmt.Println(err)
		response := &ResponseError{
			Status:  "error",
			Message: err.Error(),
		}

		return c.JSON(http.StatusGone, response)
	}

	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseKeySuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route PUT /keys/:keyId keys updateKey
//
// update a user api key (protected)
//
// The user can only update the description of an added key. The secret will not be returned.
//
// responses:
//  200: responseKeySuccess "data" will contain key info with "status": "success"
//  400: responseError missing params with "status": "fail"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *ApiKeyController) HandleUpdateKey(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["jti"].(string)
	keyId := c.Param("keyId")

	keyRequest := UpdateKeyRequest{}

	err := json.NewDecoder(c.Request().Body).Decode(&keyRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "fail",
			Message: err.Error(),
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	// client can only update description
	updateRequest := keyProto.ApiKeyRequest{
		ApiKeyId:    keyId,
		UserId:      userId,
		Description: keyRequest.Description,
	}

	r, err := controller.Client.UpdateApiKeyDescription(context.Background(), &updateRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "error",
			Message: err.Error(),
		}

		return c.JSON(http.StatusGone, response)
	}

	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseKeySuccess{
		Status: "success",
		Data:   r.Data,
	}

	return c.JSON(http.StatusOK, response)
}

// swagger:route DELETE /keys/:keyId keys deleteKey
//
// remove user api key (protected)
//
// This will remove the api key from the system.
//
// responses:
//  200: responseKeySuccess data will be null with "status": "success"
//  500: responseError the message will state what the internal server error was with "status": "error"
func (controller *ApiKeyController) HandleDeleteKey(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["jti"].(string)
	keyId := c.Param("keyId")

	removeRequest := keyProto.RemoveApiKeyRequest{
		ApiKeyId: keyId,
		UserId:   userId,
	}

	r, err := controller.Client.RemoveApiKey(context.Background(), &removeRequest)
	if err != nil {
		response := &ResponseError{
			Status:  "error",
			Message: "the apikey-service is not available",
		}

		return c.JSON(http.StatusGone, response)
	}

	if r.Status != "success" {
		response := &ResponseError{
			Status:  r.Status,
			Message: r.Message,
		}

		if r.Status == "fail" {
			return c.JSON(http.StatusBadRequest, response)
		}
		if r.Status == "error" {
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := &ResponseKeySuccess{
		Status: "success",
	}

	return c.JSON(http.StatusOK, response)
}
