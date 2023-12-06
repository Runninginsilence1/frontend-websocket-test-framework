package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/olahol/melody"
	"github.com/spf13/cobra"
	"gowscat/common"
	"gowscat/utils"
	"log"
	"net/http"
	"time"
)

var serverPort int

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s", "s"},
	Short:   "Launch gin and websocket server",
	Long:    "Launch gin and websocket server",
	Example: "-M GET|POST",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre run method command")
	},
	Run: func(cmd *cobra.Command, args []string) {
		GinServer(fmt.Sprintf(":%v", serverPort))
	},
}

func init() {
	serverCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "server port")
}

// 通过读取配置文件的方式，来做成类似于 postman websocket调试的那种效果；
func GinServer(port string) {
	r := gin.Default()

	// 普通接口

	m := NewMelody()
	_ = m

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
	//r.Run(global.ServerPort)
	r.Run(port)
}

func NewMelody() *melody.Melody {
	m := melody.New()
	// configuration melody
	melodyCfg := &melody.Config{
		WriteWait:  10 * time.Second,
		PongWait:   1 * time.Second,
		PingPeriod: (1 * time.Second * 9) / 10,
		// MaxMessageSize 这个值可以设置大一点
		MaxMessageSize:    2048, // 坑：如果json的数据多的话会直接down掉该次websocket session; default 512
		MessageBufferSize: 1024,
	}

	m.Config = melodyCfg

	m.HandleConnect(func(s *melody.Session) {
	})

	// 是显示呢还是覆盖到文件中呢？
	m.HandleMessage(func(s *melody.Session, rawBytes []byte) {

	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Printf("ws client %v disconnected\n", s.RemoteAddr())
	})

	m.HandleClose(func(session *melody.Session, i int, s string) error {
		fmt.Println("hook close")
		fmt.Println("string:", s)
		fmt.Println("int:", i)
		return nil
	})

	return m
}

func ShowMsg() {
	addr := "localhost"
	msg := ""

	fmt.Printf(`%v:
%v`, addr, msg)
}
