package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Tool"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
)

var server_index = 0

func StartCHAT(c *gin.Context) { //chat_gjson
	str_user := c.PostForm("str_user")
	str_user = GlobalData.Prompt + str_user
	str_system := GlobalData.System
	if server_index >= len(GlobalData.LocalServerList) {
		server_index = 0
	}
	CHAT := Tool.POST_FormData(map[string]string{
		"user":   str_user,
		"system": str_system,
	}, GlobalData.LocalServerList[server_index]+"/CHAT")
	println(CHAT)
	CHAT_gjson := gjson.Get(CHAT, "choices.0.message.content")
	println(CHAT_gjson.String())
	c.String(http.StatusOK, CHAT_gjson.String())
	server_index += 1

}
