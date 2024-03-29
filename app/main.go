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

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	openapi "github.com/k-jun/waxflower/go"
	"github.com/k-jun/waxflower/registry"

	"github.com/gorilla/handlers"
)

func main() {
	log.Printf("Server started")

	conf := &mysql.Config{
		Addr:   "127.0.0.1:3306",
		User:   "mysql",
		Passwd: "mysql",
		DBName: "mydb",
	}
	db := sqlx.MustConnect("mysql", conf.FormatDSN())
	sql := registry.NewMySQL(db)

	GameAPIService := openapi.NewGameAPIService(sql)
	GameAPIController := openapi.NewGameAPIController(GameAPIService)

	SeatAPIService := openapi.NewSeatAPIService(sql)
	SeatAPIController := openapi.NewSeatAPIController(SeatAPIService)

	TicketAPIService := openapi.NewTicketAPIService(sql)
	TicketAPIController := openapi.NewTicketAPIController(TicketAPIService)

	UserAPIService := openapi.NewUserAPIService(sql)
	UserAPIController := openapi.NewUserAPIController(UserAPIService)

	MainAPIService := openapi.NewMainAPIService(sql)
	MainAPIController := openapi.NewMainAPIController(MainAPIService)

	router := openapi.NewRouter(GameAPIController, SeatAPIController, TicketAPIController, UserAPIController, MainAPIController)

	// log.Fatal(http.ListenAndServe(":8080", router))

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)))

}
