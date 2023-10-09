package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// "ch_user":   make(chan string, 1),
// "ch_system": make(chan string, 1),
// "ch_return": make(chan string,1),
// "status":    0, //1代表正在连接，0代表准备中，-1代表销毁
var chat_index = 0
var mutex sync.Mutex

func CHAT(c *gin.Context) {
	//Localmain
	user := c.PostForm("user")
	system := c.PostForm("system")
	mutex.Lock()
	for {
		if chat_index >= len(GlobalData.LocalChannelWithStatus.ChannelWithStatus) {
			chat_index = 0
		}
		if GlobalData.LocalChannelWithStatus.ChannelWithStatus[chat_index]["status"] == 1 {
			ch_system := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[chat_index]["ch_system"])
			ch_user := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[chat_index]["ch_user"])
			ch_return := Tool.InterfaceToChannel_string(GlobalData.LocalChannelWithStatus.ChannelWithStatus[chat_index]["ch_return"])
			ch_system <- system
			ch_user <- user
			chat_index += 1
			c.String(http.StatusOK, <-ch_return)
			break
		}
		chat_index += 1
	}
	mutex.Unlock()

}
