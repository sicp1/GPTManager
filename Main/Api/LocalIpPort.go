package Api

import (
	"GPTManager/GlobalData"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LocalIpPort(c *gin.Context) {
	ip := c.ClientIP()
	port := c.Query("port")
	GlobalData.LocalServerList = append(GlobalData.LocalServerList, "http://"+ip+":"+port)
	println("连接成功:" + ip + ":" + port)
	c.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
