package Api

import (
	"GPTManager/GlobalData"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetSystem(c *gin.Context) {
	system := c.PostForm("System")
	GlobalData.System = system
	println("设置system成功")
	c.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
func GetSystem(c *gin.Context) {
	c.String(http.StatusOK, GlobalData.System)
}
