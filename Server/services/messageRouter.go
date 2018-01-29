package services

// MessageRouter - MessageRouter
type MessageRouter struct {
	clients map[*Client]bool
}

// NewMessageRouter - NewMessageRouter
func NewMessageRouter() *MessageRouter {
	return &MessageRouter{
		clients: make(map[*Client]bool),
	}
}

// RegisterClient - registers new client
func (router *MessageRouter) RegisterClient(c *Client) {
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
