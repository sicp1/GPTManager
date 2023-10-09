package Api

import (
	"GPTManager/GlobalData"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setprompt(c *gin.Context) {
	GlobalData.Prompt = c.PostForm("prompt")
	println("prompt :" + GlobalData.Prompt)
	c.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
