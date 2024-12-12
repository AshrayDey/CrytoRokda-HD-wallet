package routes

import (
	"HD-Wallet/utils"

	"github.com/gin-gonic/gin"
)

func testFunc(context *gin.Context) {
	res := utils.GenerateMneunonic("Test")
	context.JSON(200, res)
}
