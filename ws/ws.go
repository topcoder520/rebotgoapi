package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type TextMessage struct {
	cmd string
}

func CaptureScreen(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	typ, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	if typ == websocket.TextMessage {
		reader := strings.NewReader(string(message))
		decoder := json.NewDecoder(reader)
		msg := TextMessage{}
		fmt.Println(decoder.Decode(&msg))
		fmt.Printf("%v", msg)
	}
	log.Printf("recv: %s", message)
	w, h := robotgo.GetScreenSize()
	bitmap := robotgo.CaptureScreen(0, 0, w, h)
	image := robotgo.ToImage(bitmap)
	//robotgo.SaveJpeg(image, "E:\\huangjing\\GoWorkSpace\\src\\github.com\\topcoder520\\rebotgoapi\\test.jpg")
	imageByte := robotgo.ToByteImg(image) //base64
	err = c.WriteMessage(websocket.TextMessage, imageByte)
	if err != nil {
		log.Println("write image:", err)
	}
}
