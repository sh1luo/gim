package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			// fmt.Println("当前连接池大小为: ", len(pool.Clients))
			for client := range pool.Clients {
				// fmt.Println(client)
				_ = client.Conn.WriteJSON(Message{
					Type:     2,
					Nickname: "",
					Body:     fmt.Sprintf("%s%d", "New User Joined...Remaining number is ", len(pool.Clients)),
				})
			}
			break
		case client := <-pool.UnRegister:
			delete(pool.Clients, client)
			//fmt.Println("当前连接池大小为: ", len(pool.Clients))
			for client := range pool.Clients {
				_ = client.Conn.WriteJSON(Message{
					Type:     2,
					Nickname: "",
					Body:     fmt.Sprintf("%s%d", "User Disconnected...Remaining number is ", len(pool.Clients)),
				})
			}
			break
		case message := <-pool.Broadcast:
			//fmt.Println("正在给所有人发送消息")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
			break
		}
	}
}
