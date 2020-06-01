package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	frontBuildFolder = "./lizee/build"
)

// Server is the interface to interact with server
type Server struct {
	*gin.Engine
}

// Setup init server and return it's instance
func Setup() *Server {
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile(frontBuildFolder, true)))
	server := &Server{router}

	// create productsApi
	server.productsAPI()
	return server
}

func (s *Server) productsAPI() {
	// Setup route group for the API
	productAPI := s.Group("/products")
	productAPI.GET("/availability", checkProductAvailability)
	productAPI.POST("/products", post)

	categoryAPI := s.Group("/categories")
	categoryAPI.GET("/", listCategories)
	categoryAPI.GET("/products", checkCategoryAvailability)
}

// Serve starts server on provided port
func (s *Server) Serve(port string) {
	log.Print("starting server in dev mode on port ", port)
	err := s.Run(port)
	if err != nil {
		panic(err)
	}
}
