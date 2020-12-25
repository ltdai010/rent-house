package chatcontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"rent-house/middlewares"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice/models"
	"time"
)

type WebsocketController struct {
	beego.Controller
}

// @Title WebSocket
// @Description this is WebSocket, don't try it
// @router /
func (w *WebsocketController) Join() {
	ws, err := models.Upgrader.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
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
	models.Clients[ownerID] = ws
	models.BcAdmin[ownerID] = make(chan models.BroadCastToAdmin)
	err = ws.WriteJSON(response.NewErr(response.Success))
	if err != nil {
		ws.Close()
		delete(models.Clients, ownerID)
		return
	}
	go broadcastToAdmin(models.BcAdmin[ownerID])
	for {
		var msg models.OwnerMessage
		// Read in a new messagebody as JSON and map it to a MoveMessage object
		_, b, err = ws.ReadMessage()
		err = json.Unmarshal(b, &msg)
		if err != nil {
			err = ws.WriteJSON(response.BadRequest)
			if err != nil {
				ws.Close()
				delete(models.Clients, ownerID)
				return
			}
		}
		bc := &models.BroadCastToAdmin{
			OwnerID:      ownerID,
			SendTime:     time.Now().Unix(),
			OwnerMessage: msg,
			Type: models.OWNER_MESSAGE,
		}
		go bc.PutItem()
		// Send the newly received messagebody to the broadcastbody channel
		if models.BcAdmin != nil {
			if models.BcAdmin[ownerID] != nil {
				models.BcAdmin[ownerID] <- *bc
				err = ws.WriteJSON(response.NewErr(response.Success))
				if err != nil {
					ws.Close()
					delete(models.Clients, ownerID)
					return
				}
			}
		}
	}
}

func broadcastToAdmin(msg <- chan models.BroadCastToAdmin) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range msg {
			for _, v := range models.Admin {
				v.WriteJSON(i)
			}
		}
		// Send it out to every admin that is currently connected
	}
}


// @Title WebSocket
// @Description this is WebSocket, don't try it
// @router /admin
func (w *WebsocketController) JoinAdmin() {
	ws, err := models.Upgrader.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
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
	models.Admin[adminID] = ws
	err = ws.WriteJSON(response.NewErr(response.Success))
	if err != nil {
		delete(models.Admin, adminID)
		return
	}
	models.BcOwner[adminID] = make(chan models.BroadCastToOwner)
	go broadcastToOwner(models.BcOwner[adminID])
	for {
		var msg models.AdminMessage
		// Read in a new messagebody as JSON and map it to a MoveMessage object
		_, b, err = ws.ReadMessage()
		err = json.Unmarshal(b, &msg)
		if err != nil {
			err = ws.WriteJSON(response.BadRequest)
			if err != nil {
				ws.Close()
				delete(models.Admin, adminID)
				return
			}
			continue
		}
		bc := &models.BroadCastToOwner{
			AdminID:      adminID,
			SendTime:     time.Now().Unix(),
			AdminMessage: msg,
			Type: models.ADMIN_MESSAGE,
		}

		go bc.PutItem()

		// Send the newly received messagebody to the broadcastbody channel
		if models.BcOwner != nil {
			if models.BcOwner[adminID] != nil {
				models.BcOwner[adminID] <- *bc
				err = ws.WriteJSON(response.NewErr(response.Success))
				if err != nil {
					ws.Close()
					delete(models.Admin, adminID)
					return
				}
			}
		}
	}
}

func broadcastToOwner(msg <- chan models.BroadCastToOwner) {
	for {
		// Grab the next messagebody from the broadcastbody channel
		for i := range msg {
			if models.Clients[i.OwnerID] != nil {
				err := models.Clients[i.OwnerID].WriteJSON(i)
				if err != nil {
					delete(models.Clients, i.OwnerID)
					return
				}
			}
		}
		// Send it out to every owner that is currently connected
	}
}

