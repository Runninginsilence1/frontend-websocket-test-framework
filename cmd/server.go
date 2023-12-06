package cmd

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var serverPort int
var senderFilePath string

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
	serverCmd.Flags().StringVarP(&senderFilePath, "msg", "m", "message.json", "websocket broadcast message path")
}

// 通过读取配置文件的方式，来做成类似于 postman websocket调试的那种效果；
func GinServer(port string) {
	r := gin.Default()
	// 普通接口
	m := NewMelody()
	go func() {
		var err error
		_ = err
		var input string
		for {
			_, _ = fmt.Scan(&input)
			if input == "send" {
				BroadcastMsg(m, senderFilePath)
			}
		}
	}()
	r.GET("/ws", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			log.Fatalln("melody handler request error:", err)
		}
	})
	_ = r.Run(port)
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

	// 直接显示内容
	m.HandleMessage(func(s *melody.Session, rawBytes []byte) {
		addr := s.RemoteAddr()
		msg := string(rawBytes)

		fmt.Printf(`%v says:
%v`, addr, msg)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Printf("ws client %v disconnected\n", s.RemoteAddr())
	})

	m.HandleClose(func(session *melody.Session, i int, s string) error {
		return nil
	})

	return m
}

// 向websocket中发送信息
func BroadcastMsg(m *melody.Melody, filePath string) {
	content, _ := fileutil.ReadFileToString(filePath)
	_ = m.Broadcast([]byte(content))
}
