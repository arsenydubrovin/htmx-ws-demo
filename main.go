package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Static("/", "./public")

	app.Get("/ws", websocket.New(handleWebSocket))

	go broadcastNumbers()

	log.Fatal(app.Listen(":3000"))
}

func handleWebSocket(c *websocket.Conn) {
	clients[c] = true
	log.Println("client connected")

	defer func() {
		delete(clients, c)
		c.Close()
		log.Println("client disconnected")
	}()

	for {
		messageType, _, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		if messageType == websocket.CloseMessage {
			break
		}
	}
}

func broadcastNumbers() {
	for {
		randomNum := rand.Intn(100)
		html := fmt.Sprintf(`<div id="number">%d</div>`, randomNum)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(html))
			if err != nil {
				log.Println("write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		time.Sleep(time.Second)
	}
}
