package chatcontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"rent-house/middlewares"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice/models"
	"time"
)

type WebsocketController struct {
	beego.Controller
}

var (
	//connected client
	Clients  = make(map[string]*websocket.Conn)     // connected Clients
	//admin receiver
	Admin    = make(map[string]*websocket.Conn)
	//broadcast to admin channel
	BcAdmin = make(map[string]chan models.BroadCastToAdmin) // broadcastbody channel
	//broadcast to client channel
	BcOwner = make(map[string]chan models.BroadCastToOwner)
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
	BcAdmin[ownerID] = make(chan models.BroadCastToAdmin)
	ws.WriteJSON(response.NewErr(response.Success))
	go broadcastToAdmin(BcAdmin[ownerID])
	for {
		var msg models.OwnerMessage
		// Read in a new messagebody as JSON and map it to a MoveMessage object
		_, b, err = ws.ReadMessage()
		err = json.Unmarshal(b, &msg)
		if err != nil {
			ws.WriteJSON(response.BadRequest)
			ws.Close()
			delete(Clients, ownerID)
			return
		}
		bc := &models.BroadCastToAdmin{
			OwnerID:      ownerID,
			SendTime:     time.Now().Unix(),
			OwnerMessage: msg,
			Type: models.OWNER_MESSAGE,
		}
		go bc.PutItem()
		// Send the newly received messagebody to the broadcastbody channel
		BcAdmin[ownerID] <- *bc
	}
}

func broadcastToAdmin(msg <- chan models.BroadCastToAdmin) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range msg {
			for _, v := range Admin {
				v.WriteJSON(i)
			}
		}
		// Send it out to every admin that is currently connected
	}
}


// @Title WebSocket
// @Description WebSocket
// @router /admin
func (w *WebsocketController) JoinAdmin() {
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
	adminID := middlewares.GetAdminFromToken(tokenString)
	if adminID == "" {
		ws.WriteJSON(response.NewErr(response.BadRequest))
		return
	}
	Admin[adminID] = ws
	err = ws.WriteJSON(response.NewErr(response.Success))
	if err != nil {
		delete(Admin, adminID)
		return
	}
	BcOwner[adminID] = make(chan models.BroadCastToOwner)
	go broadcastToOwner(BcOwner[adminID])
	for {
		var msg models.AdminMessage
		// Read in a new messagebody as JSON and map it to a MoveMessage object
		_, b, err = ws.ReadMessage()
		err = json.Unmarshal(b, &msg)
		if err != nil {
			ws.WriteJSON(response.BadRequest)
			ws.Close()
			return
		}
		bc := &models.BroadCastToOwner{
			AdminID:      adminID,
			SendTime:     time.Now().Unix(),
			AdminMessage: msg,
			Type: models.ADMIN_MESSAGE,
		}

		go bc.PutItem()

		// Send the newly received messagebody to the broadcastbody channel
		BcOwner[adminID] <- *bc
	}
}

func broadcastToOwner(msg <- chan models.BroadCastToOwner) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range msg {
			if Clients[i.OwnerID] != nil {
				err := Clients[i.OwnerID].WriteJSON(i)
				if err != nil {
					delete(Clients, i.OwnerID)
				}
			}
		}
		// Send it out to every owner that is currently connected
	}
}

