package router

import (
	"server/server/internal/user"
	"server/server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler,wsHandler *ws.Handler){
	r = gin.Default()
	r.POST("/signup",userHandler.CreateUser)
	r.POST("/Login",userHandler.Login)
	r.GET("/Logout",userHandler.Logout)

	r.POST("ws/createRoom",wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start (addr string) error{
    return r.Run(addr)
}