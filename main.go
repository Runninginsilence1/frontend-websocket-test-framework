package main

import (
	"encoding/json"
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
		sendExit := make(chan bool)
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
					fmt.Print("请输入你的指令\n")
					_, err = fmt.Scanln(&input)
					if err != nil {
						continue
					}
					if input != "send" {
						log.Printf("指令有误，请重新输入")
					} else { // 使用类似于广播的形式把配置文件里面的数据读进来然后发送出去；
						utils.ReadJSON(payload)
						err = wsConn.WriteJSON(payload)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
		}()

		// 这里接收到后开始写
		payload := &common.R{}
		for {
			err = wsConn.ReadJSON(payload)
			if err != nil {
				log.Printf("%v可能前端已经断开ws连接，或者我这里json设置的数据不对", err)
				sendExit <- true
				return
			}
			var R, _ = json.Marshal(payload)
			msg := string(R)
			fmt.Println("打印一下payload")
			fmt.Println(msg)
			//wsConn.WriteJSON(payload)
		}
	})
	r.Run(global.ServerPort)
}
