package websocket

func CreateHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		brodcast:   make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()

		case client := <-h.unregister:
			h.Lock()
			if _, exist := h.clients[client]; exist {
				close(client.send)
				delete(h.clients, client)
			}
			h.Unlock()

		case msg := <-h.brodcast:
			h.RLock()

			h.messages = append(h.messages, msg)

			for client := range h.clients {
				if client.user_id == msg.Client_id {
					continue
				}

				select {

				case client.send <- getMessageTemplate(msg):

				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.RUnlock()

		}
	}
}
