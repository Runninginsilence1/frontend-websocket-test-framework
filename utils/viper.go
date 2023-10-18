package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"risk/websocket/global"
)

func init() {
	err := LoadConfig()
	if err != nil {
		fmt.Println("[ERROR] LoadConfig")
		return
	}
}

func ReadSendMsg(data any) {
	viper.SetConfigFile("./message.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("读取配置文件出现错误: %v", err)
		return
	}
	err = viper.Unmarshal(data)
	if err != nil {
		log.Printf("解析配置文件时出现错误 %v", err)
		return
	}
}

func LoadConfig() error {
	viper.SetConfigName("config") // 配置文件的文件名（不需要扩展名）
	viper.SetConfigType("yaml")   // 配置文件的类型（可以是 yaml、json、toml 等）
	viper.AddConfigPath("./")     // 可选：指定配置文件的搜索路径（默认为当前目录）

	err := viper.ReadInConfig() // 加载配置文件
	if err != nil {
		fmt.Println("Load config.yaml failed, check err output.")
		return err
	}

	// Put your global variable here:
	global.ServerPort = ":" + viper.GetString("serverPort")
	fmt.Println("serverPort from config is ", global.ServerPort)
	return nil
}

// need type assertion
//func ViperMustGetKey[T any](key string) T {
//	if !viper.IsSet(key) {
//		panic(fmt.Sprintf("%v is invalid key! Please check file: config.yaml", key))
//	}
//	return viper.Get(key)
//}
