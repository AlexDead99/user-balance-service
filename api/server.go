package api

import (
	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store, router: router}

	router.POST("/accounts", server.createAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.GET("/accounts/:id", server.getAccountInfo)
	router.POST("/transfers", server.createTransfer)
	router.PUT("/transfers", server.fulfilTransfer)
	router.POST("/report", server.CreateMonthReport)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/static", "./app/reports")

	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
