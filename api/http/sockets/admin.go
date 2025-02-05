// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sockets

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
)

var (
	// adminUpgrade is the websocket upgrader for any
	// admin routes for the SPA.
	adminUpgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// AdminData represents the struct of data to be sent
// back to the SPA via the websocket.
type AdminData struct {
	Theme domain.ThemeConfig `json:"theme"`
}

// AdminHub are the channels of broadcast and receiving
// messages for the admin socket.
var AdminHub = struct {
	// Channel for sending AdminData
	Broadcast chan AdminData
	// Channel for receiving any data.
	Receive chan []byte
	// Channel for closing the socket
	close chan bool
}{
	Broadcast: make(chan AdminData),
	Receive:   nil,
	close:     make(chan bool),
}

// Handler is the http.Handler for upgrading the http request
// to a websocket.
// Logs errors.INVALID if the request could not be upgraded.
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "Admin.Handler"
		conn, err := adminUpgrade.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error upgrading request to websocket", Operation: op, Err: err})
			return
		}
		run(conn)
	}
}

// run runs the web socket connection, marshals the AdminData on
// the Broadcast channel and closes the connection if any data
// is received on AdminHub.Close
func run(conn *websocket.Conn) {
	const op = "Sockets.Admin.Run"
	for {
		select {
		case as := <-AdminHub.Broadcast:
			b, err := json.Marshal(as)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error sending admin websocket broadcast", Operation: op, Err: err}).Error()
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error sending admin websocket broadcast", Operation: op, Err: err}).Error()
				return
			}
			logger.Info("Sent admin websocket")
		case <-AdminHub.close:
			logger.Debug("Closing admin websocket connection")
			conn.Close()
		}
	}
	conn.Close()
}
