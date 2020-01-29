package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type Client struct {
	Nickname string
	Conn     *websocket.Conn
	Pool     *Pool
}

type Message struct {
	Type     int    `json:"type"`
	Nickname string `json:"nickname"`
	Body     string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.UnRegister <- c
		_ = c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		msg := json.Get(p, "content").ToString()

		message := Message{
			Type:     1,
			Nickname: c.Nickname,
			Body:     msg,
		}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}