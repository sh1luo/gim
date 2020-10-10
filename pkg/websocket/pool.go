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
		Register:   make(chan *Client, 50),
		UnRegister: make(chan *Client, 20),
		Clients:    make(map[*Client]bool, 50),
		Broadcast:  make(chan Message, 100),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			for c := range pool.Clients {
				_ = c.Conn.WriteJSON(Message{
					Type:     2,
					Nickname: "",
					Body: fmt.Sprintf("%s%s%s%d",
						"欢迎新用户：", client.Nickname, " 加入...当前房间用户数： ",
						len(pool.Clients)),
				})
			}
			break
		case client := <-pool.UnRegister:
			delete(pool.Clients, client)
			//fmt.Println("当前连接池大小为: ", len(pool.Clients))
			for c := range pool.Clients {
				_ = c.Conn.WriteJSON(Message{
					Type:     2,
					Nickname: "",
					Body: fmt.Sprintf("%s%d",
						"用户", c.Nickname, "已经退出群聊，剩余人数为： ",
						len(pool.Clients)),
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
