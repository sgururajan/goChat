package main

import (
	"flag"
	"fmt"
	"goChat/Server/db"
	"goChat/Server/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userRepo := db.NewMongoUserRepository("", "goChat")
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	port := flag.Int("port", 5020, "Port number for the server to use")
	flag.Parse()

	// router.Handle("/", http.FileServer(http.Dir("../Client/dist")))
	router.HandleFunc("/api/login", authService.AuthenticateHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/register", userService.SignupHandlerWithNext(authService.AuthenticateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/api/getcontacts", authService.AuthenticationMiddleware(userService.GetContactListHandler)).Methods(http.MethodGet)

	serverAddr := fmt.Sprintf(":%d", *port)
	log.Printf("Server running at %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))

}
