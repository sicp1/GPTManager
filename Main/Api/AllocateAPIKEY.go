package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"GPTManager/Tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllocateAPIKEY(c *gin.Context) {
	var i = 0
	for {

		tmp_ApiKey, flag := GlobalData.ApiKeyQueue_Main.Pop()
		apikey_withstatus := Implement.InterfaceToApiKey(tmp_ApiKey)
		Tool.POST_FormData(map[string]string{
			"ApiKey": apikey_withstatus.ApiKeyNum,
			"Status": apikey_withstatus.Status,
		}, GlobalData.LocalServerList[i]+"/GetApiKey")
		if flag == 0 {
			break
		}
		i += 1
		if i >= len(GlobalData.LocalServerList) {
			i = 0
		}

	}
	c.String(http.StatusOK, "完成！")
}
