package websocket

import (
	"bytes"
	"encoding/json"
	"time"
)

func (c *Client) readMessage(link string) {
	//function read the message sent from client
	//read it from the other client

	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(data string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		wsMsg := &WsMessage{}
		reader := bytes.NewReader(message)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(wsMsg)

		if err != nil {
			break
		}

		c.hub.brodcast <- &Message{Client_id: c.user_id, Message: wsMsg.Message, Link: link}
	}
}
