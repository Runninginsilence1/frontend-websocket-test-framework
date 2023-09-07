package service

import (
	"frontend-websocket-test-framework/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouterList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    []string{"R1C"},
		})
	}
}

func GetUsableList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data": []common.DeviceInfo{
				{
					Device:     "sda1",
					Mountpoint: "/mnt1",
					Fstype:     "设备",
					Total:      "25GB",
					Free:       "17GB",
					Used:       "5GB",
				},
				{
					Device:     "sda2",
					Mountpoint: "/mnt2",
					Fstype:     "设备",
					Total:      "25GB",
					Free:       "17GB",
					Used:       "5GB",
				},
				{
					Device:     "sda3",
					Mountpoint: "/mnt3",
					Fstype:     "设备",
					Total:      "25GB",
					Free:       "17GB",
					Used:       "5GB",
				},
				{
					Device:     "sda4",
					Mountpoint: "/mnt4",
					Fstype:     "设备",
					Total:      "2TB",
					Free:       "1.75TB",
					Used:       "221GB",
				},
			},
		})
	}
}
