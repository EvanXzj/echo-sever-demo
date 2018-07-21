package api

import (
	"github.com/evanxzj/echo-server/models"
	"github.com/evanxzj/echo-server/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server : the echo server manager
type Server struct {
	Db *models.DB
	e  *echo.Echo
}

type someRoutes interface {
	mount()
}

type route struct {
	group *echo.Group
	db    *models.DB
}

// Start : init routes and start the server
func (s *Server) Start() {

	s.e = echo.New()

	// get config params
	port := utils.GetEnv("APP_PORT")
	// jwtSecret := utils.GetEnv("JWT_SECRET")

	// some general middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// create the route groups
	rootGroup := s.e.Group("")
	userGroup := s.e.Group("/users")

	// auth middleware
	// userGroup.Use(middleware.JWT([]byte(jwtSecret)))

	// define base paths
	routes := []someRoutes{
		rootRoutes{route{rootGroup, s.Db}},
		userRoutes{route{userGroup, s.Db}},
	}

	// wire up sub paths
	for _, r := range routes {
		r.mount()
	}

	s.e.Logger.Fatal(s.e.Start(":" + port))

}
