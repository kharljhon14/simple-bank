package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/kharljhon14/simple-bank/db/sqlc"
	"github.com/kharljhon14/simple-bank/token"
)

// Server serves http requests
type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// Creates a new http server and setup routing
func NewServer(store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker("WOW,MuchShibe,ToDoggesadasdasdadas")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker, %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.mountRouters()
	return server, nil
}

func (s *Server) mountRouters() {
	router := gin.Default()

	// Add routes to router
	router.GET("/api/v1/health", s.healthCheckHandler)

	router.POST("/api/v1/users", s.createUserHandler)
	router.POST("/api/v1/login", s.loginUserHandler)

	autRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	autRoutes.GET("/api/v1/accounts/", s.getAccountListHandler)
	autRoutes.GET("/api/v1/accounts/:id", s.getAccountHandler)
	autRoutes.POST("/api/v1/accounts", s.createAccountHandler)
	autRoutes.POST("/api/v1/transfer", s.transferHandler)

	s.router = router

}

// Start runs the HTTP server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
