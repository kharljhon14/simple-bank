package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kharljhon14/simple-bank/db/sqlc"
)

// Server serves http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add routes to router
	router.GET("/api/v1/health", server.healthCheckHandler)

	router.GET("/api/v1/accounts/", server.getAccountListHandler)
	router.GET("/api/v1/accounts/:id", server.getAccountHandler)
	router.POST("/api/v1/accounts", server.createAccountHandler)
	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
