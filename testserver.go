package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/greenac/testserver/endpoints"
	"github.com/greenac/testserver/endpointcontrollers"
)

const hostAndPort = ":7979"
var userEndPoints = endpoints.UserEndpoints{}
var userController = endpointcontrollers.UserController{}
var dummyDataController = endpointcontrollers.DummyDataController{}

func main() {
	userController.Initialize()
	dummyDataController.Initialize()
	fmt.Println("Starting Test Server on", hostAndPort)
	r := mux.NewRouter()
	r.HandleFunc(userEndPoints.SaveUser(), userController.SaveUser)
	r.HandleFunc(dummyDataController.EndPoints.GetSenatorUri(), dummyDataController.GetSenators)
	r.HandleFunc("/", handle404)
	log.Fatal(http.ListenAndServe(hostAndPort, r))
}

func handle404(w http.ResponseWriter, r *http.Request) {
	fmt.Println("No url setup for:", r.URL.Path)
	fmt.Fprintf(w, "No url setup for: %s", r.URL.Path)
}
