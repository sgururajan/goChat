package services

import (
	"goChat/Server/models"
	"goChat/Server/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client - Client
type Client struct {
	router  *MessageRouter
	conn    *websocket.Conn
	send    chan models.Message
	OwnerID string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// AddClient - AddClient
func (c *Client) AddClient(mRouter *MessageRouter, w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.JSONInternalServerErrorResponse(w, nil)
		return
	}
	client := Client{
		conn:   connection,
		router: mRouter,
		send:   make(chan models.Message),
	}

	mRouter.RegisterClient(&client)
}
