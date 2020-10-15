package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/url"
	"strconv"
	"time"
)

var (
	MaxUserNums   int           // 最大登陆人数
	LoginInterval time.Duration // 登陆间隔
	MsgInterval   time.Duration // 每条消息发送间隔
	//Host          string        // 目标服务器地址+端口
	connectUrl url.URL // 构建的 websocket 连接
)

func init() {
	flag.IntVar(&MaxUserNums, "u", 100, "")
	flag.DurationVar(&LoginInterval, "l", 1e3, "")
	flag.DurationVar(&MsgInterval, "m", 20*time.Second, "")
	//flag.StringVar(&Host, "h", "127.0.0.1:8080", "")
	connectUrl = url.URL{
		Scheme:     "ws",
		Host:       "39.106.55.151:8000",
		Path:       "/",
		ForceQuery: true,
	}
}

func main() {
	flag.Parse()

	for i := 0; i < MaxUserNums; i++ {
		go AddOneFakeUser("user" + strconv.Itoa(i))
		time.Sleep(LoginInterval)
	}

	select {}
}

func AddOneFakeUser(username string) {
	u := url.Values{"username": {username}}
	connectUrl.RawQuery = u.Encode()
	conn, _, err := websocket.DefaultDialer.Dial(connectUrl.String(), nil)
	if err != nil {
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(MsgInterval)
	defer ticker.Stop()
	done := make(chan bool)

	msgToSends := make(map[string]string,1)
	msgToSends["content"] = "This is a message from :" + username

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if err := conn.WriteJSON(msgToSends); err != nil {
				continue
			}
		}
	}
}
