package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type User struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Avatar    string `json:"avatar"`
}

type Msg struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Content string `json:"content"`
}

type Client struct {
	User
	net.Conn
}

// var Users [3]User = [
// 	User{"伊藤诚","xxxxxxxxxx",":8081/images/VVosbJSBNBn6eesWgk0YuT3nPP4hUndKYat7aJB6CmqO5dHFS1blruaKeT0tQp4n.png"},
// ]
var client map[net.Addr]net.Conn
var userMsgChan chan []byte
var sysMsgChan chan []byte

func main() {
	client = make(map[net.Addr]net.Conn)
	userMsgChan = make(chan []byte)
	sysMsgChan = make(chan []byte)

	go func() {
		for {
			select {
			case msg := <-sysMsgChan:
				SentMsgToAllUser(msg)

			case msg := <-userMsgChan:
				SentMsgToAllUser(msg)
			}
		}
	}()

	http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}

		clientAddr := conn.RemoteAddr()
		sysMsgChan <- []byte(conn.RemoteAddr().String() + "加入聊天室")
		client[clientAddr] = conn

		go func() {
			defer func() {
				fmt.Println(clientAddr.String() + "退出聊天室")
				delete(client, clientAddr)
				conn.Close()
			}()
			for {
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
					break
				}
				userMsgChan <- msg
			}
		}()
	}))
}

func SentMsgToAllUser(msg []byte) {
	for _, x := range client {
		msg := Msg{Name: "伊藤诚", Avatar: "./assets/VVosbJSBNBn6eesWgk0YuT3nPP4hUndKYat7aJB6CmqO5dHFS1blruaKeT0tQp4n.png", Content: string(msg)}
		json, _ := json.Marshal(msg)
		wsutil.WriteServerText(x, json)
	}
}
