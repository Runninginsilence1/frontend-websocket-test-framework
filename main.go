package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"risk/websocket/common"
	"risk/websocket/global"
	. "risk/websocket/router"
	"risk/websocket/utils"
)

// 通过读取配置文件的方式，来做成类似于 postman websocket调试的那种效果；
func main() {
	r := gin.Default()
	// 普通接口
	Router(r)

	// ws服务： 读取配置文件来进行send: send.json里面放一个测试用的数据；
	upgrade := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}

	r.GET("/ws", func(c *gin.Context) {
		defer c.JSON(200, "")
		wsConn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket升级失败：%v", err)
			c.JSON(401, common.R{
				Code:    401,
				Message: "无法启动webSocket协议。请更换最新版本Chrome浏览器。",
				Data:    nil,
			})
			return
		}

		// 开一个go协程，用来处理发送数据；
		sendExit := make(chan struct{})
		fmt.Printf("敲击Enter发送数据:\n")
		go func() {
			var payload = &common.R{}
			var input string
			for {
				// 从控制台读取信息，判断是
				select {
				case <-sendExit:
					log.Printf("协程已退出\n")
					return
				default:
					_, err = fmt.Scan(&input)
					if err != nil {
						continue
					}
					utils.ReadSendMsg(payload)
					err = wsConn.WriteJSON(payload)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
		}()

		// 这里接收到后开始写
		var msgBytes []byte
		for {
			_, msgBytes, err = wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("[Error] func (conn *WebSocket)ReadJson!\nErr: %v\n", err)
				sendExit <- struct{}{}
				return
			}
			fmt.Println("打印接收到的数据：")
			fmt.Println(string(msgBytes))
		}
	})
	r.Run(global.ServerPort)
}
