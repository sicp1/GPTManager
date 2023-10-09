package Api

import (
	"GPTManager/GlobalData"
	"GPTManager/Implement"
	"github.com/gin-gonic/gin"
)

func Init(c *gin.Context) {
	GlobalData.ApiKeyQueue_Main.Data = Implement.ApiKeyArrToInterfaceArr(GlobalData.ApiKey_WithStatus)
	println("main初始化完成")
}
