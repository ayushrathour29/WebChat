package ws

type Hub struct{
	Clients map[* Client]bool
	Broadcast chan OutboundMessage
	Register chan * Client
	Unregister chan * Client
}

func NewHub() *Hub{
	return &Hub{
		Clients: make(map[*Client]bool),
		Broadcast: make(chan OutboundMessage),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
	}
}


// infine loop with select statement to handle register, unregister and broadcast
func (h *Hub) Run(){
	for{
		select{
		case client := <- h.Register:
			h.Clients[client] = true
		case client := <- h.Unregister:
			if _,ok:= h.Clients[client]; ok{
				delete(h.Clients, client)
				close(client.Send)
			}
		case msg := <- h.Broadcast:
			for client := range h.Clients{
				if client.ID == msg.ReceiverID || client.ID == msg.SenderID{
					client.Send <- msg
				}
			}
		
	}

}
}