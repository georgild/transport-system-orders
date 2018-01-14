package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Route : defines route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : defines array of routes
type Routes []Route

// NewRouter : returns new router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend")))
	return router
}

var routes = Routes{
	Route{
		"PostOrder",
		"POST",
		"/api/v1/orders",
		PostOrder,
	},
	Route{
		"GetSeats",
		"GET",
		"/api/v1/orders/routes/{routeID}/reservedseats",
		GetSeats,
	},
}

// PostOrder :
func PostOrder(w http.ResponseWriter, r *http.Request) {

	var order Order
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &order); err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	log.Print("ABOUT TO CREATE GAME...")

	created, insertError := CreateOrder(order)

	if insertError != nil {
		log.Panic(insertError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		panic(err)
	}
}

// GetSeats :
func GetSeats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	routeID := vars["routeID"]

	var seats []Seat

	log.Print("ABOUT TO CREATE TRY...")

	seats, tryError := GetReservedSeats(routeID)

	if tryError != nil {
		panic(tryError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(seats); err != nil {
		panic(err)
	}
}
