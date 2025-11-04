package api

import (
	"net/http"
	"webchat/internal/model"
	"webchat/internal/store"
	"webchat/internal/ws"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	checkOrigin: func(r *http.Request) bool { return true },
}

func RegisterRoutes(e *echo.Echo, hub *ws.Hub) {
	e.GET("/", func(c echo.Context) error {
		return c.File("client/index.html")
	})
	e.GET("/ws", func(c echo.Context) error {
		uderId := c.QueryParam("user_id")
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		client := &ws.Client{
			ID:   userId,
			Conn: conn,
			Send: make(chan ws.OutboundMessage),
			Hub:  hub,
		}
		hub.Register <- client

		go client.ReadPump()
		go client.WritePump()
		return nil
	})
	e.GET("/messages/:user_id", func(c echo.Context) error {
		userId := c.Param("user_id")
		var msgs []model.Message
		store.DB.Where("sender_id=? OR receiver_id=?", userId, userId).Order("created_at asc").Find(&msgs)
		return c.JSON(http.StatusOK, msgs)
	})
}
