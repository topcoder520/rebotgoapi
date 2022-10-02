package keyboard

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/topcoder520/rebotgoapi/util"
)

//KeyInputString 输入内容
func KeyInputString(ctx *gin.Context) {
	data := make(map[string]interface{})
	ctx.BindJSON(&data)
	fmt.Println(data)
	v, ok := data["content"]
	if ok {
		content := v.(string)
		if len(content) > 0 {
			robotgo.TypeStr(content)
		} else {
			robotgo.KeyTap(robotgo.Enter)
		}
	}
	ctx.JSON(http.StatusOK, util.Success())
}

//KeyTap 按键
func KeyTap(ctx *gin.Context) {
	key := ctx.DefaultQuery("key", "")
	fmt.Println(key)
	if len(key) > 0 {
		if strings.Contains(key, ",") {
			arr := strings.Split(key, ",")
			if len(arr) > 1 {
				robotgo.KeyTap(arr[0], arr[1:])
			} else {
				robotgo.KeyTap(arr[0])
			}
		} else {
			robotgo.KeyTap(key)
		}

	}
	ctx.JSON(http.StatusOK, util.Success())
}
