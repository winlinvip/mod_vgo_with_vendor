package core

import (
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"net/url"
)

func NewClient(netConn net.Conn, u *url.URL, requestHeader http.Header, readBufSize, writeBufSize int) (c *websocket.Conn, response *http.Response, err error) {
	return websocket.NewClient(netConn, u, requestHeader, readBufSize, writeBufSize)
}
