package ws

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (hub *Hub) Run() {
	go func() {
		hub.CaptureScreen()
	}()
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
			}
		}
	}
}

func (hub *Hub) CaptureScreen() {
	ticker := time.NewTicker(time.Second * 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	for {
		select {
		case <-sig:
			for c := range hub.clients {
				c.conn.Close()
			}
		case <-ticker.C:
			if len(hub.clients) > 0 {
				w, h := robotgo.GetScreenSize()
				bitmap := robotgo.CaptureScreen(0, 0, w, h)
				image := robotgo.ToImage(bitmap)
				//robotgo.SaveJpeg(image, "E:\\huangjing\\GoWorkSpace\\src\\github.com\\topcoder520\\rebotgoapi\\test.jpg")
				imageByte := robotgo.ToByteImg(image) //base64
				for c, b := range hub.clients {
					if b {
						err := c.conn.WriteMessage(websocket.TextMessage, imageByte)
						if err != nil {
							log.Println("write image:", err)
						}
					}
				}
			}
		}
	}
}
