package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/topcoder520/rebotgoapi/keyboard"
	"github.com/topcoder520/rebotgoapi/mouse"
	"github.com/topcoder520/rebotgoapi/util"
)

func GetRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(Cors()) //解决跨域

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, util.Success())
	})

	//鼠标操作
	v1 := router.Group("/mouse")
	{
		v1.GET("/left", mouse.MouseMoveLeft)
		v1.GET("/right", mouse.MouseMoveRight)
		v1.GET("/up", mouse.MouseMoveUp)
		v1.GET("/down", mouse.MouseMoveDown)
		v1.GET("/move", mouse.MouseMove)
		v1.GET("/move2", mouse.MouseMove2)
		v1.GET("/leftclick", mouse.MouseLeftClick)
		v1.GET("/leftdoubleclick", mouse.MouseLeftDoubleClick)
		v1.GET("/rightclick", mouse.MouseRightClick)

		v1.GET("/lefttoggle", mouse.MouseLeftToggle)
		v1.GET("/lefttoggleup", mouse.MouseLeftToggleUp)
		v1.GET("/righttoggle", mouse.MouseRightoggle)
		v1.GET("/righttoggleup", mouse.MouseRightoggleUp)

		v1.GET("/centerclick", mouse.MouseCenterClick)
		v1.GET("/wheeldown", mouse.MouseWheelDown)
		v1.GET("/wheelup", mouse.MouseWheelUp)
		v1.GET("/wheelleft", mouse.MouseWheelLeft)
		v1.GET("/wheelright", mouse.MouseWheelRight)
		v1.GET("/wheelscrollup", mouse.MouseScrollMouseUp)
		v1.GET("/wheelscrolldown", mouse.MouseScrollMouseDown)
	}
	//键盘操作
	v2 := router.Group("/key")
	{
		v2.POST("/Keyinputstring", keyboard.KeyInputString)
		v2.GET("/keytap", keyboard.KeyTap)
	}
	return router
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
