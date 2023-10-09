package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Status    string //-1报错状态   0 失效状态   1 正常状态
func AppendApiKey(c *gin.Context) {
	apikey := c.PostForm("ApiKey")
	GlobalData.ApiKey_WithStatus = append(GlobalData.ApiKey_WithStatus, Implement.ApiKey{
		ApiKeyNum: apikey,
		Status:    "1",
	})
	c.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
