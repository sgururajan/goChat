package services

import (
	"goChat/Server/models"
	"goChat/Server/viewModels"
	"log"
)

// MessageRouter - MessageRouter
type MessageRouter struct {
	clients        map[*Client]bool
	register       chan *Client
	unregister     chan *Client
	routeMessage   chan viewModels.Message
	convService    *ConversationService
	userService    *UserService
	messageService *MessageService
}

// NewMessageRouter - NewMessageRouter
func NewMessageRouter(userSvc *UserService, convSvc *ConversationService, msgSvc *MessageService) *MessageRouter {
	return &MessageRouter{
		clients:        make(map[*Client]bool),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		routeMessage:   make(chan viewModels.Message),
		userService:    userSvc,
		convService:    convSvc,
		messageService: msgSvc,
	}
}

// RegisterClient - registers new client
func (router *MessageRouter) registerClient(c *Client) {
	found := false
	for client := range router.clients {
		if client.OwnerID == c.OwnerID {
			found = true
			break
		}
	}

	if found {
		return
	}

	router.clients[c] = true
}

// Run - runs the messages router
func (router *MessageRouter) Run() {
	// start the infinite loop
	for {
		select {
		case client := <-router.register:
			router.registerClient(client)
		case client := <-router.unregister:
			if _, ok := router.clients[client]; ok {
				delete(router.clients, client)
				close(client.send)
			}
		case msg := <-router.routeMessage:
			conv, err := router.convService.GetConversationByID(msg.ConversationID)
			if err != nil {
				log.Printf("Error getting conversation: %v", err)
				continue
			}

			var partcipantState []models.MessageParicipantState
			for _, p := range conv.Participants {
				partcipantState = append(partcipantState, models.MessageParicipantState{ParticipantID: p, IsRead: p == msg.AuthorID})
			}

			dbMsg := models.Message{
				ConversationID:    msg.ConversationID,
				Author:            msg.AuthorID,
				Body:              msg.Body,
				ParticipantsState: partcipantState,
			}

			router.messageService.AddMessage(dbMsg)

			for _, s := range conv.Participants {
				for c := range router.clients {
					if c.OwnerID == s {
						c.send <- msg
					}
				}
			}
		}
	}
}
