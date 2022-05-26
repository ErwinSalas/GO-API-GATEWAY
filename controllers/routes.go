package controllers

import (
	"fmt"

	"github.com/ErwinSalas/webui/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	fmt.Println("Home")

	//Users routes
	s.Router.HandleFunc("/inventory", middlewares.SetMiddlewareJSON(s.ListItems)).Methods("GET")
	fmt.Println("/inventory")
	s.Router.HandleFunc("/inventory/{id}", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("GET")
	fmt.Println("/inventory/{id}")

}
