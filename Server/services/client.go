package services

import (
	"goChat/Server/utils"
	"goChat/Server/viewModels"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// Client - Client
type Client struct {
	router  *MessageRouter
	conn    *websocket.Conn
	send    chan viewModels.Message
	OwnerID string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// AddClient - AddClient
func AddClient(mRouter *MessageRouter, w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.JSONInternalServerErrorResponse(w, nil)
		return
	}
	client := Client{
		conn:   connection,
		router: mRouter,
		send:   make(chan viewModels.Message),
	}

	mRouter.register <- &client

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.router.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// start the infinite loop
	for {
		var msg viewModels.Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseUnsupportedData) {
				log.Printf("Websocket Error: %v", err)
			} else {
				log.Fatalf("Websocket Error: %v", err)
			}
			break
		}
		msg.TimeStamp = time.Now()
		msg.Append = true

		c.router.routeMessage <- msg

	}
}

func (c *Client) writePump() {
	pingTimer := time.NewTicker(pingPeriod)
	defer func() {
		pingTimer.Stop()
		c.conn.Close()
	}()

	// start the infinite loop
	for {
		select {
		case _, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
		case <-pingTimer.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
