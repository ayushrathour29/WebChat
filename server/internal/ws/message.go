package ws

type InboundMessage struct {
	SenderID string                 `json:"sender_id"`
	ReceiverID string               `json:"receiver_id"`
	Content string 			  `json:"content"`
	// Type string                 `json:"type"`
	// Data map[string]interface{} `json:"data"`
}

type OutboundMessage struct {
	SenderID string                 `json:"sender_id"`
	ReceiverID string               `json:"receiver_id"`
	Content string 			  `json:"content"`
}

