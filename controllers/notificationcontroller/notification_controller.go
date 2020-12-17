package notificationcontroller

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"rent-house/middlewares"
	"rent-house/restapi/response"
	"rent-house/websocket/notificationservice/models"
)

type WebsocketController struct {
	beego.Controller
}

var (
	//connected client
	Clients  = make(map[string]*websocket.Conn)     // connected Clients
	//broadcast to owner
	Broadcast = make(map[string]chan models.Notification) // broadcast channel
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
	tokenString := string(b)
	ownerID := middlewares.GetOwnernameFromToken(tokenString)
	if ownerID == "" {
		ws.WriteJSON(response.BadRequest)
		return
	}
	Clients[ownerID] = ws
	Broadcast[ownerID] = make(chan models.Notification)
	ws.WriteJSON(response.Success)
	go broadcastToUser(Broadcast[ownerID])
	for {
		// Read in a new messagebody as JSON, if err disconnect to client
		_, _, err = ws.ReadMessage()
		if err != nil {
			ws.Close()
			delete(Clients, ownerID)
			return
		}
	}
}

func broadcastToUser(msg <- chan models.Notification) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range msg {
			err := Clients[i.OwnerID].WriteJSON(i)
			if err != nil {
				Clients[i.OwnerID].Close()
				delete(Clients, i.OwnerID)
				return
			}
		}
		// Send it out to every admin that is currently connected
	}
}