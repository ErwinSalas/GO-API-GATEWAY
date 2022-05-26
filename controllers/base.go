package controllers

import (
	"fmt"
	"log"
	"net/http"

	inventorypb "github.com/ErwinSalas/inventory-service/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

const (
	address = "inventory-service:8081"
)

type Server struct {
	Router          *mux.Router
	InventoryClient inventorypb.InventoryClient
}

func (server *Server) Initialize() {
	fmt.Println("Initialize")

	server.Router = mux.NewRouter()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	server.InventoryClient = inventorypb.NewInventoryClient(conn)
	fmt.Println("initializeRoutes")

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
