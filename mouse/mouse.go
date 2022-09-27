package mouse

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/topcoder520/rebotgoapi/util"
)

func init() {
	robotgo.MouseSleep = 100
}

//MouseMoveLeft 鼠标左移
func MouseMoveLeft(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)

	x, y := robotgo.GetMousePos()
	robotgo.Move(x-int(move), y)

	ctx.JSON(http.StatusOK, util.Success())
}

//MouseMoveRight 鼠标右移
func MouseMoveRight(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)

	x, y := robotgo.GetMousePos()
	robotgo.Move(x+int(move), y)

	ctx.JSON(http.StatusOK, util.Success())
}

//MouseMoveUp 鼠标上移
func MouseMoveUp(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)

	x, y := robotgo.GetMousePos()
	robotgo.Move(x, y-int(move))

	ctx.JSON(http.StatusOK, util.Success())
}

//MouseMoveDown 鼠标下移
func MouseMoveDown(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)

	x, y := robotgo.GetMousePos()
	robotgo.Move(x, y+int(move))

	ctx.JSON(http.StatusOK, util.Success())
}

//MouseMove 移动鼠标
func MouseMove(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)
	strangle := ctx.DefaultQuery("angle", "90") //移动角度
	angle, _ := strconv.ParseInt(strangle, 10, 64)
	angle = angle % 360
	radian := math.Pi / float64(180) * float64(angle) //弧度

	y1 := math.Sin(radian) * float64(move)
	x1 := math.Cos(radian) * float64(move)
	x, y := robotgo.GetMousePos()
	robotgo.Move(x+int(x1), y+int(y1))
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseLeftClick 鼠标左键单击
func MouseLeftClick(ctx *gin.Context) {
	robotgo.Click()
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseLeftDoubleClick 鼠标左键双击
func MouseLeftDoubleClick(ctx *gin.Context) {
	robotgo.Click("left", true)
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseRightClick 鼠标右键单击
func MouseRightClick(ctx *gin.Context) {
	robotgo.Click("right")
	ctx.JSON(http.StatusOK, util.Success())
}

/***********拖拽******************/
//MouseLeftToggle 鼠标左键下压
func MouseLeftToggle(ctx *gin.Context) {
	robotgo.Toggle("left")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseLeftToggleUp 鼠标左键下压解放
func MouseLeftToggleUp(ctx *gin.Context) {
	robotgo.Toggle("left", "up")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseRightoggle 鼠标右键下压
func MouseRightoggle(ctx *gin.Context) {
	robotgo.Toggle("right")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseRightoggleUp 鼠标右键下压解放
func MouseRightoggleUp(ctx *gin.Context) {
	robotgo.Toggle("right", "up")
	ctx.JSON(http.StatusOK, util.Success())
}

/**************************/

//MouseCenterClick 鼠标滚轮单击
func MouseCenterClick(ctx *gin.Context) {
	robotgo.Click("center")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseWheelDown 鼠标滚轮下压
func MouseWheelDown(ctx *gin.Context) {
	robotgo.Click("wheelDown")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseWheelUp 鼠标滚轮上还原
func MouseWheelUp(ctx *gin.Context) {
	robotgo.Click("wheelUp")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseWheelLeft 鼠标滚轮左
func MouseWheelLeft(ctx *gin.Context) {
	robotgo.Click("wheelLeft")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseWheelRight 鼠标滚轮右
func MouseWheelRight(ctx *gin.Context) {
	robotgo.Click("wheelRight")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseScrollMouseUp 鼠标滑轮向上滚动
func MouseScrollMouseUp(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)
	robotgo.ScrollMouse(int(move), "up")
	ctx.JSON(http.StatusOK, util.Success())
}

//MouseScrollMouseDown 鼠标滑轮向下滚动
func MouseScrollMouseDown(ctx *gin.Context) {
	strMove := ctx.DefaultQuery("move", "10") //移动距离
	move, _ := strconv.ParseInt(strMove, 10, 64)
	robotgo.ScrollMouse(int(move), "down")
	ctx.JSON(http.StatusOK, util.Success())
}
