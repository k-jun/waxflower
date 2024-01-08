/*
 * WaxFlower
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 * Contact: keijun091221@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/k-jun/waxflower/go"
)

func main() {
	log.Printf("Server started")

	db := openapi.NewMemoryDB()
	GameAPIService := openapi.NewGameAPIService(db)
	GameAPIController := openapi.NewGameAPIController(GameAPIService)

	UserAPIService := openapi.NewUserAPIService(db)
	UserAPIController := openapi.NewUserAPIController(UserAPIService)

	SeatAPIService := openapi.NewSeatAPIService(db)
	SeatAPIController := openapi.NewSeatAPIController(SeatAPIService)

	TicketAPIService := openapi.NewTicketAPIService(db)
	TicketAPIController := openapi.NewTicketAPIController(TicketAPIService)

	router := openapi.NewRouter(GameAPIController, UserAPIController, SeatAPIController, TicketAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
