package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	frontBuildFolder = "./frontBuild"
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

// productsAPI initialize all API concerning products
func (s *Server) productsAPI() {
	// Setup route group for the API
	productAPI := s.Group("/products")
	// Check if one product is available
	// ====> GET http://localhost:5000/products/availability?product_id=6&from=2020-06-04&to=2020-06-05
	productAPI.GET("/availability", checkProductAvailability)
	// Pass an order
	// ====> POST http://localhost:5000/products/order
	// ====> [{"product":{"availability":4,"id":1,"name":"tente trekking UL3","picture":""},"quantity":1,"from":"2020-06-03","to":"2020-06-04"}]
	productAPI.POST("/order", postOrder)

	categoryAPI := s.Group("/categories")
	// Get all existing categories of product
	// ====> GET http://localhost:5000/categories
	categoryAPI.GET("/", listCategories)
	// Check which products are available with category
	// ====> GET http://localhost:5000/categories/products?categoryID=1&from=2020-06-04&to=2020-06-05
	categoryAPI.GET("/products", checkCategoryAvailability)

	// Exercize purpose basic API
	avaiabilityAPI := s.Group("/availability")
	// Modify quantity of corresponding product in database
	// =====> POST to http://localhost:5000/availability/changeQuantity
	// =====> {"product_id":int, "quantity": int}
	// note that product_id is integer anywhere else than
	// checkProductsAvailability to correspond to your demand.
	avaiabilityAPI.POST("/modifyquantity", modifyQuantity)
	// Get all available product between these dates
	// =====> POST http://localhost:5000/availability/
	// =====> {"from":"2023-06-04","to":"2023-06-05"}
	avaiabilityAPI.POST("/", checkProductsAvailability)
}

// Serve starts server on provided port
func (s *Server) Serve(port string) {
	log.Print("starting server in dev mode on port ", port)
	err := s.Run(port)
	if err != nil {
		panic(err)
	}
}
