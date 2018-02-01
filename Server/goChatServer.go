package main

import (
	"flag"
	"fmt"
	"goChat/Server/db"
	"goChat/Server/inMemoryDatabase"
	"goChat/Server/mongo"
	"goChat/Server/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	port := flag.Int("port", 5020, "Port number for the server to use")
	inMemoryDb := flag.Bool("inmemory", false, "Flag to use In-Memory database. Default is false")
	flag.Parse()

	log.Printf("Using inmemory db: %v", *inMemoryDb)

	var userRepo db.IUserRepository
	var messageRepo db.IMessageRepository
	var convRepo db.IConversationRepository

	if *inMemoryDb {
		userRepo = inMemoryDatabase.NewUserRepository()
		messageRepo = inMemoryDatabase.NewMessageRepository()
		convRepo = inMemoryDatabase.NewConversationRepository()
	} else {
		userRepo = mongo.NewUserRepository("", "goChat")
		messageRepo = mongo.NewMessageRepository("", "goChat")
		convRepo = mongo.NewConversationRepository("", "goChat")
	}

	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	messageService := services.NewMessageService(messageRepo)
	convService := services.NewConversationService(convRepo, userService)

	mRouter := services.NewMessageRouter(userService, convService, messageService)

	// router.Handle("/", http.FileServer(http.Dir("../Client/dist")))
	router.HandleFunc("/api/login", authService.AuthenticateHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/register", userService.SignupHandlerWithNext(authService.AuthenticateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/api/conversations", authService.AuthenticationMiddleware(convService.GetConversationHandler)).Methods(http.MethodGet)
	router.HandleFunc("/api/getmessages/{conversationID}/{page}/{count}", authService.AuthenticationMiddleware(messageService.GetMessageHandler)).Methods(http.MethodGet)
	router.HandleFunc("/api/updatemessageasread/{messageId}", authService.AuthenticationMiddleware(messageService.UpdateMessageAsRead)).Methods(http.MethodPost)

	router.HandleFunc("/api/ws", authService.AuthenticationMiddleware(func(w http.ResponseWriter, r *http.Request) {
		services.AddClient(mRouter, w, r)
	}))

	go mRouter.Run()

	serverAddr := fmt.Sprintf(":%d", *port)
	log.Printf("Server running at %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))

}
