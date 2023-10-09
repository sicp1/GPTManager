package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetApiKey(c *gin.Context) {
	apikey_num := c.PostForm("ApiKey")
	apikey_status := c.PostForm("Status")
	apikey_interface := Implement.ApiKey{ApiKeyNum: apikey_num, Status: apikey_status}
	GlobalData.ApiKeyQueue_Local.Push(apikey_interface)
	println("apikey 推入成功")
	//构建channel
	GlobalData.LocalChannelWithStatus.Mutex.Lock()
	GlobalData.LocalChannelWithStatus.ChannelWithStatus = append(GlobalData.LocalChannelWithStatus.ChannelWithStatus, map[string]interface{}{
		"ch_user":   make(chan string, 1),
		"ch_system": make(chan string, 1),
		"ch_return": make(chan string, 1),
		"status":    0, //1代表正在连接，0代表准备中，-1代表销毁
	})
	//	last_index := len(GlobalVar.LocalChannelWithStatus.ChannelWithStatus) - 1

	GlobalData.LocalChannelWithStatus.Mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
