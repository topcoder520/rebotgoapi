package router

import (
	"github.com/gin-gonic/gin"
	"github.com/topcoder520/rebotgoapi/mouse"
)

func GetRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//鼠标操作
	v1 := router.Group("/mouse")
	{
		v1.GET("/left", mouse.MouseMoveLeft)
		v1.GET("/right", mouse.MouseMoveRight)
		v1.GET("/up", mouse.MouseMoveUp)
		v1.GET("/down", mouse.MouseMoveDown)
		v1.GET("/move", mouse.MouseMove)
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
	return router
}
