package chatcontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"rent-house/websocket/chatservice/models"
	"time"
)

type WebsocketController struct {
	beego.Controller
}

var (
	Clients  = make(map[string]map[*websocket.Conn]bool)     // connected Clients
	Bc       = make(map[string]chan models.Broadcast) // broadcastbody channel
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// @Title WebSocket
// @Description WebSocket
// @router /
func (w *WebsocketController) Join() {
	ws, err := upgrader.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	_, b, err := ws.ReadMessage()
	if err != nil {
		w.Ctx.WriteString(err.Error())
		return
	}
	ownerID := string(b)
	Clients[ownerID] = make(map[*websocket.Conn]bool)
	Clients[ownerID][ws] = true
	Bc[ownerID] = make(chan models.Broadcast)
	ws.WriteJSON(map[string]string{
		"code" : "200",
		"message" : "success",
	})
	go broadcastMessages(Bc[ownerID])
	for {
		var batm models.Message
		broadcast := models.Broadcast{}
		// Read in a new messagebody as JSON and map it to a MoveMessage object
		_, b, err = ws.ReadMessage()
		err = json.Unmarshal(b, &batm)
		if err != nil {
			ws.WriteJSON(map[string]interface{}{
				"code" : "400",
				"message" : err.Error(),
			})
			ws.Close()
			return
		}
		broadcast.Message = batm
		broadcast.SendTime = time.Now().Unix()
		// Send the newly received messagebody to the broadcastbody channel
		Bc[ownerID] <- broadcast
	}
}

func broadcastMessages(bs <- chan models.Broadcast) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range bs {
			for client := range Clients[i.OwnerID] {
				err := client.WriteJSON(i)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(Clients[i.OwnerID], client)
				}
			}
		}
		// Send it out to every client that is currently connected
	}
}