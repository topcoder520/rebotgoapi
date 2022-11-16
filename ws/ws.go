package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func CaptureScreen(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)
	w, h := robotgo.GetScreenSize()
	bitmap := robotgo.CaptureScreen(0, 0, w, h)
	image := robotgo.ToImage(bitmap)
	imageByte := robotgo.ToByteImg(image) //base64
	err = c.WriteMessage(websocket.TextMessage, imageByte)
	if err != nil {
		log.Println("write image:", err)
	}
}
