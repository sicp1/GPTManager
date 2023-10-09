package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateApi(c *gin.Context) {
	nowURL := c.PostForm("URL")
	c.String(http.StatusOK, "http://"+nowURL+"/StartCHAT")
	println("API : " + "http://" + nowURL + "/StartCHAT")
}
