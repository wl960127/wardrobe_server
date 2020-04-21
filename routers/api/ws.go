package api

import (
	"time"
	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/msg"

	// "fmt"

	wsservice "wardrobe_server/service/wsService"

	"net/http"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

/** http://www.websocket-test.com/  测试 */

// WsPage .
func WsPage(c *gin.Context) {

	appG := app.Gin{C: c}

	var (
		conn   *wsservice.Connection
		err    error
		data   []byte
		wsConn *websocket.Conn
	)

	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		appG.Response(http.StatusBadRequest, msg.INVALID_PARAMS, nil)

		return
	}

	if conn, err = wsservice.InitConnection(wsConn); err != nil {
		goto ERR
	}

	/**
	  发送心跳
	*/
	go func() {

		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("心跳")); err != nil {
				return
			}
			time.Sleep(3 * time.Second)
		}

	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
