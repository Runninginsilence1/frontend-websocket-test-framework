package utils

import (
	"github.com/spf13/viper"
	"log"
)

func ReadJSON(data any) {
	viper.SetConfigFile("./send1.json")
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
